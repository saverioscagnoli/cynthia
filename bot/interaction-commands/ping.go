package interactioncommands

import (
	"cynthia/ds"
	"cynthia/ds/dstypes"
	"cynthia/ds/payloads"
)

func Ping(client *ds.Client, int payloads.InteractionCreate) {
	client.Api.InteractionReply(int.ID, int.Token, dstypes.InteractionResponse{
		Type: dstypes.InteractionCallbackTypeChannelMessageWithSource,
		Data: &dstypes.InteractionCallbackData{
			Content: "Pong! Latency: `" + client.Latency().String() + "ms`",
		},
	})
}
