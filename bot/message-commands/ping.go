package messagecommands

import (
	"cynthia/ds"
	"cynthia/ds/payloads"
	"fmt"
)

type PingCommand struct{}

func (p PingCommand) Name() string {
	return "ping"
}

func (p PingCommand) Description() string {
	return "Healthceck. Returns the websocket latency."
}

func (p PingCommand) Handler(client *ds.Client, i payloads.InteractionCreate) error {
	return client.Api.InteractionReplyText(&i, fmt.Sprintf("Pong! Latency: `%dms`", client.Latency().Milliseconds()))
}
