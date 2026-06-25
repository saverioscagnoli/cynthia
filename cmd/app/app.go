package main

import (
	"camilla/ds"
	"camilla/service/api"
	"camilla/service/database"
	"context"
	"os"
)

type App struct {
	ds        *ds.Client
	db        database.AppDatabase
	rt        api.Router
	pkapiStop context.CancelFunc
}

func NewApp(db database.AppDatabase, rt api.Router, pkapiStop context.CancelFunc) *App {
	token := os.Getenv("TOKEN")
	appID := os.Getenv("APP_ID")
	ds := ds.NewClient(token, appID, ds.WithIntents(ds.IntentAll))

	return &App{
		ds:        ds,
		db:        db,
		rt:        rt,
		pkapiStop: pkapiStop,
	}
}
