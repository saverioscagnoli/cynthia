package main

import (
	"cynthia/ds"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	ds   *ds.Client
	pool *pgxpool.Pool
}

func NewApp(pool *pgxpool.Pool) *App {
	token := os.Getenv("TOKEN")
	appID := os.Getenv("APP_ID")
	ds := ds.NewClient(token, appID, ds.WithIntents(ds.IntentAll))

	return &App{
		ds:   ds,
		pool: pool,
	}
}
