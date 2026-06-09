//go:build ignore

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type PokemonList struct {
	Count   int    `json:"count"`
	Next    string `json:"next"`
	Results []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type PokemonSpecies struct {
	ID                int
	Name              string
	NationalDexNumber int
	GenderChance      int
	CaptureRate       int
	BaseHappiness     int
	IsBaby            bool
	IsLegendary       bool
	IsMythic          bool
	HatchCounter      int
	GenderDifferences bool
	FormsSwitchable   bool
	Color             string
}

func downloadSprite(url, path string) {
	if url == "" {
		return
	}
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		fmt.Printf("  failed to download %s\n", path)
		return
	}
	defer resp.Body.Close()

	file, err := os.Create(path)
	if err != nil {
		fmt.Printf("  failed to create %s\n", path)
		return
	}
	defer file.Close()
	io.Copy(file, resp.Body)
}

func createSpeciesTable(db *sql.DB) {
	_, err := db.Exec(`
    CREATE TABLE IF NOT EXISTS pokemon_species (
        id                INTEGER PRIMARY KEY,
        name              TEXT    NOT NULL,
        dex_number        INTEGER NOT NULL,
        gender_chance     INTEGER NOT NULL,
        capture_rate      INTEGER NOT NULL,
        base_happiness    INTEGER NOT NULL,
        baby              BOOLEAN NOT NULL,
        legendary         BOOLEAN NOT NULL,
        mythic            BOOLEAN NOT NULL,
        hatch_counter     INTEGER NOT NULL,
        gender_differences BOOLEAN NOT NULL,
        forms_switchable  BOOLEAN NOT NULL,
        color             TEXT    NOT NULL
    )`)

	if err != nil {
		log.Fatal(err)
	}
}

func insertSpecies(db *sql.DB, p PokemonSpecies) error {
	_, err := db.Exec(`
        INSERT OR REPLACE INTO pokemon_species
            (id, name, dex_number, gender_chance, capture_rate, base_happiness,
             baby, legendary, mythic, hatch_counter, gender_differences, forms_switchable, color)
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		p.ID, p.Name, p.NationalDexNumber, p.GenderChance, p.CaptureRate, p.BaseHappiness,
		p.IsBaby, p.IsLegendary, p.IsMythic, p.HatchCounter, p.GenderDifferences, p.FormsSwitchable, p.Color,
	)
	return err
}

func main() {
	db, err := sql.Open("sqlite3", "./pokemon.db?_busy_timeout=5000")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	url := "http://localhost:8000/api/v2/pokemon-species?limit=100&offset=0"

	createSpeciesTable(db)

	for url != "" {
		resp, _ := http.Get(url)
		var list PokemonList
		json.NewDecoder(resp.Body).Decode(&list)
		resp.Body.Close()

		for _, p := range list.Results {
			r, _ := http.Get(p.URL)
			var species map[string]any
			json.NewDecoder(r.Body).Decode(&species)
			r.Body.Close()

			var dexNumber int

			for _, entry := range species["pokedex_numbers"].([]any) {
				e := entry.(map[string]any)
				pokedex := e["pokedex"].(map[string]any)

				if pokedex["name"].(string) == "national" {
					dexNumber = int(e["entry_number"].(float64))
					break
				}
			}

			p := PokemonSpecies{
				ID:                int(species["id"].(float64)),
				Name:              species["name"].(string),
				NationalDexNumber: dexNumber,
				GenderChance:      int(species["gender_rate"].(float64)),
				CaptureRate:       int(species["capture_rate"].(float64)),
				BaseHappiness:     int(species["base_happiness"].(float64)),
				IsBaby:            species["is_baby"].(bool),
				IsLegendary:       species["is_legendary"].(bool),
				IsMythic:          species["is_mythical"].(bool),
				HatchCounter:      int(species["hatch_counter"].(float64)),
				GenderDifferences: species["has_gender_differences"].(bool),
				FormsSwitchable:   species["forms_switchable"].(bool),
				Color:             species["color"].(map[string]any)["name"].(string),
			}

			if err := insertSpecies(db, p); err != nil {
				log.Printf("failed to insert %s: %v", p.Name, err)
			}

			log.Println("Wrote", p.Name)
		}

		url = list.Next
	}
}
