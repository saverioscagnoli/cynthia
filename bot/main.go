package main

import (
	"cynthia/bot/handlers"
	"cynthia/ds"
	"cynthia/ds/dstypes"
	"log/slog"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/lmittmann/tint"
)

func registerEventHandlers(client *ds.Client) {
	client.OnReady(handlers.OnReady)
	client.OnMessageCreate(handlers.OnMessageCreate)

	slog.Info("Registered 2 event handlders.")
}

func main() {
	slog.SetDefault(slog.New(tint.NewHandler(os.Stderr, &tint.Options{
		AddSource:  true,
		Level:      slog.LevelInfo,
		TimeFormat: time.Kitchen,
	})))

	godotenv.Load()

	client := ds.NewClient()

	registerEventHandlers(client)

	client.Start(os.Getenv("TOKEN"), dstypes.IntentsGuilds|dstypes.IntentsGuildMessages|dstypes.IntentsMessageContent)
}
