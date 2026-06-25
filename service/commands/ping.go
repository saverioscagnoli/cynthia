package commands

import (
	"camilla/ds"
	"fmt"
)

type Ping struct{}

func (p Ping) Name() string {
	return "ping"
}

func (p Ping) Description() string {
	return "Returns gateway latency."
}

func (p Ping) Handler(client *ds.Client, i *ds.InteractionCreate) error {
	msg := fmt.Sprintf("Pong! `%dms`", client.Latency().Milliseconds())
	err := client.Api.InteractionReplyText(i, msg)

	return err
}
