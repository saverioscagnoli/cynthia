package handlers

import (
	"cynthia/ds"
	"cynthia/ds/dstypes"
	"cynthia/ds/payloads"
	"fmt"
)

func OnMessageCreate(client *ds.Client, msg payloads.MessageCreate) {
	if msg.Content == "!ping" {
		client.SendMessage(msg.ChannelID, fmt.Sprintf("Pong! Latency: `%dms`", client.Latency().Milliseconds()))
		client.SendEmbed(msg.ChannelID, dstypes.NewEmbed().Title("Hello kl;adjfasl;j").Description("Hello world").Color(0x00ffff).Build())
	}
}
