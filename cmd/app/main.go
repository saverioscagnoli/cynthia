package main

import (
	"cynthia/ds"
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

	app.ds.Login()
}
