package handlers

import (
	"cynthia/ds"
	"cynthia/ds/payloads"
	"fmt"
)

func OnMessageCreate(client *ds.Client, msg payloads.MessageCreate) {
	if msg.Content == "!ping" {
		client.Api.SendMessageText(msg.ChannelID, fmt.Sprintf("Pong! Latency: `%dms`", client.Latency().Milliseconds()))
	}
}
