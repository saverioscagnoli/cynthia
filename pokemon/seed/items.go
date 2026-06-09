//go:build ignore

package main

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type List struct {
	Count   int    `json:"count"`
	Next    string `json:"next"`
	Results []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type Item struct {
	ID          int
	Name        string
	Cost        int
	FlingPower  *int // pointer = nullable
	FlingEffect *string
	Sprite      []byte
}

func createTable(db *sql.DB) {
	_, err := db.Exec(`
    CREATE TABLE IF NOT EXISTS items (
        id           INTEGER PRIMARY KEY,
        name         TEXT    NOT NULL,
        cost         INTEGER NOT NULL,
        fling_power  INTEGER,
        fling_effect TEXT,
        sprite       BLOB
    )`)
	if err != nil {
		log.Fatal(err)
	}
}

func getEnglishEffect(raw map[string]any) *string {
	effect, ok := raw["fling_effect"]
	if !ok || effect == nil {
		return nil
	}

	effectMap, ok := effect.(map[string]any)
	if !ok {
		return nil
	}

	effectURL, ok := effectMap["url"].(string)
	if !ok {
		return nil
	}

	resp, err := http.Get(effectURL)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	var flingEffect struct {
		Name string
	}
	json.NewDecoder(resp.Body).Decode(&flingEffect)

	return &flingEffect.Name
}

func downloadSprite(raw map[string]any) []byte {
	sprites, ok := raw["sprites"].(map[string]any)
	if !ok {
		return nil
	}

	spriteURL, ok := sprites["default"].(string)
	if !ok || spriteURL == "" {
		return nil
	}

	resp, err := http.Get(spriteURL)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil
	}
	return data
}

func main() {
	db, err := sql.Open("sqlite3", "./pokemon.db?_busy_timeout=5000")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	createTable(db)

	url := "http://localhost:8000/api/v2/item?limit=100&offset=0"
	for url != "" {
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		var list List
		json.NewDecoder(resp.Body).Decode(&list)
		resp.Body.Close()

		for _, p := range list.Results {
			r, err := http.Get(p.URL)
			if err != nil {
				log.Println("Failed to fetch", p.Name, err)
				continue
			}
			var raw map[string]any
			json.NewDecoder(r.Body).Decode(&raw)
			r.Body.Close()

			item := Item{
				ID:   int(raw["id"].(float64)),
				Name: raw["name"].(string),
				Cost: int(raw["cost"].(float64)),
			}

			if fp, ok := raw["fling_power"].(float64); ok {
				v := int(fp)
				item.FlingPower = &v
			}

			item.FlingEffect = getEnglishEffect(raw)

			item.Sprite = downloadSprite(raw)

			_, err = db.Exec(
				`INSERT OR REPLACE INTO items (id, name, cost, fling_power, fling_effect, sprite)
                 VALUES (?, ?, ?, ?, ?, ?)`,
				item.ID, item.Name, item.Cost, item.FlingPower, item.FlingEffect, item.Sprite,
			)
			if err != nil {
				log.Println("Insert failed for", item.Name, err)
				continue
			}

			log.Println("Wrote", item.ID, item.Name)
		}

		url = list.Next
	}
}
