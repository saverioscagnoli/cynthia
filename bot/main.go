package main

import (
	"cynthia/bot/handlers"
	"cynthia/ds"
	"cynthia/ds/dsapi"
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
	client.OnInteractionCreate(handlers.InteractionCreate)

	slog.Info("Registered 3 event handlders.")
}

func main() {
	slog.SetDefault(slog.New(tint.NewHandler(os.Stderr, &tint.Options{
		AddSource:  true,
		Level:      slog.LevelInfo,
		TimeFormat: time.Kitchen,
	})))

	godotenv.Load()

	client := ds.NewClient(os.Getenv("TOKEN"), os.Getenv("APP_ID"))

	registerEventHandlers(client)

	err := client.CreateGuildCommand(dsapi.CreateGuildCommandBody{
		Name:        "ping",
		Description: "bruh",
	}, os.Getenv("TEST_GUILD"))

	if err != nil {
		slog.Error("Failed to create guild command", "error", err)
	}

	client.Start(dstypes.IntentsGuilds | dstypes.IntentsGuildMessages | dstypes.IntentsMessageContent)

}
