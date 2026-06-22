package main

import (
	"cynthia/ds"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	dbPath := flag.String("db-path", "assets/pokemon.db", "Database path from where to extract information")
	flag.Parse()

	app, err := Init(*dbPath)

	if err != nil {
		slog.Error("Failed to initialize app", "err", err)
		return
	}

	slog.Info("App setup completed")

	ds.On(app.ds, ds.EventReady, func(c *ds.Client, a *ds.Ready) {
		slog.Info("Ready event received", "version", a.Version)
	})

	ds.On(app.ds, ds.EventMessageCreate, func(c *ds.Client, msg *ds.MessageCreate) {
		if msg.Content == "!ping" {
			_, err := c.Api.SendMessageText(msg.ChannelID, fmt.Sprintf("Pong! `%dms`", c.Latency().Milliseconds()))

			if err != nil {
				slog.Error("Error in message create", "err", err)
				return
			}
		}
	})

	ds.On(app.ds, ds.EventInteractionCreate, func(c *ds.Client, i *ds.InteractionCreate) {
		if i.Type == ds.InteractionTypeApplicationCommand {
			data, _ := i.ApplicationCommandData()

			if cmd, ok := c.Commands[data.Name]; ok {
				err := cmd.Handler(c, i)

				if err != nil {
					slog.Error("Error during command execution", "err", err)
				}
			} else {
				slog.Error("Slash command not found", "cmd", data.Name)
			}
		}
	})

	go app.ds.Login()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	slog.Warn("Shutting down")

	app.pkapiStop()

	if err := app.rt.Close(); err != nil {
		slog.Error("Error closing router", "err", err)
	}
}
