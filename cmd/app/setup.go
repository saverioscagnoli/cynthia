package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/lmittmann/tint"
)

func createSchema(ctx context.Context, pool *pgxpool.Pool) error {
	stmts := []string{
		`CREATE TABLE IF NOT EXISTS users (
    		discord_id  TEXT      PRIMARY KEY NOT NULL,
    		username    TEXT      NOT NULL,
    		avatar      TEXT,
    		trainer_id  INTEGER   DEFAULT 0,
    		created_at  TIMESTAMP NOT NULL DEFAULT NOW()
      	)`,
		`CREATE TABLE IF NOT EXISTS trainers (
			id         TEXT      PRIMARY KEY NOT NULL,
			money      INTEGER   NOT NULL DEFAULT 0,
			created_at TIMESTAMP NOT NULL DEFAULT NOW(),
			trainer_sprite_id	   INTEGER
		)`,
		`CREATE TABLE IF NOT EXISTS owned_pokemons (
			id           SERIAL    PRIMARY KEY,
			trainer_id   TEXT      NOT NULL REFERENCES trainers(id),
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
		`CREATE TABLE IF NOT EXISTS bag (
			trainer_id TEXT    NOT NULL REFERENCES trainers(id),
			item_id    INTEGER NOT NULL,
			quantity   INTEGER NOT NULL DEFAULT 1 CHECK(quantity > 0),
			PRIMARY KEY (trainer_id, item_id)
		)`,
	}

	for _, s := range stmts {
		if _, err := pool.Exec(ctx, s); err != nil {
			return fmt.Errorf("bot schema error: %w\nstatement: %s", err, s)
		}
	}

	return nil
}

func SetupLogging() {
	level := slog.LevelInfo

	if os.Getenv("APP_ENV") == "dev" {
		level = slog.LevelDebug
	}

	slog.SetDefault(slog.New(tint.NewHandler(os.Stderr, &tint.Options{
		Level:      level,
		TimeFormat: time.TimeOnly,
	})))
}

func SetupDatabase() (*pgxpool.Pool, error) {
	username := os.Getenv("DB_USERNAME")
	passwd := os.Getenv("DB_PASSWD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	ctx := context.Background()
	addr := fmt.Sprintf("postgres://%s:%s@%s:%s/pokemon", username, passwd, host, port)
	pool, err := pgxpool.New(ctx, addr)

	if err != nil {
		return nil, err
	}

	err = pool.Ping(ctx)

	if err != nil {
		return nil, err
	}

	return pool, nil
}

func Init() (*App, error) {
	err := godotenv.Load()

	SetupLogging()

	if err != nil {
		return nil, err
	}

	slog.Debug("Secrets loaded")

	pool, err := SetupDatabase()

	if err != nil {
		return nil, err
	}

	slog.Info("Connected to database.")

	return NewApp(pool), nil
}
