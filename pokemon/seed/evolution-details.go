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

// Evolution chain list uses unnamed resources (no name, only url)
type ChainRef struct {
	URL string `json:"url"`
}

type ChainList struct {
	Count   int        `json:"count"`
	Next    string     `json:"next"`
	Results []ChainRef `json:"results"`
}

type EvolutionDetail struct {
	SpeciesID          int
	EvolvesToSpeciesID int
	PokemonID          int
	Trigger            string
	Gender             *string
	HeldItemID         *int
	KnownMoveID        *int
	KnownMoveTypeID    *int
	LocationID         *int
	MinLevel           int
	MinHappiness       *int
	MinBeauty          *int
	MinAffection       *int
	NeedsMultiplayer   bool
	NeedsRain          bool
	PartySpeciesID     *int
	PartyTypeID        *int
	RelativePhysStats  *int
	TimeOfDay          *string
	TradeSpeciesID     *int
	TurnUpsideDown     bool
	UsedMoveID         *int
	MinMoveCount       int
	MinSteps           *int
	MinDamageTaken     *int
}

var triggerMap = map[string]string{
	"level-up":               "LevelUp",
	"trade":                  "Trade",
	"use-item":               "UseItem",
	"shed":                   "Shed",
	"spin":                   "Spin",
	"tower-of-darkness":      "TowerOfDarkness",
	"tower-of-waters":        "TowerOfWaters",
	"three-critical-hits":    "ThreeCriticalHits",
	"take-damage":            "TakeDamage",
	"other":                  "Other",
	"agile-style-move":       "AgileStyleMove",
	"strong-style-move":      "StrongStyleMove",
	"recoil-damage":          "RecoilDamage",
	"use-move":               "UseMove",
	"three-defeated-bisharp": "ThreeDefeatedBisharp",
	"gimmighoul-coins":       "GimmighoulCoins",
}

var genderMap = map[float64]string{
	1: "Female",
	2: "Male",
	3: "Genderless",
}

func createEvolutionDetailsTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS evolution_details (
		species_id              INTEGER NOT NULL,
		evolves_to_species_id   INTEGER NOT NULL,
		pokemon_id              INTEGER NOT NULL REFERENCES pokemons(id),
		trigger                 TEXT    NOT NULL,
		gender                  TEXT    CHECK(gender IN ('Male', 'Female', 'Genderless')),
		held_item_id            INTEGER,
		known_move_id           INTEGER,
		known_move_type_id      INTEGER REFERENCES types(id),
		location_id             INTEGER,
		min_level               INTEGER NOT NULL DEFAULT 0,
		min_happiness           INTEGER,
		min_beauty              INTEGER,
		min_affection           INTEGER,
		needs_multiplayer       BOOLEAN NOT NULL DEFAULT 0,
		needs_rain              BOOLEAN NOT NULL DEFAULT 0,
		party_species_id        INTEGER,
		party_type_id           INTEGER REFERENCES types(id),
		relative_physical_stats INTEGER,
		time_of_day             TEXT,
		trade_species_id        INTEGER,
		turn_upside_down        BOOLEAN NOT NULL DEFAULT 0,
		used_move_id            INTEGER REFERENCES moves(id),
		min_move_count          INTEGER NOT NULL DEFAULT 0,
		min_steps               INTEGER,
		min_damage_taken        INTEGER,
		PRIMARY KEY (species_id, evolves_to_species_id, trigger)
	)`)
	if err != nil {
		log.Fatal(err)
	}
}

func idFromURL(url, pattern string) int {
	var id int
	fmt.Sscanf(url, pattern, &id)
	return id
}

func nullableIntFromMap(m map[string]any, key string) *int {
	v, ok := m[key]
	if !ok || v == nil {
		return nil
	}
	f, ok := v.(float64)
	if !ok {
		return nil
	}
	i := int(f)
	return &i
}

func nullableResourceID(m map[string]any, key, pattern string) *int {
	v, ok := m[key]
	if !ok || v == nil {
		return nil
	}
	res, ok := v.(map[string]any)
	if !ok {
		return nil
	}
	u, _ := res["url"].(string)
	id := idFromURL(u, pattern)
	if id == 0 {
		return nil
	}
	return &id
}

func nullableString(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

// walkChain recursively walks the chain link tree and emits EvolutionDetail rows.
// parentSpeciesID is the species_id of the node we evolved FROM.
func walkChain(link map[string]any, parentSpeciesID int, rows *[]EvolutionDetail) {
	species, _ := link["species"].(map[string]any)
	speciesURL, _ := species["url"].(string)
	speciesID := idFromURL(speciesURL, "http://localhost:8000/api/v2/pokemon-species/%d/")

	evolvesTo, _ := link["evolves_to"].([]any)
	evolutionDetails, _ := link["evolution_details"].([]any)

	// Emit one row per evolution_detail entry (there can be multiple trigger conditions)
	if parentSpeciesID != 0 && len(evolutionDetails) > 0 {
		for _, ed := range evolutionDetails {
			edMap, _ := ed.(map[string]any)

			triggerRes, _ := edMap["trigger"].(map[string]any)
			triggerName, _ := triggerRes["name"].(string)
			triggerEnum, ok := triggerMap[triggerName]
			if !ok {
				triggerEnum = "Other"
			}

			// gender: API returns 1=female, 2=male, 3=genderless, null=any
			var gender *string
			if g, ok := edMap["gender"].(float64); ok {
				if gStr, ok := genderMap[g]; ok {
					gender = &gStr
				}
			}

			timeOfDay, _ := edMap["time_of_day"].(string)

			needsOverworldRain, _ := edMap["needs_overworld_rain"].(bool)
			turnUpsideDown, _ := edMap["turn_upside_down"].(bool)

			row := EvolutionDetail{
				SpeciesID:          parentSpeciesID,
				EvolvesToSpeciesID: speciesID,
				PokemonID:          speciesID, // species_id == pokemon_id for base forms
				Trigger:            triggerEnum,
				Gender:             gender,
				HeldItemID:         nullableResourceID(edMap, "held_item", "http://localhost:8000/api/v2/item/%d/"),
				KnownMoveID:        nullableResourceID(edMap, "known_move", "http://localhost:8000/api/v2/move/%d/"),
				KnownMoveTypeID:    nullableResourceID(edMap, "known_move_type", "http://localhost:8000/api/v2/type/%d/"),
				LocationID:         nullableResourceID(edMap, "location", "http://localhost:8000/api/v2/location/%d/"),
				MinLevel: func() int {
					v := nullableIntFromMap(edMap, "min_level")
					if v != nil {
						return *v
					}
					return 0
				}(),
				MinHappiness:      nullableIntFromMap(edMap, "min_happiness"),
				MinBeauty:         nullableIntFromMap(edMap, "min_beauty"),
				MinAffection:      nullableIntFromMap(edMap, "min_affection"),
				NeedsMultiplayer:  false, // not in API; reserved
				NeedsRain:         needsOverworldRain,
				PartySpeciesID:    nullableResourceID(edMap, "party_species", "http://localhost:8000/api/v2/pokemon-species/%d/"),
				PartyTypeID:       nullableResourceID(edMap, "party_type", "http://localhost:8000/api/v2/type/%d/"),
				RelativePhysStats: nullableIntFromMap(edMap, "relative_physical_stats"),
				TimeOfDay:         nullableString(timeOfDay),
				TradeSpeciesID:    nullableResourceID(edMap, "trade_species", "http://localhost:8000/api/v2/pokemon-species/%d/"),
				TurnUpsideDown:    turnUpsideDown,
				UsedMoveID:        nullableResourceID(edMap, "used_move", "http://localhost:8000/api/v2/move/%d/"),
				MinMoveCount:      0,
				MinSteps:          nullableIntFromMap(edMap, "min_steps"),
				MinDamageTaken:    nullableIntFromMap(edMap, "min_damage_taken"),
			}
			*rows = append(*rows, row)
		}
	}

	for _, child := range evolvesTo {
		childMap, _ := child.(map[string]any)
		walkChain(childMap, speciesID, rows)
	}
}

func fetchEvolutionChain(url string) ([]EvolutionDetail, error) {
	r, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("fetching chain: %w", err)
	}
	defer r.Body.Close()

	var raw map[string]any
	if err := json.NewDecoder(r.Body).Decode(&raw); err != nil {
		return nil, fmt.Errorf("decoding chain: %w", err)
	}

	chain, _ := raw["chain"].(map[string]any)

	var rows []EvolutionDetail
	walkChain(chain, 0, &rows)
	return rows, nil
}

func main() {
	db, err := sql.Open("sqlite3", "./pokemon.db?_busy_timeout=5000")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	createEvolutionDetailsTable(db)

	// evolution-chain is an unnamed endpoint (no name field, only url)
	var allChains []ChainRef
	url := "http://localhost:8000/api/v2/evolution-chain?limit=100&offset=0"
	for url != "" {
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		var list ChainList
		json.NewDecoder(resp.Body).Decode(&list)
		resp.Body.Close()
		allChains = append(allChains, list.Results...)
		url = list.Next
		fmt.Printf("Collected %d chains so far...\n", len(allChains))
	}
	fmt.Printf("Total: %d evolution chains to fetch\n", len(allChains))

	const workers = 10
	jobs := make(chan ChainRef, len(allChains))
	results := make(chan []EvolutionDetail, workers)

	var fetchWg sync.WaitGroup
	for i := 0; i < workers; i++ {
		fetchWg.Add(1)
		go func() {
			defer fetchWg.Done()
			for c := range jobs {
				rows, err := fetchEvolutionChain(c.URL)
				if err != nil {
					log.Println("Failed to fetch chain", c.URL, err)
					continue
				}
				if len(rows) > 0 {
					results <- rows
				}
			}
		}()
	}

	go func() {
		fetchWg.Wait()
		close(results)
	}()

	for _, c := range allChains {
		jobs <- c
	}
	close(jobs)

	for rows := range results {
		tx, err := db.Begin()
		if err != nil {
			log.Println("Failed to begin tx:", err)
			continue
		}

		for _, row := range rows {
			_, err := tx.Exec(`
				INSERT OR REPLACE INTO evolution_details (
					species_id, evolves_to_species_id, pokemon_id,
					trigger, gender,
					held_item_id, known_move_id, known_move_type_id, location_id,
					min_level, min_happiness, min_beauty, min_affection,
					needs_multiplayer, needs_rain,
					party_species_id, party_type_id,
					relative_physical_stats, time_of_day, trade_species_id,
					turn_upside_down, used_move_id,
					min_move_count, min_steps, min_damage_taken
				) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
				row.SpeciesID, row.EvolvesToSpeciesID, row.PokemonID,
				row.Trigger, row.Gender,
				row.HeldItemID, row.KnownMoveID, row.KnownMoveTypeID, row.LocationID,
				row.MinLevel, row.MinHappiness, row.MinBeauty, row.MinAffection,
				row.NeedsMultiplayer, row.NeedsRain,
				row.PartySpeciesID, row.PartyTypeID,
				row.RelativePhysStats, row.TimeOfDay, row.TradeSpeciesID,
				row.TurnUpsideDown, row.UsedMoveID,
				row.MinMoveCount, row.MinSteps, row.MinDamageTaken,
			)
			if err != nil {
				log.Printf("Insert failed species %d->%d: %v\n", row.SpeciesID, row.EvolvesToSpeciesID, err)
			}
		}

		if err := tx.Commit(); err != nil {
			log.Println("Commit failed:", err)
		} else {
			log.Printf("Wrote %d evolution rows\n", len(rows))
		}
	}

	fmt.Println("Done.")
}
