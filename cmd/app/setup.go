package main

import (
	"context"
	"cynthia/ds"
	"cynthia/pkapi"
	"cynthia/store"
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
		`CREATE TABLE IF NOT EXISTS bag (
			user_id TEXT    NOT NULL REFERENCES users(id),
			item_id    INTEGER NOT NULL,
			quantity   INTEGER NOT NULL DEFAULT 1 CHECK(quantity > 0),
			PRIMARY KEY (user_id, item_id)
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
		AddSource:  true,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.LevelKey {
				if lvl, ok := a.Value.Any().(slog.Level); ok && lvl == slog.LevelDebug {
					return tint.Attr(33, a)
				}
			}

			return a
		},
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

	err = createSchema(ctx, pool)

	if err != nil {
		return nil, err
	}

	slog.Debug("Database schema checked")

	err = pool.Ping(ctx)

	if err != nil {
		return nil, err
	}

	slog.Debug("Database ping successful")

	return pool, nil
}

func SetupBackend(addr string, port string, pool *pgxpool.Pool) (*backend, error) {
	frontendURL := os.Getenv("FRONTEND_URL")
	jwtSecret := os.Getenv("JWT_SECRET")
	oauth2ClientID := os.Getenv("OAUTH2_CLIENT_ID")
	oauth2Secret := os.Getenv("OAUTH2_SECRET")
	oauth2RedirectURI := os.Getenv("OAUTH2_REDIRECT_URI")

	b := &backend{
		db:                db{pool: pool},
		frontendURL:       frontendURL,
		jwtSecret:         jwtSecret,
		oauth2ClientID:    oauth2ClientID,
		oauth2Secret:      oauth2Secret,
		oauth2RedirectURI: oauth2RedirectURI,
	}

	publicMux := http.NewServeMux()
	protectedMux := http.NewServeMux()

	publicMux.HandleFunc("GET /healthcheck", b.Healthcheck)
	publicMux.HandleFunc("GET /auth/login", b.AuthLogin)
	publicMux.HandleFunc("GET /auth/callback", b.AuthCallback)

	publicMux.Handle("/user/", b.authMiddleware(protectedMux))

	protectedMux.HandleFunc("GET /user/me", b.GetCurrentUser)
	protectedMux.HandleFunc("GET /user/banner", b.GetBanner)
	protectedMux.HandleFunc("PUT /user/banner", b.UpdateBanner)

	go func() {
		c := cors.Default()
		addr = fmt.Sprintf("%s:%s", addr, port)
		err := http.ListenAndServe(addr, c.Handler(publicMux))

		if err != nil {
			slog.Error("Backend serve error", "err", err)
			return
		}
	}()

	return b, nil
}

func SetupDiscordCommands(app *App, testGuild ds.Snowflake) (bool, error) {
	app.ds.AddCommand(Ping{})

	if testGuild != "" {
		err := app.ds.RegisterGuildCommands(testGuild)
		return false, err
	}

	err := app.ds.RegisteGlobalCommands()
	return true, err
}

func Init(dbPath string) (*App, error) {
	err := godotenv.Load()

	SetupLogging()

	if err != nil {
		return nil, err
	}

	slog.Debug("Secrets loaded")

	slog.Debug("Extracting data...")
	store.Extract(dbPath)
	slog.Info("Pokemon store filled")

	pkapiPort := os.Getenv("PKAPI_PORT")

	slog.Debug("Starting store api...")
	go pkapi.Start(pkapiPort)

	pool, err := SetupDatabase()

	if err != nil {
		return nil, err
	}

	slog.Info("Connected to database.")

	addr := os.Getenv("BACKEND_ADDR")
	port := os.Getenv("BACKEND_PORT")

	b, err := SetupBackend(addr, port, pool)

	if err != nil {
		return nil, err
	}

	app := NewApp(pool, b)

	testGuild := os.Getenv("TEST_GUILD")
	global, err := SetupDiscordCommands(app, testGuild)

	if err != nil {
		return app, err
	}

	if global {
		slog.Info("Registered commands", "count", len(app.ds.Commands))
	} else {
		slog.Info("Registered guild commands", "count", len(app.ds.Commands), "guild", testGuild)
	}

	return app, nil
}
