package main

import (
	"cynthia/dstypes"
	"cynthia/gateway"
	"cynthia/payloads"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	client := gateway.NewClient()
	token := os.Getenv("TOKEN")

	gateway.ReadyEvent.AddHandler(client, func(c *gateway.Client, p payloads.Ready) {
		slog.Info("Ready received!", "user", p.User.Username, "id", p.User.ID, "v", p.Version)
	})

	gateway.MessageCreate.AddHandler(client, func(c *gateway.Client, msg payloads.MessageCreate) {
		if msg.Content == "!ping" {
			err := c.SendMessage(msg.ChannelID, "pong!")

			if err != nil {
				slog.Error("Error while sending message", "err", err)
			}
		}
	})

	client.Start(token, dstypes.IntentsGuilds|dstypes.IntentsGuildMessages|dstypes.IntentsMessageContent)
}
