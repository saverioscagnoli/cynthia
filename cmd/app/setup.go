package main

import (
	"camilla/ds"
	"camilla/pkapi"
	"camilla/service/api"
	"camilla/service/commands"
	"camilla/service/database"
	"camilla/store"
	"context"
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

func SetupLogging() {
	level := slog.LevelInfo

	if os.Getenv("APP_ENV") == "dev" {
		level = slog.LevelDebug
	}

	slog.SetDefault(slog.New(tint.NewHandler(os.Stderr, &tint.Options{
		Level:      level,
		TimeFormat: time.TimeOnly,
		AddSource:  false,
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

func SetupDatabase() (database.AppDatabase, error) {
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

	db, err := database.New(pool)

	if err != nil {
		return nil, err
	}

	err = db.Ping(ctx)

	if err != nil {
		return nil, err
	}

	slog.Debug("Database ping successful")

	return db, nil
}

func SetupBackend(addr string, port string, db database.AppDatabase, gen *ds.SnowflakeGenerator) (api.Router, error) {
	router, err := api.New(api.Config{
		Logger:             slog.Default(),
		Database:           db,
		SnowflakeGenerator: gen,
	})

	if err != nil {
		return nil, err
	}

	h := router.Handler()

	h, err = registerWebUI(h)

	if err != nil {
		return nil, err
	}

	go func() {

		var c *cors.Cors

		if os.Getenv("APP_ENV") == "dev" {
			c = cors.New(cors.Options{
				AllowedOrigins:   []string{"http://localhost:5173"},
				AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
				AllowedHeaders:   []string{"*"},
				AllowCredentials: true,
			})
		} else {
			c = cors.Default()
		}

		if err := http.ListenAndServe(addr+":"+port, c.Handler(h)); err != nil {
			slog.Error("Backend stopped", "err", err)
		}
	}()

	slog.Info("Backend started", "addr", addr, "port", port)

	return router, nil
}

func SetupDiscordCommands(app *App, testGuild *ds.Snowflake) error {
	registry := commands.New(app.db)

	return registry.Register(app.ds, testGuild)
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

	slog.Debug("Initialized pkhelper.")
	slog.Debug("Starting store api...")

	ctx, cancelPkapi := context.WithCancel(context.Background())
	gen := ds.NewSnowflakeGenerator(1, 0)

	go func() {
		if err := pkapi.Start(ctx, pkapiPort, gen); err != nil {
			slog.Error("Store API stopped", "err", err)
		}
	}()

	db, err := SetupDatabase()

	if err != nil {
		cancelPkapi()
		return nil, err
	}

	slog.Info("Connected to database.")

	addr := os.Getenv("BACKEND_ADDR")
	port := os.Getenv("BACKEND_PORT")

	rt, err := SetupBackend(addr, port, db, gen)

	if err != nil {
		cancelPkapi()
		return nil, err
	}

	app := NewApp(db, rt, cancelPkapi)

	var testGuild *ds.Snowflake = nil

	g, found := os.LookupEnv("TEST_GUILD")

	if found {
		testGuild = &g
	}

	err = SetupDiscordCommands(app, testGuild)

	if err != nil {
		return app, err
	}

	if found {
		slog.Info("Registered guild commands", "count", len(app.ds.Commands), "guild", *testGuild)
	} else {
		slog.Info("Registered commands", "count", len(app.ds.Commands))
	}

	return app, nil
}
