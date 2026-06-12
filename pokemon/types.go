package main

import (
	"cynthia/util"
	"database/sql"
	"log"
)

var (
	Pokemons map[int]Pokemon = map[int]Pokemon{}
	Types    map[int]Type    = map[int]Type{}
	Moves    map[int]Move    = map[int]Move{}
	Items    map[int]Item    = map[int]Item{}
)

func Init(path string) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	loadTypes(db)
	loadMoves(db)
	loadItems(db)
	loadPokemons(db)
}

func loadTypes(db *sql.DB) {
	rows, err := db.Query(`SELECT id, name, sprite FROM types`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var t Type
		var sprite []byte
		if err := rows.Scan(&t.ID, &t.Name, &sprite); err != nil {
			log.Println("scan type:", err)
			continue
		}
		if sprite != nil {
			t.Sprite = &sprite
		}
		Types[t.ID] = t
	}
}

func loadMoves(db *sql.DB) {
	rows, err := db.Query(`
        SELECT id, name, accuracy, effect_chance, base_pp, priority, power,
               damage_class, type_id, min_hits, max_hits, min_turns, max_turns,
               drain_or_recoil, healing, crit_rate_bonus, flinch_chance, stat_change_chance
        FROM moves`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var m Move
		var typeID int
		var damageClass *string
		err := rows.Scan(
			&m.ID, &m.Name, &m.Accuracy, &m.EffectChance, &m.BasePP, &m.Priority, &m.Power,
			&damageClass, &typeID, &m.MinHits, &m.MaxHits, &m.MinTurns, &m.MaxTurns,
			&m.DrainOrRecoil, &m.Healing, &m.CritRateBonus, &m.FlinchChance, &m.StatChangeChance,
		)
		if err != nil {
			log.Println("scan move:", err)
			continue
		}
		if damageClass != nil {
			dc := MoveDamageClass(*damageClass)
			m.DamageClass = &dc
		}
		m.Type = Types[typeID]
		Moves[m.ID] = m
	}
}

func loadItems(db *sql.DB) {
	rows, err := db.Query(`SELECT id, name, cost, fling_power, fling_effect FROM items`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var i Item
		if err := rows.Scan(&i.ID, &i.Name, &i.Cost, &i.FlingPower, &i.FlingEffect); err != nil {
			log.Println("scan item:", err)
			continue
		}
		Items[i.ID] = i
	}
}

func loadPokemons(db *sql.DB) {
	// load base pokemon
	rows, err := db.Query(`
        SELECT p.id, p.name, p.base_exp, p.height, p.default_form, p.weight,
               s.id, s.name, s.dex_number, s.gender_chance, s.capture_rate,
               s.base_happiness, s.baby, s.legendary, s.mythic, s.hatch_counter,
               s.gender_differences, s.forms_switchable, s.color,
               p.sprite_front, p.sprite_front_shiny, p.sprite_back, p.sprite_back_shiny,
               p.sprite_front_female, p.sprite_front_shiny_female,
               p.sprite_back_female, p.sprite_back_shiny_female
        FROM pokemons p
        JOIN pokemon_species s ON s.id = p.species_id`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var p Pokemon
		var color string
		var front, frontShiny, back, backShiny []byte
		var frontFemale, frontShinyFemale, backFemale, backShinyFemale []byte

		err := rows.Scan(
			&p.ID, &p.Name, &p.BaseExp, &p.Height, &p.DefaultForm, &p.Weight,
			&p.Species.ID, &p.Species.Name, &p.Species.NationalDexNumber,
			&p.Species.GenderChance, &p.Species.CaptureRate, &p.Species.BaseHappiness,
			&p.Species.IsBaby, &p.Species.IsLegendary, &p.Species.IsMythic,
			&p.Species.HatchCounter, &p.Species.GenderDifferences, &p.Species.FormsSwitchable,
			&color,
			&front, &frontShiny, &back, &backShiny,
			&frontFemale, &frontShinyFemale, &backFemale, &backShinyFemale,
		)
		if err != nil {
			log.Println("scan pokemon:", err)
			continue
		}

		p.Species.Color = PokemonColor(color)
		p.Sprites = PokemonSprites{
			Front:            toPtr(front),
			FrontShiny:       toPtr(frontShiny),
			Back:             toPtr(back),
			BackShiny:        toPtr(backShiny),
			FrontFemale:      toPtr(frontFemale),
			FrontShinyFemale: toPtr(frontShinyFemale),
			BackFemale:       toPtr(backFemale),
			BackShinyFemale:  toPtr(backShinyFemale),
		}

		Pokemons[p.ID] = p
	}

	loadPokemonTypes(db)
	loadPokemonMoves(db)
	loadPokemonItems(db)
}

func loadPokemonTypes(db *sql.DB) {
	rows, err := db.Query(`SELECT pokemon_id, type_id FROM pokemon_types ORDER BY slot`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var pokemonID, typeID int
		if err := rows.Scan(&pokemonID, &typeID); err != nil {
			continue
		}
		p := Pokemons[pokemonID]
		p.Types = append(p.Types, Types[typeID])
		Pokemons[pokemonID] = p
	}
}

func loadPokemonMoves(db *sql.DB) {
	rows, err := db.Query(`SELECT pokemon_id, move_id, learn_method FROM pokemon_moves`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var pokemonID, moveID int
		var learnMethod string
		if err := rows.Scan(&pokemonID, &moveID, &learnMethod); err != nil {
			continue
		}
		lm := MoveLearnMethod(learnMethod)
		p := Pokemons[pokemonID]
		p.Moves = append(p.Moves, &lm)
		Pokemons[pokemonID] = p
	}
}

func loadPokemonItems(db *sql.DB) {
	rows, err := db.Query(`SELECT pokemon_id, item_id, rarity FROM held_items`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var pokemonID, itemID, rarity int
		if err := rows.Scan(&pokemonID, &itemID, &rarity); err != nil {
			continue
		}
		p := Pokemons[pokemonID]
		p.Items = append(p.Items, &HeldItem{Item: Items[itemID], Rarity: rarity})
		Pokemons[pokemonID] = p
	}
}

func toPtr(b []byte) *[]byte {
	if b == nil {
		return nil
	}
	return &b
}

type Type struct {
	ID     int
	Name   string
	Sprite *[]byte
}

type MoveDamageClass string

const (
	MoveDamageClassPhysical MoveDamageClass = "Physical"
	MoveDamageClassSpecial  MoveDamageClass = "Special"
)

func (m *MoveDamageClass) UnmarshalJSON(b []byte) error {
	return util.UnmarshalString(b, m)
}

type Move struct {
	ID               int
	Name             string
	Accuracy         *int
	EffectChance     *int
	BasePP           int
	Priority         int
	Power            *int
	DamageClass      *MoveDamageClass
	Type             Type
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
}

type HeldItem struct {
	Item
	Rarity int
}

type MoveLearnMethod string

const (
	MoveLearnMethodLevelUp MoveLearnMethod = "LevelUp"
	MoveLearnMethodEgg     MoveLearnMethod = "Egg"
	MoveLearnMethodTutor   MoveLearnMethod = "Tutor"
	MoveLearnMethodMachine MoveLearnMethod = "Machine"
)

type PokemonMove struct {
	Move
	LearnMethod MoveLearnMethod
}

type PokemonColor string

const (
	PokemonColorRed    PokemonColor = "red"
	PokemonColorBlue   PokemonColor = "blue"
	PokemonColorYellow PokemonColor = "yellow"
	PokemonColorGreen  PokemonColor = "green"
	PokemonColorBlack  PokemonColor = "black"
	PokemonColorBrown  PokemonColor = "brown"
	PokemonColorPurple PokemonColor = "purple"
	PokemonColorGray   PokemonColor = "gray"
	PokemonColorWhite  PokemonColor = "white"
	PokemonColorPink   PokemonColor = "pink"
)

func (c *PokemonColor) UnmarshalJSON(b []byte) error {
	return util.UnmarshalString(b, c)
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
	Color             PokemonColor
}

type PokemonSprites struct {
	Front            *[]byte
	FrontShiny       *[]byte
	Back             *[]byte
	BackShiny        *[]byte
	FrontFemale      *[]byte
	FrontShinyFemale *[]byte
	BackFemale       *[]byte
	BackShinyFemale  *[]byte
}

type Pokemon struct {
	ID          int
	Name        string
	Types       []Type
	BaseExp     int
	Height      int
	DefaultForm bool
	Weight      int
	Species     PokemonSpecies
	Moves       []*MoveLearnMethod
	Items       []*HeldItem
	Sprites     PokemonSprites
}
