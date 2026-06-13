//go:build ignore

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

type PokemonTypeRef struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PokemonTypeList struct {
	Count   int              `json:"count"`
	Next    string           `json:"next"`
	Results []PokemonTypeRef `json:"results"`
}

type PokemonTypeRow struct {
	PokemonID int
	TypeID    int
	Slot      int
}

func createPokemonTypesTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS pokemon_types (
		pokemon_id  INTEGER NOT NULL REFERENCES pokemons(id),
		type_id     INTEGER NOT NULL REFERENCES types(id),
		slot        INTEGER NOT NULL,
		PRIMARY KEY (pokemon_id, slot)
	)`)
	if err != nil {
		log.Fatal(err)
	}
}

func fetchPokemonTypes(url string) ([]PokemonTypeRow, error) {
	r, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("fetching pokemon: %w", err)
	}
	defer r.Body.Close()

	var raw map[string]any
	if err := json.NewDecoder(r.Body).Decode(&raw); err != nil {
		return nil, fmt.Errorf("decoding pokemon: %w", err)
	}

	pokemonID := int(raw["id"].(float64))

	typesRaw, _ := raw["types"].([]any)

	rows := make([]PokemonTypeRow, 0, len(typesRaw))
	for _, t := range typesRaw {
		tMap, _ := t.(map[string]any)

		slot := int(tMap["slot"].(float64))

		typeRes, _ := tMap["type"].(map[string]any)
		typeURL, _ := typeRes["url"].(string)

		var typeID int
		fmt.Sscanf(typeURL, "http://localhost:8000/api/v2/type/%d/", &typeID)
		if typeID == 0 {
			continue
		}

		rows = append(rows, PokemonTypeRow{
			PokemonID: pokemonID,
			TypeID:    typeID,
			Slot:      slot,
		})
	}

	return rows, nil
}

func main() {
	db, err := sql.Open("sqlite3", "./pokemon.db?_busy_timeout=5000")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	createPokemonTypesTable(db)

	var allPokemon []PokemonTypeRef
	url := "http://localhost:8000/api/v2/pokemon?limit=100&offset=0"
	for url != "" {
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		var list PokemonTypeList
		json.NewDecoder(resp.Body).Decode(&list)
		resp.Body.Close()
		allPokemon = append(allPokemon, list.Results...)
		url = list.Next
		fmt.Printf("Collected %d pokemon so far...\n", len(allPokemon))
	}
	fmt.Printf("Total: %d pokemon to process\n", len(allPokemon))

	const workers = 10
	jobs := make(chan PokemonTypeRef, len(allPokemon))
	results := make(chan []PokemonTypeRow, workers)

	var fetchWg sync.WaitGroup
	for i := 0; i < workers; i++ {
		fetchWg.Add(1)
		go func() {
			defer fetchWg.Done()
			for p := range jobs {
				rows, err := fetchPokemonTypes(p.URL)
				if err != nil {
					log.Println("Failed to fetch", p.Name, err)
					continue
				}
				results <- rows
			}
		}()
	}

	go func() {
		fetchWg.Wait()
		close(results)
	}()

	for _, p := range allPokemon {
		jobs <- p
	}
	close(jobs)

	for rows := range results {
		if len(rows) == 0 {
			continue
		}
		pokemonID := rows[0].PokemonID

		tx, err := db.Begin()
		if err != nil {
			log.Println("Failed to begin tx for pokemon", pokemonID, err)
			continue
		}

		for _, row := range rows {
			_, err := tx.Exec(`
				INSERT OR REPLACE INTO pokemon_types (pokemon_id, type_id, slot)
				VALUES (?, ?, ?)`,
				row.PokemonID, row.TypeID, row.Slot,
			)
			if err != nil {
				log.Println("Insert failed for pokemon", row.PokemonID, "type", row.TypeID, err)
			}
		}

		if err := tx.Commit(); err != nil {
			log.Println("Commit failed for pokemon", pokemonID, err)
			continue
		}

		log.Printf("Wrote %d type rows for pokemon %d\n", len(rows), pokemonID)
	}

	fmt.Println("Done.")
}
