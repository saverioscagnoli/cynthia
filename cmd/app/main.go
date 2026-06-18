package main

import (
	"cynthia/ds"
	"fmt"
	"log/slog"
)

func main() {
	app, err := Init()

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

	app.ds.Login()
}
