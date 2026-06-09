package handlers

import (
	"cynthia/ds"
	"cynthia/ds/payloads"
	"log/slog"
)

func OnReady(client *ds.Client, p payloads.Ready) {
	slog.Info("Received Ready event", "username", p.User.Username, "discriminator", p.User.Discriminator)
}
