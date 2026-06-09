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

type PokemonMoveRef struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PokemonMoveList struct {
	Count   int              `json:"count"`
	Next    string           `json:"next"`
	Results []PokemonMoveRef `json:"results"`
}

type PokemonMoveRow struct {
	PokemonID   int
	MoveID      int
	MinLevel    int
	LearnMethod string
}

func createPokemonMovesTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS pokemon_moves (
		pokemon_id    INTEGER NOT NULL REFERENCES pokemons(id),
		move_id       INTEGER NOT NULL REFERENCES moves(id),
		min_level     INTEGER NOT NULL,
		learn_method  TEXT    NOT NULL CHECK(learn_method IN ('LevelUp', 'Egg', 'Tutor', 'Machine')),
		PRIMARY KEY (pokemon_id, move_id, learn_method)
	)`)
	if err != nil {
		log.Fatal(err)
	}
}

// Maps API learn method names to our enum values.
// Only the four methods in the enum are kept; others are skipped.
var learnMethodMap = map[string]string{
	"level-up": "LevelUp",
	"egg":      "Egg",
	"tutor":    "Tutor",
	"machine":  "Machine",
}

func fetchPokemonMoves(url string) ([]PokemonMoveRow, error) {
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

	movesRaw, _ := raw["moves"].([]any)

	// Deduplicate: keep the lowest min_level per (pokemon_id, move_id, learn_method)
	type key struct {
		MoveID      int
		LearnMethod string
	}
	best := make(map[key]int) // key -> min_level

	for _, m := range movesRaw {
		mMap, _ := m.(map[string]any)

		moveRes, _ := mMap["move"].(map[string]any)
		moveURL, _ := moveRes["url"].(string)
		var moveID int
		fmt.Sscanf(moveURL, "http://localhost:8000/api/v2/move/%d/", &moveID)
		if moveID == 0 {
			continue
		}

		versionDetails, _ := mMap["version_group_details"].([]any)
		for _, vd := range versionDetails {
			vdMap, _ := vd.(map[string]any)

			methodRes, _ := vdMap["move_learn_method"].(map[string]any)
			methodName, _ := methodRes["name"].(string)

			enumVal, ok := learnMethodMap[methodName]
			if !ok {
				continue // skip methods not in our enum
			}

			level := 0
			if v, ok := vdMap["level_learned_at"].(float64); ok {
				level = int(v)
			}

			k := key{MoveID: moveID, LearnMethod: enumVal}
			if existing, seen := best[k]; !seen || level < existing {
				best[k] = level
			}
		}
	}

	rows := make([]PokemonMoveRow, 0, len(best))
	for k, level := range best {
		rows = append(rows, PokemonMoveRow{
			PokemonID:   pokemonID,
			MoveID:      k.MoveID,
			MinLevel:    level,
			LearnMethod: k.LearnMethod,
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

	createPokemonMovesTable(db)

	// Collect all pokemon URLs across all pages
	var allPokemon []PokemonMoveRef
	url := "http://localhost:8000/api/v2/pokemon?limit=100&offset=0"
	for url != "" {
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		var list PokemonMoveList
		json.NewDecoder(resp.Body).Decode(&list)
		resp.Body.Close()
		allPokemon = append(allPokemon, list.Results...)
		url = list.Next
		fmt.Printf("Collected %d pokemon so far...\n", len(allPokemon))
	}
	fmt.Printf("Total: %d pokemon to process\n", len(allPokemon))

	const workers = 10
	jobs := make(chan PokemonMoveRef, len(allPokemon))
	results := make(chan []PokemonMoveRow, workers)

	var fetchWg sync.WaitGroup
	for i := 0; i < workers; i++ {
		fetchWg.Add(1)
		go func() {
			defer fetchWg.Done()
			for p := range jobs {
				rows, err := fetchPokemonMoves(p.URL)
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
				INSERT OR REPLACE INTO pokemon_moves (pokemon_id, move_id, min_level, learn_method)
				VALUES (?, ?, ?, ?)`,
				row.PokemonID, row.MoveID, row.MinLevel, row.LearnMethod,
			)
			if err != nil {
				log.Println("Insert failed for pokemon", row.PokemonID, "move", row.MoveID, err)
			}
		}

		if err := tx.Commit(); err != nil {
			log.Println("Commit failed for pokemon", pokemonID, err)
			continue
		}

		log.Printf("Wrote %d move rows for pokemon %d\n", len(rows), pokemonID)
	}

	fmt.Println("Done.")
}
