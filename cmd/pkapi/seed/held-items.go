//go:build ignore

package main

import (
	"database/sql"
	"encoding/json"
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

type HeldItem struct {
	ItemID    int
	PokemonID int
	Rarity    int
}

func createTable(db *sql.DB) {
	_, err := db.Exec(`
    CREATE TABLE IF NOT EXISTS held_items (
    	item_id INTEGER NOT NULL,
     	pokemon_id INTEGER NOT NULL,
        rarity INTEGER
    )`)

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	db, err := sql.Open("sqlite3", "./pokemon.db?_busy_timeout=5000")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	createTable(db)

	url := "http://localhost:8000/api/v2/pokemon?limit=100&offset=0"

	for url != "" {
		resp, _ := http.Get(url)
		var list List
		json.NewDecoder(resp.Body).Decode(&list)
		resp.Body.Close()

		for _, p := range list.Results {
			r, _ := http.Get(p.URL)
			var pokemon map[string]any
			json.NewDecoder(r.Body).Decode(&pokemon)
			r.Body.Close()

			var heldItems []interface{}

			heldItems = pokemon["held_items"].([]interface{})

			log.Println(len(heldItems))

			if len(heldItems) > 0 {
				for _, item := range heldItems {
					itemCasted := item.(map[string]any)["item"].(map[string]any)
					itemUrl := itemCasted["url"]
					versionInfo := item.(map[string]any)["version_details"].([]interface{})

					var lastVersionInfo map[string]any
					for i := len(versionInfo) - 1; i >= 0; i-- {
						candidate := versionInfo[i].(map[string]any)
						if candidate["version"].(map[string]any)["name"].(string) != "xd" {
							lastVersionInfo = candidate
							break
						}
					}

					if lastVersionInfo == nil {
						log.Println("Skipped, all versions are xd")
						continue
					}

					rarity := lastVersionInfo["rarity"]

					res, _ := http.Get(itemUrl.(string))

					var item map[string]any
					json.NewDecoder(res.Body).Decode(&item)
					res.Body.Close()

					i := HeldItem{
						ItemID:    int(item["id"].(float64)),
						PokemonID: int(pokemon["id"].(float64)),
						Rarity:    int(rarity.(float64)),
					}

					_, err = db.Exec(
						`INSERT OR REPLACE INTO held_items (item_id, pokemon_id, rarity) VALUES (?, ?, ?)`,
						i.ItemID, i.PokemonID, i.Rarity,
					)
					if err != nil {
						log.Println("Insert failed:", err)
						continue
					}

					log.Println("Wrote item id", i.ItemID, "for pokemon id", i.PokemonID, "rarity =", i.Rarity)

				}
			} else {
				log.Println("Skipped")
			}
		}

		url = list.Next
	}
}
