package main

import (
	"cynthia/ds"
	"cynthia/ds/payloads"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	client := ds.NewClient()

	ds.ReadyEvent.AddHandler(client, func(c *ds.Client, p payloads.Ready) {
		slog.Info("Ready received!", "user", p.User.Username, "id", p.User.ID, "v", p.Version)
	})

	client.Start(os.Getenv("TOKEN"))
}
