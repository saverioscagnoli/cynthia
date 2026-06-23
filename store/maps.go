package store

import (
	"cynthia/service/util"
	"database/sql"
	"log/slog"
)

var (
	Pokemons         map[int]*Pokemon        = map[int]*Pokemon{}
	PokemonSpriteMap map[int]*PokemonSprites = map[int]*PokemonSprites{}
	Types            map[int]*Type           = map[int]*Type{}
	TypeSprites      map[int]*[]byte         = map[int]*[]byte{}
	Moves            map[int]*Move           = map[int]*Move{}
	Stats            map[int]*Stat           = map[int]*Stat{}
	Items            map[int]*Item           = map[int]*Item{}
	ItemSprites      map[int]*[]byte         = map[int]*[]byte{}
	TrainerSprites   map[int]*[]byte         = map[int]*[]byte{}
)

func Extract(path string) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		slog.Error("Failed to open database", "err", err)
		return
	}

	defer db.Close()

	loadStats(db)
	loadTypes(db)
	loadMoves(db)
	loadItems(db)
	loadPokemons(db)
	loadTrainerSprites(db)
}

func loadTypes(db *sql.DB) {
	rows, err := db.Query(`SELECT id, name, sprite FROM types`)
	if err != nil {
		slog.Error("Failed to get types table", "err", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var t Type
		var sprite []byte

		if err := rows.Scan(&t.ID, &t.Name, &sprite); err != nil {
			slog.Error("Scan type", "err", err)
			continue
		}

		Types[t.ID] = &t

		if sprite != nil {
			TypeSprites[t.ID] = &sprite
		}
	}
}

func loadMoves(db *sql.DB) {
	rows, err := db.Query(`
		SELECT id, name, accuracy, effect_chance, base_pp, priority, power,
		       damage_class, type_id, min_hits, max_hits, min_turns, max_turns,
		       drain_or_recoil, healing, crit_rate_bonus, flinch_chance, stat_change_chance
		FROM moves`)
	if err != nil {
		slog.Error("Failed to get moves table", "err", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var m Move
		var typeID int
		var damageClass sql.NullString

		err := rows.Scan(
			&m.ID, &m.Name, &m.Accuracy, &m.EffectChance, &m.BasePP, &m.Priority, &m.Power,
			&damageClass, &typeID, &m.MinHits, &m.MaxHits, &m.MinTurns, &m.MaxTurns,
			&m.DrainOrRecoil, &m.Healing, &m.CritRateBonus, &m.FlinchChance, &m.StatChangeChance,
		)

		if err != nil {
			slog.Error("Scan move", "err", err)
			continue
		}

		if damageClass.Valid {
			dc := MoveDamageClass(damageClass.String)
			m.DamageClass = &dc
		}

		m.Type = *Types[typeID]
		Moves[m.ID] = &m
	}
}

func loadItems(db *sql.DB) {
	rows, err := db.Query(`SELECT id, name, cost, fling_power, fling_effect FROM items`)
	if err != nil {
		slog.Error("Failed to get items table", "err", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var i Item

		if err := rows.Scan(&i.ID, &i.Name, &i.Cost, &i.FlingPower, &i.FlingEffect); err != nil {
			slog.Error("Scan item", "err", err)
			continue
		}

		Items[i.ID] = &i
	}

	loadItemSprites(db)
}

func loadStats(db *sql.DB) {
	rows, err := db.Query(`SELECT id, name FROM stats`)
	if err != nil {
		slog.Error("Failed to get stats table", "err", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var s Stat
		if err := rows.Scan(&s.ID, &s.Name); err != nil {
			slog.Error("Scan stat", "err", err)
			continue
		}
		Stats[s.ID] = &s
	}
}

func loadItemSprites(db *sql.DB) {
	rows, err := db.Query(`SELECT item_id, sprite FROM item_sprites`)
	if err != nil {
		slog.Error("Failed to get item sprites", "err", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var sprite []byte

		if err := rows.Scan(&id, &sprite); err != nil {
			slog.Error("Scan item sprite", "err", err)
			continue
		}

		if sprite != nil {
			ItemSprites[id] = &sprite
		}
	}
}

func loadPokemons(db *sql.DB) {
	rows, err := db.Query(`
		SELECT p.id, p.name, p.base_exp, p.height, p.default_form, p.weight,
		       s.id, s.name, s.dex_number, s.gender_chance, s.capture_rate,
		       s.base_happiness, s.baby, s.legendary, s.mythic, s.hatch_counter,
		       s.gender_differences, s.forms_switchable, s.color
		FROM pokemons p
		JOIN pokemon_species s ON s.id = p.species_id`)
	if err != nil {
		slog.Error("Failed to get pokemon table", "err", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		p := &Pokemon{}
		var color string

		err := rows.Scan(
			&p.ID, &p.Name, &p.BaseExp, &p.Height, &p.DefaultForm, &p.Weight,
			&p.Species.ID, &p.Species.Name, &p.Species.NationalDexNumber,
			&p.Species.GenderChance, &p.Species.CaptureRate, &p.Species.BaseHappiness,
			&p.Species.IsBaby, &p.Species.IsLegendary, &p.Species.IsMythic,
			&p.Species.HatchCounter, &p.Species.GenderDifferences, &p.Species.FormsSwitchable,
			&color,
		)
		if err != nil {
			slog.Error("Scan pokemon", "err", err)
			continue
		}

		p.Species.Color = PokemonColor(color)
		Pokemons[p.ID] = p
	}

	loadPokemonSprites(db)
	loadPokemonTypes(db)
	loadPokemonMoves(db)
	loadPokemonItems(db)
	loadPokemonStats(db)
}

func loadPokemonSprites(db *sql.DB) {
	rows, err := db.Query(`
		SELECT id, sprite_front, sprite_front_shiny, sprite_back, sprite_back_shiny,
		       sprite_front_female, sprite_front_shiny_female,
		       sprite_back_female, sprite_back_shiny_female
		FROM pokemons`)
	if err != nil {
		slog.Error("Failed to get pokemon sprites", "err", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var front, frontShiny, back, backShiny []byte
		var frontFemale, frontShinyFemale, backFemale, backShinyFemale []byte

		err := rows.Scan(
			&id,
			&front, &frontShiny, &back, &backShiny,
			&frontFemale, &frontShinyFemale, &backFemale, &backShinyFemale,
		)
		if err != nil {
			slog.Error("Scan pokemon sprites", "err", err)
			continue
		}

		PokemonSpriteMap[id] = &PokemonSprites{
			Front:            util.Ptr(front),
			FrontShiny:       util.Ptr(frontShiny),
			Back:             util.Ptr(back),
			BackShiny:        util.Ptr(backShiny),
			FrontFemale:      util.Ptr(frontFemale),
			FrontShinyFemale: util.Ptr(frontShinyFemale),
			BackFemale:       util.Ptr(backFemale),
			BackShinyFemale:  util.Ptr(backShinyFemale),
		}
	}
}

func loadPokemonTypes(db *sql.DB) {

	rows, err := db.Query(`SELECT pokemon_id, type_id FROM pokemon_types ORDER BY slot`)
	if err != nil {
		slog.Error("Failed to get pokemon_types table", "err", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var pokemonID, typeID int
		if err := rows.Scan(&pokemonID, &typeID); err != nil {
			continue
		}

		p, ok := Pokemons[pokemonID]
		if !ok {
			continue
		}
		p.Types = append(p.Types, *Types[typeID])
	}
}

func loadPokemonMoves(db *sql.DB) {
	rows, err := db.Query(`
		SELECT pokemon_id, move_id, learn_method
		FROM pokemon_moves
	`)
	if err != nil {
		slog.Error("Failed to get pokemon_moves table", "err", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var pokemonID, moveID int
		var learnMethod string

		if err := rows.Scan(&pokemonID, &moveID, &learnMethod); err != nil {
			continue
		}

		p, ok := Pokemons[pokemonID]
		if !ok {
			continue
		}

		move := Moves[moveID]

		pm := &PokemonMove{
			Move:        *move,
			LearnMethod: MoveLearnMethod(learnMethod),
		}

		p.Moves = append(p.Moves, pm)
	}
}

func loadPokemonItems(db *sql.DB) {
	rows, err := db.Query(`SELECT pokemon_id, item_id, rarity FROM held_items`)
	if err != nil {
		slog.Error("Failed to get held_items table", "err", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var pokemonID, itemID, rarity int

		if err := rows.Scan(&pokemonID, &itemID, &rarity); err != nil {
			continue
		}

		p, ok := Pokemons[pokemonID]
		if !ok {
			continue
		}

		p.Items = append(p.Items, &HeldItem{
			Item:   Items[itemID],
			Rarity: rarity,
		})
	}
}

func loadPokemonStats(db *sql.DB) {
	rows, err := db.Query(`SELECT pokemon_id, stat_id, base_stat FROM pokemon_base_stats`)
	if err != nil {
		slog.Error("Failed to get pokemon_base_stats table", "err", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var pokemonID, statID, baseStat int
		if err := rows.Scan(&pokemonID, &statID, &baseStat); err != nil {
			slog.Error("Scan pokemon stat", "err", err)
			continue
		}

		p, ok := Pokemons[pokemonID]
		if !ok {
			continue
		}

		stat := Stats[statID]
		p.BaseStats = append(p.BaseStats, &PokemonStat{
			Stat: Stat{
				ID:   stat.ID,
				Name: stat.Name,
			},
			Value: baseStat,
		})
	}
}

func loadTrainerSprites(db *sql.DB) {
	rows, err := db.Query(`SELECT id, sprite FROM trainer_sprites`)

	if err != nil {
		slog.Error("Failed to get trainer_sprites table", "err", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var sprite []byte

		if err := rows.Scan(&id, &sprite); err != nil {
			slog.Error("Scan trainer sprite", "err", err)
			continue
		}

		if sprite != nil {
			TrainerSprites[id] = &sprite
		}
	}
}
