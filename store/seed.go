//go:build ignore

package main

import (
	"camilla/assets"
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"sort"
	"strings"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

const (
	baseURL = "http://localhost:8000/api/v2"
	workers = 10
)

// ─── Generic list types ───────────────────────────────────────────────────────

type Ref struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type NamelessRef struct {
	URL string `json:"url"`
}

type Page[T any] struct {
	Count   int    `json:"count"`
	Next    string `json:"next"`
	Results []T    `json:"results"`
}

// ─── Domain structs ───────────────────────────────────────────────────────────

type Type struct {
	ID     int
	Name   string
	Sprite []byte
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

type PokemonTypeRow struct {
	PokemonID int
	TypeID    int
	Slot      int
}

type PokemonMoveRow struct {
	PokemonID   int
	MoveID      int
	MinLevel    int
	LearnMethod string
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

type Item struct {
	ID          int
	Name        string
	Cost        int
	FlingPower  *int
	FlingEffect *string
	Sprite      []byte
}

type HeldItem struct {
	ItemID    int
	PokemonID int
	Rarity    int
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

type StatDef struct {
	ID   int
	Name string
}

type StatEntry struct {
	PokemonID int
	StatID    int
	BaseStat  int
}

func intOrZero(v any) int {
	if i := nullInt(v); i != nil {
		return *i
	}
	return 0
}

// ─── Schema ───────────────────────────────────────────────────────────────────

func createSchema(db *sql.DB) {
	stmts := []string{
		`CREATE TABLE IF NOT EXISTS types (
			id     INTEGER PRIMARY KEY,
			name   TEXT NOT NULL,
			sprite BLOB
		)`,
		`CREATE TABLE IF NOT EXISTS stats (
			id   INTEGER PRIMARY KEY,
			name TEXT NOT NULL
		)`,
		`CREATE TABLE IF NOT EXISTS pokemon_species (
			id                 INTEGER PRIMARY KEY,
			name               TEXT    NOT NULL,
			dex_number         INTEGER NOT NULL,
			gender_chance      INTEGER NOT NULL,
			capture_rate       INTEGER NOT NULL,
			base_happiness     INTEGER NOT NULL,
			baby               BOOLEAN NOT NULL,
			legendary          BOOLEAN NOT NULL,
			mythic             BOOLEAN NOT NULL,
			hatch_counter      INTEGER NOT NULL,
			gender_differences BOOLEAN NOT NULL,
			forms_switchable   BOOLEAN NOT NULL,
			color              TEXT    NOT NULL
		)`,
		`CREATE TABLE IF NOT EXISTS pokemons (
			id                        INTEGER PRIMARY KEY,
			name                      TEXT    NOT NULL,
			base_exp                  INTEGER NOT NULL,
			height                    INTEGER NOT NULL,
			default_form              BOOLEAN NOT NULL,
			weight                    INTEGER NOT NULL,
			species_id                INTEGER NOT NULL REFERENCES pokemon_species(id),
			sprite_front              BLOB,
			sprite_front_shiny        BLOB,
			sprite_back               BLOB,
			sprite_back_shiny         BLOB,
			sprite_front_female       BLOB,
			sprite_front_shiny_female BLOB,
			sprite_back_female        BLOB,
			sprite_back_shiny_female  BLOB
		)`,
		`CREATE TABLE IF NOT EXISTS pokemon_base_stats (
			pokemon_id INTEGER NOT NULL REFERENCES pokemons(id),
			stat_id    INTEGER NOT NULL REFERENCES stats(id),
			base_stat  INTEGER NOT NULL,
			PRIMARY KEY (pokemon_id, stat_id)
		)`,
		`CREATE TABLE IF NOT EXISTS items (
    		id           INTEGER PRIMARY KEY,
    		name         TEXT    NOT NULL,
    		cost         INTEGER NOT NULL,
    		fling_power  INTEGER,
    		fling_effect TEXT
      	)`,
		`CREATE TABLE IF NOT EXISTS item_sprites (
       		item_id INTEGER PRIMARY KEY REFERENCES items(id),
       		sprite  BLOB    NOT NULL
       	)`,
		`CREATE TABLE IF NOT EXISTS held_items (
			item_id    INTEGER NOT NULL REFERENCES items(id),
			pokemon_id INTEGER NOT NULL REFERENCES pokemons(id),
			rarity     INTEGER,
			PRIMARY KEY (item_id, pokemon_id)
		)`,
		`CREATE TABLE IF NOT EXISTS moves (
			id                 INTEGER PRIMARY KEY,
			name               TEXT    NOT NULL,
			accuracy           INTEGER,
			effect_chance      INTEGER,
			base_pp            INTEGER NOT NULL,
			priority           INTEGER NOT NULL,
			power              INTEGER,
			damage_class       TEXT CHECK(damage_class IN ('Physical', 'Special')),
			type_id            INTEGER NOT NULL REFERENCES types(id),
			min_hits           INTEGER,
			max_hits           INTEGER,
			min_turns          INTEGER,
			max_turns          INTEGER,
			drain_or_recoil    INTEGER,
			healing            INTEGER,
			crit_rate_bonus    INTEGER NOT NULL,
			flinch_chance      INTEGER NOT NULL,
			stat_change_chance INTEGER NOT NULL
		)`,
		`CREATE TABLE IF NOT EXISTS pokemon_types (
			pokemon_id INTEGER NOT NULL REFERENCES pokemons(id),
			type_id    INTEGER NOT NULL REFERENCES types(id),
			slot       INTEGER NOT NULL,
			PRIMARY KEY (pokemon_id, slot)
		)`,
		`CREATE TABLE IF NOT EXISTS pokemon_moves (
			pokemon_id   INTEGER NOT NULL REFERENCES pokemons(id),
			move_id      INTEGER NOT NULL REFERENCES moves(id),
			min_level    INTEGER NOT NULL,
			learn_method TEXT    NOT NULL CHECK(learn_method IN ('LevelUp','Egg','Tutor','Machine')),
			PRIMARY KEY (pokemon_id, move_id, learn_method)
		)`,
		`CREATE TABLE IF NOT EXISTS evolution_details (
			species_id              INTEGER NOT NULL,
			evolves_to_species_id   INTEGER NOT NULL,
			pokemon_id              INTEGER NOT NULL REFERENCES pokemons(id),
			trigger                 TEXT    NOT NULL,
			gender                  TEXT    CHECK(gender IN ('Male','Female','Genderless')),
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
		)`,
		`CREATE TABLE IF NOT EXISTS trainer_sprites (
           id     INTEGER PRIMARY KEY NOT NULL,
           name   TEXT    NOT NULL,
           sprite BLOB    NOT NULL
        )`,
	}

	for _, s := range stmts {
		if _, err := db.Exec(s); err != nil {
			log.Fatalf("schema error: %v\nstatement: %s", err, s)
		}
	}
}

// ─── Helpers ──────────────────────────────────────────────────────────────────

func getJSON(url string, dst any) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(dst)
}

func fetchBytes(url string) []byte {
	if url == "" {
		return nil
	}
	r, err := http.Get(url)
	if err != nil {
		return nil
	}
	defer r.Body.Close()
	b, _ := io.ReadAll(r.Body)
	return b
}

func nullInt(v any) *int {
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

func nullStr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

// idFromURL extracts the trailing numeric ID from a PokeAPI URL.
// e.g. "http://localhost:8000/api/v2/type/13/" → 13
func idFromURL(url string) int {
	if url == "" {
		return 0
	}
	n := len(url)
	if url[n-1] == '/' {
		url = url[:n-1]
	}
	var id int
	for i := len(url) - 1; i >= 0; i-- {
		if url[i] == '/' {
			fmt.Sscanf(url[i+1:], "%d", &id)
			break
		}
	}
	return id
}

func nullResourceID(m map[string]any, key string) *int {
	v, ok := m[key]
	if !ok || v == nil {
		return nil
	}
	res, ok := v.(map[string]any)
	if !ok {
		return nil
	}
	u, _ := res["url"].(string)
	id := idFromURL(u)
	if id == 0 {
		return nil
	}
	return &id
}

func nullIntFromMap(m map[string]any, key string) *int {
	return nullInt(m[key])
}

func collectPages[T any](startURL string) []T {
	var all []T
	url := startURL
	for url != "" {
		var page Page[T]
		if err := getJSON(url, &page); err != nil {
			log.Fatalf("paginating %s: %v", url, err)
		}
		all = append(all, page.Results...)
		url = page.Next
	}
	return all
}

// fanOut runs fetch concurrently over jobs and sends results to the returned channel.
func fanOut[J any, R any](jobs []J, fetch func(J) (R, error)) <-chan R {
	jobCh := make(chan J, len(jobs))
	resCh := make(chan R, workers)

	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := range jobCh {
				r, err := fetch(j)
				if err != nil {
					log.Println("worker error:", err)
					continue
				}
				resCh <- r
			}
		}()
	}
	go func() { wg.Wait(); close(resCh) }()

	for _, j := range jobs {
		jobCh <- j
	}
	close(jobCh)

	return resCh
}

// ─── Types ────────────────────────────────────────────────────────────────────

func fetchType(ref Ref) (Type, error) {
	var raw map[string]any
	if err := getJSON(ref.URL, &raw); err != nil {
		return Type{}, fmt.Errorf("fetchType %s: %w", ref.Name, err)
	}

	t := Type{
		ID:   int(raw["id"].(float64)),
		Name: raw["name"].(string),
	}

	sprite, err := assets.TypeSprites.ReadFile("types/" + t.Name + ".png")

	if err != nil {
		log.Printf("no sprite for type %s: %v", t.Name, err)
	}
	t.Sprite = sprite

	return t, nil
}

func seedTypes(db *sql.DB) {
	log.Println("=== seeding types ===")
	refs := collectPages[Ref](baseURL + "/type?limit=100&offset=0")
	fmt.Printf("fetching %d types\n", len(refs))

	for t := range fanOut(refs, fetchType) {
		if _, err := db.Exec(
			`INSERT OR REPLACE INTO types (id, name, sprite) VALUES (?, ?, ?)`,
			t.ID, t.Name, t.Sprite,
		); err != nil {
			log.Println("insert type:", err)
			continue
		}
		log.Println("wrote type", t.Name)
	}
}

// ─── Species ──────────────────────────────────────────────────────────────────

func fetchSpecies(ref Ref) (PokemonSpecies, error) {
	var raw map[string]any
	if err := getJSON(ref.URL, &raw); err != nil {
		return PokemonSpecies{}, fmt.Errorf("fetchSpecies %s: %w", ref.Name, err)
	}

	var dexNumber int
	for _, entry := range raw["pokedex_numbers"].([]any) {
		e := entry.(map[string]any)
		if e["pokedex"].(map[string]any)["name"].(string) == "national" {
			dexNumber = int(e["entry_number"].(float64))
			break
		}
	}

	return PokemonSpecies{
		ID:                int(raw["id"].(float64)),
		Name:              raw["name"].(string),
		NationalDexNumber: dexNumber,
		GenderChance:      int(raw["gender_rate"].(float64)),
		CaptureRate:       int(raw["capture_rate"].(float64)),
		BaseHappiness:     int(raw["base_happiness"].(float64)),
		IsBaby:            raw["is_baby"].(bool),
		IsLegendary:       raw["is_legendary"].(bool),
		IsMythic:          raw["is_mythical"].(bool),
		HatchCounter:      int(raw["hatch_counter"].(float64)),
		GenderDifferences: raw["has_gender_differences"].(bool),
		FormsSwitchable:   raw["forms_switchable"].(bool),
		Color:             raw["color"].(map[string]any)["name"].(string),
	}, nil
}

func seedSpecies(db *sql.DB) {
	log.Println("=== seeding pokemon_species ===")
	refs := collectPages[Ref](baseURL + "/pokemon-species?limit=100&offset=0")
	fmt.Printf("fetching %d species\n", len(refs))

	for s := range fanOut(refs, fetchSpecies) {
		if _, err := db.Exec(`
			INSERT OR REPLACE INTO pokemon_species
				(id, name, dex_number, gender_chance, capture_rate, base_happiness,
				 baby, legendary, mythic, hatch_counter, gender_differences, forms_switchable, color)
			VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			s.ID, s.Name, s.NationalDexNumber, s.GenderChance, s.CaptureRate, s.BaseHappiness,
			s.IsBaby, s.IsLegendary, s.IsMythic, s.HatchCounter, s.GenderDifferences, s.FormsSwitchable, s.Color,
		); err != nil {
			log.Println("insert species:", err)
			continue
		}
		log.Println("wrote species", s.Name)
	}
}

// ─── Pokémon ──────────────────────────────────────────────────────────────────

func spriteStr(sprites map[string]any, key string) string {
	v, ok := sprites[key]
	if !ok || v == nil {
		return ""
	}
	s, _ := v.(string)
	return s
}

func fetchPokemon(ref Ref) (Pokemon, error) {
	var raw map[string]any
	if err := getJSON(ref.URL, &raw); err != nil {
		return Pokemon{}, fmt.Errorf("fetchPokemon %s: %w", ref.Name, err)
	}

	sprites, _ := raw["sprites"].(map[string]any)

	// Resolve species ID from the nested species URL
	speciesID := 0
	if sp, ok := raw["species"].(map[string]any); ok {
		if u, ok := sp["url"].(string); ok {
			speciesID = idFromURL(u)
		}
	}

	return Pokemon{
		ID:                     int(raw["id"].(float64)),
		Name:                   raw["name"].(string),
		BaseExp:                intOrZero(raw["base_experience"]),
		Height:                 int(raw["height"].(float64)),
		DefaultForm:            raw["is_default"].(bool),
		Weight:                 int(raw["weight"].(float64)),
		SpeciesID:              speciesID,
		SpriteFront:            fetchBytes(spriteStr(sprites, "front_default")),
		SpriteFrontShiny:       fetchBytes(spriteStr(sprites, "front_shiny")),
		SpriteBack:             fetchBytes(spriteStr(sprites, "back_default")),
		SpriteBackShiny:        fetchBytes(spriteStr(sprites, "back_shiny")),
		SpriteFrontFemale:      fetchBytes(spriteStr(sprites, "front_female")),
		SpriteFrontShinyFemale: fetchBytes(spriteStr(sprites, "front_shiny_female")),
		SpriteBackFemale:       fetchBytes(spriteStr(sprites, "back_female")),
		SpriteBackShinyFemale:  fetchBytes(spriteStr(sprites, "back_shiny_female")),
	}, nil
}

func seedPokemons(db *sql.DB) {
	log.Println("=== seeding pokemons ===")
	refs := collectPages[Ref](baseURL + "/pokemon?limit=100&offset=0")
	fmt.Printf("fetching %d pokemon\n", len(refs))

	for pk := range fanOut(refs, fetchPokemon) {
		if _, err := db.Exec(`
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
		); err != nil {
			log.Println("insert pokemon:", err)
			continue
		}
		log.Println("wrote pokemon", pk.Name, "id", pk.ID)
	}
}

// ─── Stats ────────────────────────────────────────────────────────────────────

type statsResult struct {
	defs    []StatDef
	entries []StatEntry
}

func fetchPokemonStats(ref Ref) (statsResult, error) {
	var raw map[string]any
	if err := getJSON(ref.URL, &raw); err != nil {
		return statsResult{}, fmt.Errorf("fetchPokemonStats %s: %w", ref.Name, err)
	}

	pokemonID := int(raw["id"].(float64))
	statsRaw, _ := raw["stats"].([]any)

	var res statsResult
	for _, s := range statsRaw {
		sm := s.(map[string]any)
		statInfo := sm["stat"].(map[string]any)
		statID := idFromURL(statInfo["url"].(string))

		res.defs = append(res.defs, StatDef{
			ID:   statID,
			Name: statInfo["name"].(string),
		})
		res.entries = append(res.entries, StatEntry{
			PokemonID: pokemonID,
			StatID:    statID,
			BaseStat:  int(sm["base_stat"].(float64)),
		})
	}
	return res, nil
}

func seedStats(db *sql.DB) {
	log.Println("=== seeding stats + pokemon_base_stats ===")
	refs := collectPages[Ref](baseURL + "/pokemon?limit=100&offset=0")
	fmt.Printf("fetching stats for %d pokemon\n", len(refs))

	seenStats := map[int]bool{}
	for res := range fanOut(refs, fetchPokemonStats) {
		for _, def := range res.defs {
			if !seenStats[def.ID] {
				if _, err := db.Exec(
					`INSERT OR IGNORE INTO stats (id, name) VALUES (?, ?)`,
					def.ID, def.Name,
				); err != nil {
					log.Println("insert stat def:", err)
				}
				seenStats[def.ID] = true
			}
		}
		for _, e := range res.entries {
			if _, err := db.Exec(
				`INSERT OR REPLACE INTO pokemon_base_stats (pokemon_id, stat_id, base_stat) VALUES (?, ?, ?)`,
				e.PokemonID, e.StatID, e.BaseStat,
			); err != nil {
				log.Printf("insert base_stat pokemon=%d stat=%d: %v", e.PokemonID, e.StatID, err)
			}
		}
		if len(res.entries) > 0 {
			log.Printf("wrote stats for pokemon_id=%d", res.entries[0].PokemonID)
		}
	}
}

// ─── Pokémon types ────────────────────────────────────────────────────────────

func fetchPokemonTypes(ref Ref) ([]PokemonTypeRow, error) {
	var raw map[string]any
	if err := getJSON(ref.URL, &raw); err != nil {
		return nil, fmt.Errorf("fetchPokemonTypes %s: %w", ref.Name, err)
	}

	pokemonID := int(raw["id"].(float64))
	typesRaw, _ := raw["types"].([]any)

	var rows []PokemonTypeRow
	for _, t := range typesRaw {
		tMap, _ := t.(map[string]any)
		slot := int(tMap["slot"].(float64))
		typeRes, _ := tMap["type"].(map[string]any)
		typeID := idFromURL(typeRes["url"].(string))
		if typeID == 0 {
			continue
		}
		rows = append(rows, PokemonTypeRow{PokemonID: pokemonID, TypeID: typeID, Slot: slot})
	}
	return rows, nil
}

func seedPokemonTypes(db *sql.DB) {
	log.Println("=== seeding pokemon_types ===")
	refs := collectPages[Ref](baseURL + "/pokemon?limit=100&offset=0")
	fmt.Printf("processing types for %d pokemon\n", len(refs))

	for rows := range fanOut(refs, fetchPokemonTypes) {
		if len(rows) == 0 {
			continue
		}
		tx, err := db.Begin()
		if err != nil {
			log.Println("begin tx:", err)
			continue
		}
		for _, row := range rows {
			if _, err := tx.Exec(
				`INSERT OR REPLACE INTO pokemon_types (pokemon_id, type_id, slot) VALUES (?, ?, ?)`,
				row.PokemonID, row.TypeID, row.Slot,
			); err != nil {
				log.Printf("insert pokemon_type pokemon=%d type=%d: %v", row.PokemonID, row.TypeID, err)
			}
		}
		if err := tx.Commit(); err != nil {
			log.Println("commit pokemon_types:", err)
			continue
		}
		log.Printf("wrote %d type rows for pokemon %d", len(rows), rows[0].PokemonID)
	}
}

// ─── Pokémon moves ────────────────────────────────────────────────────────────

var learnMethodMap = map[string]string{
	"level-up": "LevelUp",
	"egg":      "Egg",
	"tutor":    "Tutor",
	"machine":  "Machine",
}

func fetchPokemonMoves(ref Ref) ([]PokemonMoveRow, error) {
	var raw map[string]any
	if err := getJSON(ref.URL, &raw); err != nil {
		return nil, fmt.Errorf("fetchPokemonMoves %s: %w", ref.Name, err)
	}

	pokemonID := int(raw["id"].(float64))
	movesRaw, _ := raw["moves"].([]any)

	type key struct {
		MoveID      int
		LearnMethod string
	}
	best := map[key]int{} // key -> lowest min_level

	for _, m := range movesRaw {
		mMap, _ := m.(map[string]any)
		moveRes, _ := mMap["move"].(map[string]any)
		moveID := idFromURL(moveRes["url"].(string))
		if moveID == 0 {
			continue
		}

		for _, vd := range mMap["version_group_details"].([]any) {
			vdMap, _ := vd.(map[string]any)
			methodName := vdMap["move_learn_method"].(map[string]any)["name"].(string)
			enumVal, ok := learnMethodMap[methodName]
			if !ok {
				continue
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

func seedPokemonMoves(db *sql.DB) {
	log.Println("=== seeding pokemon_moves ===")
	refs := collectPages[Ref](baseURL + "/pokemon?limit=100&offset=0")
	fmt.Printf("processing moves for %d pokemon\n", len(refs))

	for rows := range fanOut(refs, fetchPokemonMoves) {
		if len(rows) == 0 {
			continue
		}
		tx, err := db.Begin()
		if err != nil {
			log.Println("begin tx:", err)
			continue
		}
		for _, row := range rows {
			if _, err := tx.Exec(
				`INSERT OR REPLACE INTO pokemon_moves (pokemon_id, move_id, min_level, learn_method) VALUES (?, ?, ?, ?)`,
				row.PokemonID, row.MoveID, row.MinLevel, row.LearnMethod,
			); err != nil {
				log.Printf("insert pokemon_move pokemon=%d move=%d: %v", row.PokemonID, row.MoveID, err)
			}
		}
		if err := tx.Commit(); err != nil {
			log.Println("commit pokemon_moves:", err)
			continue
		}
		log.Printf("wrote %d move rows for pokemon %d", len(rows), rows[0].PokemonID)
	}
}

// ─── Moves ────────────────────────────────────────────────────────────────────

func fetchMove(ref Ref) (Move, error) {
	var raw map[string]any
	if err := getJSON(ref.URL, &raw); err != nil {
		return Move{}, fmt.Errorf("fetchMove %s: %w", ref.Name, err)
	}

	var damageClass *string
	if dc, ok := raw["damage_class"].(map[string]any); ok {
		switch dc["name"].(string) {
		case "physical":
			s := "Physical"
			damageClass = &s
		case "special":
			s := "Special"
			damageClass = &s
		}
	}

	typeID := 0
	if t, ok := raw["type"].(map[string]any); ok {
		typeID = idFromURL(t["url"].(string))
	}

	meta, _ := raw["meta"].(map[string]any)

	pp := 0
	if v, ok := raw["pp"].(float64); ok {
		pp = int(v)
	}
	priority := 0
	if v, ok := raw["priority"].(float64); ok {
		priority = int(v)
	}

	critRate, flinchChance, statChance := 0, 0, 0
	var minHits, maxHits, minTurns, maxTurns, drain, healing *int
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
		minHits = nullInt(meta["min_hits"])
		maxHits = nullInt(meta["max_hits"])
		minTurns = nullInt(meta["min_turns"])
		maxTurns = nullInt(meta["max_turns"])
		drain = nullInt(meta["drain"])
		healing = nullInt(meta["healing"])
	}

	return Move{
		ID:               int(raw["id"].(float64)),
		Name:             raw["name"].(string),
		Accuracy:         nullInt(raw["accuracy"]),
		EffectChance:     nullInt(raw["effect_chance"]),
		BasePP:           pp,
		Priority:         priority,
		Power:            nullInt(raw["power"]),
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
	}, nil
}

func seedMoves(db *sql.DB) {
	log.Println("=== seeding moves ===")
	refs := collectPages[Ref](baseURL + "/move?limit=100&offset=0")
	fmt.Printf("fetching %d moves\n", len(refs))

	for mv := range fanOut(refs, fetchMove) {
		if _, err := db.Exec(`
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
		); err != nil {
			log.Println("insert move:", err)
			continue
		}
		log.Println("wrote move", mv.Name)
	}
}

// ─── Items ────────────────────────────────────────────────────────────────────

func fetchItem(ref Ref) (Item, error) {
	var raw map[string]any
	if err := getJSON(ref.URL, &raw); err != nil {
		return Item{}, fmt.Errorf("fetchItem %s: %w", ref.Name, err)
	}

	item := Item{
		ID:   int(raw["id"].(float64)),
		Name: raw["name"].(string),
		Cost: int(raw["cost"].(float64)),
	}

	if fp, ok := raw["fling_power"].(float64); ok {
		v := int(fp)
		item.FlingPower = &v
	}

	// fling effect name via sub-request
	if fe, ok := raw["fling_effect"].(map[string]any); ok {
		if u, ok := fe["url"].(string); ok {
			var feRaw struct{ Name string }
			if err := getJSON(u, &feRaw); err == nil && feRaw.Name != "" {
				item.FlingEffect = &feRaw.Name
			}
		}
	}

	if sprites, ok := raw["sprites"].(map[string]any); ok {
		if url, ok := sprites["default"].(string); ok {
			item.Sprite = fetchBytes(url)
		}
	}

	return item, nil
}

func seedItems(db *sql.DB) {
	log.Println("=== seeding items + item_sprites ===")
	refs := collectPages[Ref](baseURL + "/item?limit=100&offset=0")
	fmt.Printf("fetching %d items\n", len(refs))

	for it := range fanOut(refs, fetchItem) {
		if _, err := db.Exec(
			`INSERT OR REPLACE INTO items (id, name, cost, fling_power, fling_effect) VALUES (?, ?, ?, ?, ?)`,
			it.ID, it.Name, it.Cost, it.FlingPower, it.FlingEffect,
		); err != nil {
			log.Println("insert item:", err)
			continue
		}
		if it.Sprite != nil {
			if _, err := db.Exec(
				`INSERT OR REPLACE INTO item_sprites (item_id, sprite) VALUES (?, ?)`,
				it.ID, it.Sprite,
			); err != nil {
				log.Println("insert item_sprite:", err)
			}
		}
		log.Println("wrote item", it.Name)
	}
}

// ─── Held items ───────────────────────────────────────────────────────────────

func fetchHeldItems(ref Ref) ([]HeldItem, error) {
	var raw map[string]any
	if err := getJSON(ref.URL, &raw); err != nil {
		return nil, fmt.Errorf("fetchHeldItems %s: %w", ref.Name, err)
	}

	pokemonID := int(raw["id"].(float64))
	heldRaw, _ := raw["held_items"].([]any)

	var rows []HeldItem
	for _, h := range heldRaw {
		hMap := h.(map[string]any)
		itemRes := hMap["item"].(map[string]any)
		itemID := idFromURL(itemRes["url"].(string))

		versionDetails := hMap["version_details"].([]any)

		// Pick the last non-"xd" version entry
		var lastVD map[string]any
		for i := len(versionDetails) - 1; i >= 0; i-- {
			candidate := versionDetails[i].(map[string]any)
			if candidate["version"].(map[string]any)["name"].(string) != "xd" {
				lastVD = candidate
				break
			}
		}
		if lastVD == nil {
			continue
		}

		rarity := int(lastVD["rarity"].(float64))
		rows = append(rows, HeldItem{ItemID: itemID, PokemonID: pokemonID, Rarity: rarity})
	}
	return rows, nil
}

func seedHeldItems(db *sql.DB) {
	log.Println("=== seeding held_items ===")
	refs := collectPages[Ref](baseURL + "/pokemon?limit=100&offset=0")
	fmt.Printf("processing held items for %d pokemon\n", len(refs))

	for rows := range fanOut(refs, fetchHeldItems) {
		for _, row := range rows {
			if _, err := db.Exec(
				`INSERT OR REPLACE INTO held_items (item_id, pokemon_id, rarity) VALUES (?, ?, ?)`,
				row.ItemID, row.PokemonID, row.Rarity,
			); err != nil {
				log.Printf("insert held_item item=%d pokemon=%d: %v", row.ItemID, row.PokemonID, err)
				continue
			}
			log.Printf("wrote held item %d for pokemon %d", row.ItemID, row.PokemonID)
		}
	}
}

// ─── Evolution details ────────────────────────────────────────────────────────

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

func walkChain(link map[string]any, parentSpeciesID int, rows *[]EvolutionDetail) {
	species, _ := link["species"].(map[string]any)
	speciesID := idFromURL(species["url"].(string))

	evolvesTo, _ := link["evolves_to"].([]any)
	details, _ := link["evolution_details"].([]any)

	if parentSpeciesID != 0 {
		for _, ed := range details {
			edMap, _ := ed.(map[string]any)

			triggerRes, _ := edMap["trigger"].(map[string]any)
			triggerEnum, ok := triggerMap[triggerRes["name"].(string)]
			if !ok {
				triggerEnum = "Other"
			}

			var gender *string
			if g, ok := edMap["gender"].(float64); ok {
				if gStr, ok := genderMap[g]; ok {
					gender = &gStr
				}
			}

			timeOfDay, _ := edMap["time_of_day"].(string)
			needsRain, _ := edMap["needs_overworld_rain"].(bool)
			turnUpsideDown, _ := edMap["turn_upside_down"].(bool)

			minLevel := 0
			if v := nullIntFromMap(edMap, "min_level"); v != nil {
				minLevel = *v
			}

			*rows = append(*rows, EvolutionDetail{
				SpeciesID:          parentSpeciesID,
				EvolvesToSpeciesID: speciesID,
				PokemonID:          speciesID,
				Trigger:            triggerEnum,
				Gender:             gender,
				HeldItemID:         nullResourceID(edMap, "held_item"),
				KnownMoveID:        nullResourceID(edMap, "known_move"),
				KnownMoveTypeID:    nullResourceID(edMap, "known_move_type"),
				LocationID:         nullResourceID(edMap, "location"),
				MinLevel:           minLevel,
				MinHappiness:       nullIntFromMap(edMap, "min_happiness"),
				MinBeauty:          nullIntFromMap(edMap, "min_beauty"),
				MinAffection:       nullIntFromMap(edMap, "min_affection"),
				NeedsMultiplayer:   false,
				NeedsRain:          needsRain,
				PartySpeciesID:     nullResourceID(edMap, "party_species"),
				PartyTypeID:        nullResourceID(edMap, "party_type"),
				RelativePhysStats:  nullIntFromMap(edMap, "relative_physical_stats"),
				TimeOfDay:          nullStr(timeOfDay),
				TradeSpeciesID:     nullResourceID(edMap, "trade_species"),
				TurnUpsideDown:     turnUpsideDown,
				UsedMoveID:         nullResourceID(edMap, "used_move"),
				MinMoveCount:       0,
				MinSteps:           nullIntFromMap(edMap, "min_steps"),
				MinDamageTaken:     nullIntFromMap(edMap, "min_damage_taken"),
			})
		}
	}

	for _, child := range evolvesTo {
		walkChain(child.(map[string]any), speciesID, rows)
	}
}

func fetchEvolutionChain(ref NamelessRef) ([]EvolutionDetail, error) {
	var raw map[string]any
	if err := getJSON(ref.URL, &raw); err != nil {
		return nil, fmt.Errorf("fetchEvolutionChain %s: %w", ref.URL, err)
	}
	chain, _ := raw["chain"].(map[string]any)
	var rows []EvolutionDetail
	walkChain(chain, 0, &rows)
	return rows, nil
}

func seedEvolutionDetails(db *sql.DB) {
	log.Println("=== seeding evolution_details ===")
	refs := collectPages[NamelessRef](baseURL + "/evolution-chain?limit=100&offset=0")
	fmt.Printf("fetching %d evolution chains\n", len(refs))

	for rows := range fanOut(refs, fetchEvolutionChain) {
		if len(rows) == 0 {
			continue
		}
		tx, err := db.Begin()
		if err != nil {
			log.Println("begin tx:", err)
			continue
		}
		for _, row := range rows {
			if _, err := tx.Exec(`
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
			); err != nil {
				log.Printf("insert evolution %d->%d: %v", row.SpeciesID, row.EvolvesToSpeciesID, err)
			}
		}
		if err := tx.Commit(); err != nil {
			log.Println("commit evolution_details:", err)
			continue
		}
		log.Printf("wrote %d evolution rows", len(rows))
	}
}

func seedTrainerSprites(db *sql.DB) {
	log.Println("=== seeding trainer_sprites ===")
	entries, err := assets.TrainerSprites.ReadDir("trainers")
	if err != nil {
		log.Fatal("read trainer sprites dir:", err)
	}

	i := 0

	sort.Slice(entries, func(i, j int) bool {
		ni := strings.TrimSuffix(entries[i].Name(), ".png")
		nj := strings.TrimSuffix(entries[j].Name(), ".png")

		basei := strings.SplitN(ni, "-", 2)[0]
		basej := strings.SplitN(nj, "-", 2)[0]

		if basei != basej {
			return basei < basej
		}

		hasi := strings.Contains(ni, "-")
		hasj := strings.Contains(nj, "-")

		if hasi != hasj {
			return !hasi
		}

		return ni < nj
	})

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		name := entry.Name()
		sprite, err := assets.TrainerSprites.ReadFile("trainers/" + name)
		if err != nil {
			log.Println("read sprite:", name, err)
			continue
		}

		if _, err := db.Exec(
			`INSERT OR REPLACE INTO trainer_sprites (id, name, sprite) VALUES (?, ?, ?)`,
			i, strings.TrimSuffix(name, ".png"), sprite,
		); err != nil {
			log.Println("insert trainer sprite:", err)
			continue
		}

		log.Println("wrote trainer sprite", i, name)
		i++
	}
}

// ─── Main ─────────────────────────────────────────────────────────────────────

func main() {
	path := flag.String("db path", "./assets/pokemon.db", "dp path")
	flag.Parse()

	db, err := sql.Open("sqlite3", *path+"?_busy_timeout=5000")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if _, err := db.Exec("PRAGMA journal_mode=WAL"); err != nil {
		log.Fatal("WAL mode:", err)
	}
	if _, err := db.Exec("PRAGMA foreign_keys=ON"); err != nil {
		log.Fatal("foreign keys:", err)
	}

	createSchema(db)

	// Order respects FK dependencies:
	//   types → (moves, pokemon_types, evolution_details)
	//   pokemon_species → pokemons
	//   pokemons → (pokemon_base_stats, pokemon_types, pokemon_moves, held_items, evolution_details)
	//   items → held_items
	//   moves → (pokemon_moves, evolution_details)
	//seedTypes(db)
	//seedSpecies(db)
	//seedPokemons(db)
	//seedStats(db)
	//seedMoves(db)
	//seedItems(db)
	//seedPokemonTypes(db)
	//seedPokemonMoves(db)
	//seedHeldItems(db)
	//seedEvolutionDetails(db)
	seedTrainerSprites(db)

	fmt.Println("All done.")
}
