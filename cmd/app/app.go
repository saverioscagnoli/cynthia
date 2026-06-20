package main

import (
	"cynthia/ds"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type db struct {
	pool *pgxpool.Pool
}

type backend struct {
	db                db
	frontendURL       string
	jwtSecret         string
	oauth2ClientID    string
	oauth2Secret      string
	oauth2RedirectURI string
}

type App struct {
	ds      *ds.Client
	backend *backend
}

func NewApp(pool *pgxpool.Pool, b *backend) *App {
	token := os.Getenv("TOKEN")
	appID := os.Getenv("APP_ID")
	ds := ds.NewClient(token, appID, ds.WithIntents(ds.IntentAll))

	return &App{
		ds:      ds,
		backend: b,
	}
}
