package main

import (
	"cynthia/ds"
	"cynthia/service/database"
	"os"
)

type App struct {
	ds *ds.Client
	db database.AppDatabase
}

func NewApp(db database.AppDatabase) *App {
	token := os.Getenv("TOKEN")
	appID := os.Getenv("APP_ID")
	ds := ds.NewClient(token, appID, ds.WithIntents(ds.IntentAll))

	return &App{
		ds: ds,
		db: db,
	}
}
