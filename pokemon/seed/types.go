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

type TypeRef struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type TypeList struct {
	Count   int       `json:"count"`
	Next    string    `json:"next"`
	Results []TypeRef `json:"results"`
}

type Type struct {
	ID   int
	Name string
}

func createTypesTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS types (
		id   INTEGER PRIMARY KEY,
		name TEXT    NOT NULL
	)`)
	if err != nil {
		log.Fatal(err)
	}
}

func fetchType(url string) (Type, error) {
	r, err := http.Get(url)
	if err != nil {
		return Type{}, fmt.Errorf("fetching type: %w", err)
	}
	defer r.Body.Close()

	var raw map[string]any
	if err := json.NewDecoder(r.Body).Decode(&raw); err != nil {
		return Type{}, fmt.Errorf("decoding type: %w", err)
	}

	return Type{
		ID:   int(raw["id"].(float64)),
		Name: raw["name"].(string),
	}, nil
}

func main() {
	db, err := sql.Open("sqlite3", "./pokemon.db?_busy_timeout=5000")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	createTypesTable(db)

	var allTypes []TypeRef
	url := "http://localhost:8000/api/v2/type?limit=100&offset=0"
	for url != "" {
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		var list TypeList
		json.NewDecoder(resp.Body).Decode(&list)
		resp.Body.Close()
		allTypes = append(allTypes, list.Results...)
		url = list.Next
		fmt.Printf("Collected %d types so far...\n", len(allTypes))
	}
	fmt.Printf("Total: %d types to fetch\n", len(allTypes))

	const workers = 10
	jobs := make(chan TypeRef, len(allTypes))
	results := make(chan Type, workers)

	var fetchWg sync.WaitGroup
	for i := 0; i < workers; i++ {
		fetchWg.Add(1)
		go func() {
			defer fetchWg.Done()
			for t := range jobs {
				typ, err := fetchType(t.URL)
				if err != nil {
					log.Println("Failed to fetch", t.Name, err)
					continue
				}
				results <- typ
			}
		}()
	}

	go func() {
		fetchWg.Wait()
		close(results)
	}()

	for _, t := range allTypes {
		jobs <- t
	}
	close(jobs)

	for typ := range results {
		_, err := db.Exec(`
			INSERT OR REPLACE INTO types (id, name)
			VALUES (?, ?)`,
			typ.ID, typ.Name,
		)
		if err != nil {
			log.Println("Insert failed for", typ.Name, err)
			continue
		}
		log.Println("Wrote", typ.Name)
	}

	fmt.Println("Done.")
}
