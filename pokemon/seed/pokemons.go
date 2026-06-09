//go:build ignore

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

type PokemonRef struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type List struct {
	Count   int          `json:"count"`
	Next    string       `json:"next"`
	Results []PokemonRef `json:"results"`
}

type Pokemon struct {
	ID                     int
	Name                   string
	BaseExp                int
	Height                 int
	DefaultForm            bool
	Weight                 int
	SpeciesID              int
	SpriteFront            []byte
	SpriteFrontShiny       []byte
	SpriteBack             []byte
	SpriteBackShiny        []byte
	SpriteFrontFemale      []byte
	SpriteFrontShinyFemale []byte
	SpriteBackFemale       []byte
	SpriteBackShinyFemale  []byte
}

func createTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS pokemons (
		id                        INTEGER PRIMARY KEY,
		name                      TEXT    NOT NULL,
		base_exp                  INTEGER NOT NULL,
		height                    INTEGER NOT NULL,
		default_form              BOOLEAN NOT NULL,
		weight                    INTEGER NOT NULL,
		species_id                INTEGER NOT NULL,
		sprite_front              BLOB,
		sprite_front_shiny        BLOB,
		sprite_back               BLOB,
		sprite_back_shiny         BLOB,
		sprite_front_female       BLOB,
		sprite_front_shiny_female BLOB,
		sprite_back_female        BLOB,
		sprite_back_shiny_female  BLOB
	)`)
	if err != nil {
		log.Fatal(err)
	}
}

func fetchSprite(url string) []byte {
	if url == "" {
		return nil
	}
	resp, err := http.Get(url)
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

func spriteStr(sprites map[string]any, key string) string {
	v, ok := sprites[key]
	if !ok || v == nil {
		return ""
	}
	s, _ := v.(string)
	return s
}

func fetchPokemon(url string) (Pokemon, error) {
	r, err := http.Get(url)
	if err != nil {
		return Pokemon{}, fmt.Errorf("fetching pokemon: %w", err)
	}
	defer r.Body.Close()

	var raw map[string]any
	if err := json.NewDecoder(r.Body).Decode(&raw); err != nil {
		return Pokemon{}, fmt.Errorf("decoding pokemon: %w", err)
	}

	sprites, _ := raw["sprites"].(map[string]any)
	species, _ := raw["species"].(map[string]any)

	speciesID := 0
	if speciesURL, ok := species["url"].(string); ok {
		sr, err := http.Get(speciesURL)
		if err == nil {
			var speciesRaw map[string]any
			json.NewDecoder(sr.Body).Decode(&speciesRaw)
			sr.Body.Close()
			if sid, ok := speciesRaw["id"].(float64); ok {
				speciesID = int(sid)
			}
		}
	}

	pk := Pokemon{
		ID:                     int(raw["id"].(float64)),
		Name:                   raw["name"].(string),
		BaseExp:                int(raw["base_experience"].(float64)),
		Height:                 int(raw["height"].(float64)),
		DefaultForm:            raw["is_default"].(bool),
		Weight:                 int(raw["weight"].(float64)),
		SpeciesID:              speciesID,
		SpriteFront:            fetchSprite(spriteStr(sprites, "front_default")),
		SpriteFrontShiny:       fetchSprite(spriteStr(sprites, "front_shiny")),
		SpriteBack:             fetchSprite(spriteStr(sprites, "back_default")),
		SpriteBackShiny:        fetchSprite(spriteStr(sprites, "back_shiny")),
		SpriteFrontFemale:      fetchSprite(spriteStr(sprites, "front_female")),
		SpriteFrontShinyFemale: fetchSprite(spriteStr(sprites, "front_shiny_female")),
		SpriteBackFemale:       fetchSprite(spriteStr(sprites, "back_female")),
		SpriteBackShinyFemale:  fetchSprite(spriteStr(sprites, "back_shiny_female")),
	}

	return pk, nil
}

func main() {
	db, err := sql.Open("sqlite3", "./pokemon.db?_busy_timeout=5000")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	createTable(db)

	// Collect all pokemon URLs across all pages first
	var allPokemon []PokemonRef
	url := "http://localhost:8000/api/v2/pokemon?limit=100&offset=0"
	for url != "" {
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		var list List
		json.NewDecoder(resp.Body).Decode(&list)
		resp.Body.Close()
		allPokemon = append(allPokemon, list.Results...)
		url = list.Next
		fmt.Printf("Collected %d pokemon so far...\n", len(allPokemon))
	}
	fmt.Printf("Total: %d pokemon to fetch\n", len(allPokemon))

	const workers = 10
	jobs := make(chan PokemonRef, len(allPokemon))
	results := make(chan Pokemon, workers)

	// Fan out: workers fetch pokemon concurrently
	var fetchWg sync.WaitGroup
	for i := 0; i < workers; i++ {
		fetchWg.Add(1)
		go func() {
			defer fetchWg.Done()
			for p := range jobs {
				pk, err := fetchPokemon(p.URL)
				if err != nil {
					log.Println("Failed to fetch", p.Name, err)
					continue
				}
				results <- pk
			}
		}()
	}

	// Close results when all workers are done
	go func() {
		fetchWg.Wait()
		close(results)
	}()

	// Enqueue all jobs
	for _, p := range allPokemon {
		jobs <- p
	}
	close(jobs)

	// Single writer goroutine drains results and writes to DB
	for pk := range results {
		_, err := db.Exec(`
			INSERT OR REPLACE INTO pokemons (
				id, name, base_exp, height, default_form, weight, species_id,
				sprite_front, sprite_front_shiny,
				sprite_back, sprite_back_shiny,
				sprite_front_female, sprite_front_shiny_female,
				sprite_back_female, sprite_back_shiny_female
			) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			pk.ID, pk.Name, pk.BaseExp, pk.Height, pk.DefaultForm, pk.Weight, pk.SpeciesID,
			pk.SpriteFront, pk.SpriteFrontShiny,
			pk.SpriteBack, pk.SpriteBackShiny,
			pk.SpriteFrontFemale, pk.SpriteFrontShinyFemale,
			pk.SpriteBackFemale, pk.SpriteBackShinyFemale,
		)
		if err != nil {
			log.Println("Insert failed for", pk.Name, err)
			continue
		}
		log.Println("Wrote", pk.Name)
	}

	fmt.Println("Done.")
}
