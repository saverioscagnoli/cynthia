package main

import (
	"cynthia/bot/commands"
	"cynthia/ds"
	"encoding/json"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/lmittmann/tint"
)

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

	client := ds.NewClient(os.Getenv("TOKEN"), os.Getenv("APP_ID"), ds.WithIntents(ds.IntentGuilds|ds.IntentGuildMessages|ds.IntentMessageContent))

	commands.Register(client)

	ds.On(client, ds.EventReady, func(c *ds.Client, r ds.Ready) {
		slog.Info("Ready event received.", "username", r.User.Username, "version", r.Version)
	})

	ds.On(client, ds.EventMessageCreate, func(c *ds.Client, msg ds.MessageCreate) {
		if msg.Content == "!ping" {
			text := fmt.Sprintf("Pong! `%dms`", c.Latency().Milliseconds())
			c.Api.SendMessageText(msg.ChannelID, text)
		}
	})

	ds.On(client, ds.EventInteractionCreate, func(c *ds.Client, i ds.InteractionCreate) {
		if i.Type == ds.InteractionTypeApplicationCommand || i.Type == ds.InteractionTypeApplicationCommandAutocomplete {
			if i.Data == nil {
				slog.Warn("Received command interaction with nil data")
				return
			}

			var data ds.ApplicationCommandData

			err := json.Unmarshal(*i.Data, &data)

			if err != nil {
				slog.Error("Failed to unmarshal command interaction data", "err", err)
				return
			}

			if cmd, ok := commands.Registry[data.Name]; ok {
				cmd.Handler(c, i)
			} else {
				slog.Error("Couldnt find handler for interaction command", "name", data.Name)
			}
		}
	})

	client.Login()
}
