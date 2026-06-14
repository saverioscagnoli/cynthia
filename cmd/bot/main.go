package main

import (
	"context"
	"cynthia/ds"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	*pgxpool.Pool
	context.Context
}

type App struct {
	*ds.Client
	DB *DB
}

func main() {
	app, err := Init()

	if err != nil {
		slog.Error("Initialization", "err", err)
	}

	defer app.DB.Close()

	slog.Info("App initialized.")

	RegisterCommands(app.Client, app.DB)

	ds.On(app.Client, ds.EventReady, func(c *ds.Client, r *ds.Ready) {
		slog.Info("Ready event received.", "username", r.User.Username, "version", r.Version)
	})

	ds.On(app.Client, ds.EventMessageCreate, func(c *ds.Client, msg *ds.MessageCreate) {
		if msg.Content == "!ping" {
			text := fmt.Sprintf("Pong! `%dms`", c.Latency().Milliseconds())
			c.Api.SendMessageText(msg.ChannelID, text)
		}
	})

	ds.On(app.Client, ds.EventInteractionCreate, func(c *ds.Client, i *ds.InteractionCreate) {
		if i.Type == ds.InteractionTypeApplicationCommand || i.Type == ds.InteractionTypeApplicationCommandAutocomplete {
			if i.Data == nil {
				slog.Warn("Received command interaction with nil data")
				return
			}

			data, err := i.ApplicationCommandData()

			if err != nil {
				slog.Error("Failed to unmarshal command interaction data", "err", err)
				return
			}

			if cmd, ok := CommandsRegistry[data.Name]; ok {
				cmd.Handler(c, i)
			} else {
				slog.Error("Couldnt find handler for interaction command", "name", data.Name)
			}
		}
	})

	err = app.Client.Login()

	if err != nil {
		slog.Error("Gateway", "err", err)
	}
}
