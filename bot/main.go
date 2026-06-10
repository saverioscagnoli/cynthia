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

	slog.Info("Registered event handlders.", "count", 3)
}

func registerInteractionCommandHandlersGuild(client *ds.Client) {

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
	slog.Info("Registered interaction command handlers", "count", len(handlers.InteractionCommandHandlers))

	client.Api.ClearGuildCommands(os.Getenv("TEST_GUILD"))

	slog.Warn("Cleared all previous guild commands")

	err := client.Api.BulkOverwriteGuildCommands(dstypes.Snowflake(os.Getenv("TEST_GUILD")), []dsapi.CreateGuildCommandBody{
		{
			Name:        "ping",
			Description: "bruh",
		},
		{
			Name:        "twoplustwo",
			Description: "Four!",
		},
	})

	if err != nil {
		slog.Error("Failed to create guild command", "error", err)
	}

	client.Start(dstypes.IntentsGuilds | dstypes.IntentsGuildMessages | dstypes.IntentsMessageContent)

}
