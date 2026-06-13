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

type MoveRef struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type MoveList struct {
	Count   int       `json:"count"`
	Next    string    `json:"next"`
	Results []MoveRef `json:"results"`
}

type Move struct {
	ID               int
	Name             string
	Accuracy         *int
	EffectChance     *int
	BasePP           int
	Priority         int
	Power            *int
	DamageClass      *string
	TypeID           int
	MinHits          *int
	MaxHits          *int
	MinTurns         *int
	MaxTurns         *int
	DrainOrRecoil    *int
	Healing          *int
	CritRateBonus    int
	FlinchChance     int
	StatChangeChance int
}

func createMovesTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS moves (
		id                  INTEGER PRIMARY KEY,
		name                TEXT    NOT NULL,
		accuracy            INTEGER,
		effect_chance       INTEGER,
		base_pp             INTEGER NOT NULL,
		priority            INTEGER NOT NULL,
		power               INTEGER,
		damage_class        TEXT CHECK(damage_class IN ('Physical', 'Special')),
		type_id             INTEGER NOT NULL,
		min_hits            INTEGER,
		max_hits            INTEGER,
		min_turns           INTEGER,
		max_turns           INTEGER,
		drain_or_recoil     INTEGER,
		healing             INTEGER,
		crit_rate_bonus     INTEGER NOT NULL,
		flinch_chance       INTEGER NOT NULL,
		stat_change_chance  INTEGER NOT NULL
	)`)
	if err != nil {
		log.Fatal(err)
	}
}

func nullableInt(v any) *int {
	if v == nil {
		return nil
	}
	f, ok := v.(float64)
	if !ok {
		return nil
	}
	i := int(f)
	return &i
}

func fetchMove(url string) (Move, error) {
	r, err := http.Get(url)
	if err != nil {
		return Move{}, fmt.Errorf("fetching move: %w", err)
	}
	defer r.Body.Close()

	var raw map[string]any
	if err := json.NewDecoder(r.Body).Decode(&raw); err != nil {
		return Move{}, fmt.Errorf("decoding move: %w", err)
	}

	// damage_class -> "Physical" or "Special" (capitalize first letter)
	var damageClass *string
	if dc, ok := raw["damage_class"].(map[string]any); ok {
		if name, ok := dc["name"].(string); ok {
			switch name {
			case "physical":
				s := "Physical"
				damageClass = &s
			case "special":
				s := "Special"
				damageClass = &s
			}
		}
	}

	// type -> id from URL e.g. .../type/13/
	typeID := 0
	if t, ok := raw["type"].(map[string]any); ok {
		if tURL, ok := t["url"].(string); ok {
			fmt.Sscanf(tURL[len(tURL)-len("13/"):], "%d", &typeID) // fragile; use regex below
			// safer: scan the last path segment
			var tmp int
			fmt.Sscanf(tURL, "http://localhost:8000/api/v2/type/%d/", &tmp)
			if tmp != 0 {
				typeID = tmp
			}
		}
	}

	// meta
	meta, _ := raw["meta"].(map[string]any)

	pp := 0
	if v, ok := raw["pp"].(float64); ok {
		pp = int(v)
	}
	priority := 0
	if v, ok := raw["priority"].(float64); ok {
		priority = int(v)
	}
	critRate := 0
	flinchChance := 0
	statChance := 0
	if meta != nil {
		if v, ok := meta["crit_rate"].(float64); ok {
			critRate = int(v)
		}
		if v, ok := meta["flinch_chance"].(float64); ok {
			flinchChance = int(v)
		}
		if v, ok := meta["stat_chance"].(float64); ok {
			statChance = int(v)
		}
	}

	var minHits, maxHits, minTurns, maxTurns, drain, healing *int
	if meta != nil {
		minHits = nullableInt(meta["min_hits"])
		maxHits = nullableInt(meta["max_hits"])
		minTurns = nullableInt(meta["min_turns"])
		maxTurns = nullableInt(meta["max_turns"])
		drain = nullableInt(meta["drain"])
		healing = nullableInt(meta["healing"])
	}

	mv := Move{
		ID:               int(raw["id"].(float64)),
		Name:             raw["name"].(string),
		Accuracy:         nullableInt(raw["accuracy"]),
		EffectChance:     nullableInt(raw["effect_chance"]),
		BasePP:           pp,
		Priority:         priority,
		Power:            nullableInt(raw["power"]),
		DamageClass:      damageClass,
		TypeID:           typeID,
		MinHits:          minHits,
		MaxHits:          maxHits,
		MinTurns:         minTurns,
		MaxTurns:         maxTurns,
		DrainOrRecoil:    drain,
		Healing:          healing,
		CritRateBonus:    critRate,
		FlinchChance:     flinchChance,
		StatChangeChance: statChance,
	}

	return mv, nil
}

func main() {
	db, err := sql.Open("sqlite3", "./pokemon.db?_busy_timeout=5000")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	createMovesTable(db)

	// Collect all move URLs across all pages
	var allMoves []MoveRef
	url := "http://localhost:8000/api/v2/move?limit=100&offset=0"
	for url != "" {
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		var list MoveList
		json.NewDecoder(resp.Body).Decode(&list)
		resp.Body.Close()
		allMoves = append(allMoves, list.Results...)
		url = list.Next
		fmt.Printf("Collected %d moves so far...\n", len(allMoves))
	}
	fmt.Printf("Total: %d moves to fetch\n", len(allMoves))

	const workers = 10
	jobs := make(chan MoveRef, len(allMoves))
	results := make(chan Move, workers)

	var fetchWg sync.WaitGroup
	for i := 0; i < workers; i++ {
		fetchWg.Add(1)
		go func() {
			defer fetchWg.Done()
			for m := range jobs {
				mv, err := fetchMove(m.URL)
				if err != nil {
					log.Println("Failed to fetch", m.Name, err)
					continue
				}
				results <- mv
			}
		}()
	}

	go func() {
		fetchWg.Wait()
		close(results)
	}()

	for _, m := range allMoves {
		jobs <- m
	}
	close(jobs)

	for mv := range results {
		_, err := db.Exec(`
			INSERT OR REPLACE INTO moves (
				id, name, accuracy, effect_chance, base_pp, priority, power,
				damage_class, type_id,
				min_hits, max_hits, min_turns, max_turns,
				drain_or_recoil, healing,
				crit_rate_bonus, flinch_chance, stat_change_chance
			) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			mv.ID, mv.Name, mv.Accuracy, mv.EffectChance, mv.BasePP, mv.Priority, mv.Power,
			mv.DamageClass, mv.TypeID,
			mv.MinHits, mv.MaxHits, mv.MinTurns, mv.MaxTurns,
			mv.DrainOrRecoil, mv.Healing,
			mv.CritRateBonus, mv.FlinchChance, mv.StatChangeChance,
		)
		if err != nil {
			log.Println("Insert failed for", mv.Name, err)
			continue
		}
		log.Println("Wrote", mv.Name)
	}

	fmt.Println("Done.")
}
