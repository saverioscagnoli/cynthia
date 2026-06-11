package main

import (
	"cynthia/bot/handlers"
	messagecommands "cynthia/bot/message-commands"
	"cynthia/ds"
	"cynthia/ds/dstypes"
	"flag"
	"log/slog"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/lmittmann/tint"
)

func registerEventHandlers(client *ds.Client) {
	client.OnReady(handlers.OnReady)
	client.OnMessageCreate(handlers.OnMessageCreate)
	client.OnInteractionCreate(handlers.InteractionCreate)

	slog.Info("Registered event handlders.", "count", 3)
}

func main() {
	levelFlag := flag.String("level", "info", "log level (debug, info, warn, error)")
	flag.Parse()

	var logLevel slog.Level

	if err := logLevel.UnmarshalText([]byte(*levelFlag)); err != nil {
		logLevel = slog.LevelInfo
	}

	slog.SetDefault(slog.New(tint.NewHandler(os.Stderr, &tint.Options{
		Level:      logLevel,
		TimeFormat: time.Kitchen,
	})))

	godotenv.Load()
	slog.Info(".env Loaded.")

	client := ds.NewClient(os.Getenv("TOKEN"), os.Getenv("APP_ID"))

	registerEventHandlers(client)

	count, err := messagecommands.RegisterAll(client)

	if err != nil {
		slog.Error("Failed to register application commands.", "err", err)
	} else {
		slog.Info("Successfuly registered application commands.", "count", count)
	}

	client.Start(dstypes.IntentsGuilds | dstypes.IntentsGuildMessages | dstypes.IntentsMessageContent)
}
