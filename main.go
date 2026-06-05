package main

import (
	"cynthia/ds/dstypes"
	"cynthia/ds/gateway"
	"cynthia/ds/payloads"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/lmittmann/tint"
)

func main() {
	slog.SetDefault(slog.New(tint.NewHandler(os.Stderr, &tint.Options{
		AddSource:  true,
		Level:      slog.LevelInfo,
		TimeFormat: time.Kitchen,
	})))

	godotenv.Load()

	client := gateway.NewClient()
	token := os.Getenv("TOKEN")

	client.OnMessageCreate(func(c *gateway.Client, msg payloads.MessageCreate) {
		if msg.Content == "!ping" {
			err := c.SendMessage(msg.ChannelID, fmt.Sprintf("Pong! Latency: `%dms`", c.Latency().Milliseconds()))

			if err != nil {
				slog.Error("Error while sending message", "err", err)
			}
		}

	})

	client.OnChannelCreate(func(c *gateway.Client, p payloads.ChannelCreate) {
		slog.Warn("Channel created", "channel", p)
	})

	client.Start(token, dstypes.IntentsGuilds|dstypes.IntentsGuildMessages|dstypes.IntentsMessageContent)
}
