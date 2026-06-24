package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func createSchema(ctx context.Context, pool *pgxpool.Pool) error {
	stmts := []string{
		`CREATE TABLE IF NOT EXISTS users (
            id               TEXT      PRIMARY KEY NOT NULL,
            username         TEXT      NOT NULL,
            discord_username TEXT      NOT NULL,
            avatar_hash      TEXT,
            money            INTEGER   NOT NULL DEFAULT 0,
            sprite_id        INTEGER,
            banner           BYTEA,
            banner_type      TEXT,
            created_at       TIMESTAMP NOT NULL DEFAULT NOW()
        )`,
		`CREATE TABLE IF NOT EXISTS bag (
			user_id TEXT    NOT NULL REFERENCES users(id),
			item_id    INTEGER NOT NULL,
			quantity   INTEGER NOT NULL DEFAULT 1 CHECK(quantity > 0),
			PRIMARY KEY (user_id, item_id)
		)`,
		`CREATE TABLE IF NOT EXISTS matches (
    		id            SERIAL    PRIMARY KEY,
    		player1_id    TEXT      NOT NULL REFERENCES users(id),
    		player2_id    TEXT      NOT NULL REFERENCES users(id),
    		winner_id     TEXT      REFERENCES users(id),
    		type          TEXT      NOT NULL DEFAULT 'single' CHECK(type IN ('single', 'double')),
    		status        TEXT      NOT NULL DEFAULT 'pending' CHECK(status IN ('pending', 'active', 'finished', 'cancelled')),
    		created_at    TIMESTAMP NOT NULL DEFAULT NOW(),
    		finished_at   TIMESTAMP
      	)`,
		`CREATE TABLE IF NOT EXISTS match_summaries (
           match_id             INTEGER NOT NULL REFERENCES matches(id),
           user_id              TEXT    NOT NULL REFERENCES users(id),
           damage_dealt         INTEGER NOT NULL DEFAULT 0,
           damage_received      INTEGER NOT NULL DEFAULT 0,
           pokemon_fainted      INTEGER NOT NULL DEFAULT 0,
           pokemon_lost         INTEGER NOT NULL DEFAULT 0,
           PRIMARY KEY (match_id, user_id)
        )`,
		`CREATE TABLE IF NOT EXISTS owned_pokemons (
			id           SERIAL    PRIMARY KEY,
			user_id   TEXT      NOT NULL REFERENCES users(id),
			pokemon_id   INTEGER   NOT NULL,
			nickname     TEXT,
			level        INTEGER   NOT NULL DEFAULT 1,
			exp          INTEGER   NOT NULL DEFAULT 0,
			held_item_id INTEGER,
			created_at   TIMESTAMP NOT NULL DEFAULT NOW()
		)`,
		`CREATE TABLE IF NOT EXISTS owned_pokemon_moves (
			owned_pokemon_id INTEGER NOT NULL REFERENCES owned_pokemons(id),
			move_id          INTEGER NOT NULL,
			slot             INTEGER NOT NULL CHECK(slot BETWEEN 1 AND 4),
			PRIMARY KEY (owned_pokemon_id, slot)
		)`,
		`CREATE TABLE IF NOT EXISTS owned_pokemon_stats (
			owned_pokemon_id INTEGER NOT NULL REFERENCES owned_pokemons(id),
			stat_id          INTEGER NOT NULL,
			value            INTEGER NOT NULL,
			PRIMARY KEY (owned_pokemon_id, stat_id)
		)`,
	}

	for _, s := range stmts {
		if _, err := pool.Exec(ctx, s); err != nil {
			return fmt.Errorf("bot schema error: %w\nstatement: %s", err, s)
		}
	}

	return nil
}
