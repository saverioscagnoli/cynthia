package models

import (
	"context"
	"cynthia/cmd/pkapi/store"
	"cynthia/ds"
	"cynthia/pokemon"
	"fmt"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Trainer struct {
	ID        ds.Snowflake             `json:"id"`
	Money     int                      `json:"money"`
	Pokemons  *[]*pokemon.OwnedPokemon `json:"pokemons"`
	Bag       *[]*pokemon.Item         `json:"bag"`
	SpriteID  int                      `json:"sprite_id"`
	CreatedAt string                   `json:"created_at"`
}

func CreateTrainer(ctx context.Context, pool *pgxpool.Pool, id string) (*Trainer, error) {
	_, err := pool.Exec(ctx, `
		INSERT INTO trainers (id, money) VALUES ($1, 0)
		ON CONFLICT (id) DO NOTHING`, id,
	)

	if err != nil {
		return nil, fmt.Errorf("create trainer: %w", err)
	}

	return &Trainer{
		ID:       id,
		Money:    0,
		Pokemons: &[]*pokemon.OwnedPokemon{},
		Bag:      &[]*pokemon.Item{},
	}, nil
}

func GetTrainer(ctx context.Context, pool *pgxpool.Pool, id string) (*Trainer, error) {
	start := time.Now()
	trainer := &Trainer{}

	err := pool.QueryRow(ctx, `
		SELECT id, money FROM trainers WHERE id = $1`, id,
	).Scan(&trainer.ID, &trainer.Money)

	if err != nil {
		return nil, fmt.Errorf("get trainer: %w", err)
	}

	rows, err := pool.Query(ctx, `
		SELECT id, pokemon_id, nickname, level, exp, held_item_id
		FROM owned_pokemons
		WHERE trainer_id = $1`, id,
	)

	if err != nil {
		return nil, fmt.Errorf("get owned pokemons: %w", err)
	}

	ownedByID := map[int]*pokemon.OwnedPokemon{}
	var ownedPokemons []*pokemon.OwnedPokemon

	for rows.Next() {
		op := &pokemon.OwnedPokemon{}
		var heldItemID *int
		var nickname *string

		if err := rows.Scan(&op.ID, &op.SpeciesID, &nickname, &op.Level, &op.Exp, &heldItemID); err != nil {
			return nil, fmt.Errorf("scan owned pokemon: %w", err)
		}

		if nickname != nil {
			op.Name = *nickname
		}
		if heldItemID != nil {
			item := store.Items[*heldItemID]
			op.HeldItem = &item
		}

		ownedByID[op.ID] = op
		ownedPokemons = append(ownedPokemons, op)
	}

	rows.Close()

	if len(ownedPokemons) > 0 {
		ids := make([]int, 0, len(ownedByID))
		for id := range ownedByID {
			ids = append(ids, id)
		}

		moveRows, err := pool.Query(ctx, `
			SELECT owned_pokemon_id, move_id
			FROM owned_pokemon_moves
			WHERE owned_pokemon_id = ANY($1)
			ORDER BY owned_pokemon_id, slot`, ids,
		)

		if err != nil {
			return nil, fmt.Errorf("get owned pokemon moves: %w", err)
		}

		for moveRows.Next() {
			var ownedPokemonID, moveID int

			if err := moveRows.Scan(&ownedPokemonID, &moveID); err != nil {
				return nil, fmt.Errorf("scan owned pokemon move: %w", err)
			}

			move := store.Moves[moveID]
			ownedByID[ownedPokemonID].Moves = append(ownedByID[ownedPokemonID].Moves, move)
		}

		moveRows.Close()

		statRows, err := pool.Query(ctx, `
			SELECT owned_pokemon_id, stat_id, value
			FROM owned_pokemon_stats
			WHERE owned_pokemon_id = ANY($1)`, ids,
		)

		if err != nil {
			return nil, fmt.Errorf("get owned pokemon stats: %w", err)
		}

		for statRows.Next() {
			var ownedPokemonID, statID, value int

			if err := statRows.Scan(&ownedPokemonID, &statID, &value); err != nil {
				return nil, fmt.Errorf("scan owned pokemon stat: %w", err)
			}

			stat := store.Stats[statID]

			ownedByID[ownedPokemonID].Stats = append(ownedByID[ownedPokemonID].Stats, &pokemon.PokemonStat{
				Stat:  *stat,
				Value: value,
			})
		}

		statRows.Close()
	}

	trainer.Pokemons = &ownedPokemons

	bagRows, err := pool.Query(ctx, `
		SELECT item_id, quantity FROM bag WHERE trainer_id = $1`, id,
	)

	if err != nil {
		return nil, fmt.Errorf("get bag: %w", err)
	}

	defer bagRows.Close()

	var bag []*pokemon.Item

	for bagRows.Next() {
		var itemID, quantity int

		if err := bagRows.Scan(&itemID, &quantity); err != nil {
			return nil, fmt.Errorf("scan bag: %w", err)
		}

		for range quantity {
			item := store.Items[itemID]
			bag = append(bag, &item)
		}
	}

	trainer.Bag = &bag

	slog.Debug("Trainer query", "time", time.Since(start))

	return trainer, nil
}

func GetOrCreateTrainer(ctx context.Context, pool *pgxpool.Pool, id ds.Snowflake) (*Trainer, error) {
	_, err := pool.Exec(ctx, `
		INSERT INTO trainers (id, money) VALUES ($1, 0)
		ON CONFLICT (id) DO NOTHING`, id,
	)

	if err != nil {
		return nil, fmt.Errorf("upsert trainer: %w", err)
	}

	return GetTrainer(ctx, pool, id)
}
