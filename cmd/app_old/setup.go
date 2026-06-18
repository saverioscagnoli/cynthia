package main

import (
	"context"
	"cynthia/ds"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/lmittmann/tint"
	"github.com/rs/cors"
)

func NewApp(
	token string,
	appID ds.Snowflake,
	DB *DB,
	mux *http.ServeMux,
) *App {
	return &App{
		Client: ds.NewClient(token, appID, ds.WithIntents(ds.IntentAll)),
		DB:     DB,
		mux:    mux,
	}
}

func SetupLogging() {
	var level slog.Level

	levelStr := flag.String("level", "info", "log level (debug, info, warn, error)")
	flag.Parse()

	if err := level.UnmarshalText([]byte(*levelStr)); err != nil {
		level = slog.LevelInfo
	}

	slog.SetDefault(slog.New(tint.NewHandler(os.Stderr, &tint.Options{
		Level:      level,
		TimeFormat: time.TimeOnly,
		AddSource:  true,
	})))
}

func SetupDatabase(username string, passwd string, host string, port string) (*pgxpool.Pool, context.Context, error) {
	ctx := context.Background()
	addr := fmt.Sprintf("postgres://%s:%s@%s:%s/pokemon", username, passwd, host, port)
	pool, err := pgxpool.New(ctx, addr)

	if err != nil {
		return nil, nil, err
	}

	err = pool.Ping(ctx)

	if err != nil {
		return nil, nil, err
	}

	return pool, ctx, nil
}

func createBotSchema(ctx context.Context, pool *pgxpool.Pool) error {
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

func SetupBackend(addr string, port string) *http.ServeMux {
	mux := http.NewServeMux()

	go func() {
		c := cors.Default()
		err := http.ListenAndServe(fmt.Sprintf("%s:%s", addr, port), c.Handler(mux))

		if err != nil {
			slog.Error("Backend server", "err", err)
			return
		}
	}()

	return mux
}

func Init() (*App, error) {
	SetupLogging()
	godotenv.Load()

	slog.Info("Secrets loaded.")

	dbUsername := os.Getenv("DB_USERNAME")
	dbPasswd := os.Getenv("DB_PASSWD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	pool, ctx, err := SetupDatabase(dbUsername, dbPasswd, dbHost, dbPort)

	if err != nil {
		return nil, err
	}

	slog.Info("Connected to database.")

	err = createBotSchema(ctx, pool)

	if err != nil {
		return nil, err
	}

	slog.Info("Checked database schema")

	backendAddr := os.Getenv("BACKEND_ADDR")
	backendPort := os.Getenv("BACKEND_PORT")

	mux := SetupBackend(backendAddr, backendPort)

	slog.Info("Backend listening.", "addr", fmt.Sprintf("%s:%s", backendAddr, backendPort))

	token := os.Getenv("TOKEN")
	appID := os.Getenv("APP_ID")

	return NewApp(token, appID, &DB{Pool: pool, Context: ctx}, mux), nil
}
