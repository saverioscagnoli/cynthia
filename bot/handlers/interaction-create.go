package handlers

import (
	"cynthia/ds"
	"cynthia/ds/payloads"
	"log/slog"
)

func InteractionCreate(client *ds.Client, int payloads.InteractionCreate) {
	slog.Info("Received interaction", "int", int)

}
