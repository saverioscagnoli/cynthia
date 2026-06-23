package main

import (
	"context"
	"cynthia/ds"
	"cynthia/pkapi"
	"cynthia/service/api"
	"cynthia/service/commands"
	"cynthia/service/database"
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

func SetupBackend(addr string, port string, db database.AppDatabase) (api.Router, error) {
	router, err := api.New(api.Config{
		Logger:   slog.Default(),
		Database: db,
	})

	if err != nil {
		return nil, err
	}

	go func() {
		c := cors.Default()
		if err := http.ListenAndServe(addr+":"+port, c.Handler(router.Handler())); err != nil {
			slog.Error("Backend stopped", "err", err)
		}
	}()

	slog.Info("Backend started", "addr", addr, "port", port)

	return router, nil
}

func SetupDiscordCommands(app *App, testGuild ds.Snowflake) (bool, error) {
	commands.Init(app.db)

	app.ds.AddCommand(commands.Ping{})
	app.ds.AddCommand(commands.Trainer{})

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

	slog.Debug("Initialized pkhelper.")
	slog.Debug("Starting store api...")

	ctx, cancelPkapi := context.WithCancel(context.Background())

	go func() {
		if err := pkapi.Start(ctx, pkapiPort); err != nil {
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

	rt, err := SetupBackend(addr, port, db)

	if err != nil {
		cancelPkapi()
		return nil, err
	}

	app := NewApp(db, rt, cancelPkapi)

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
