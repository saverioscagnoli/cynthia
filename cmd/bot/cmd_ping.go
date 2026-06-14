package main

import (
	"cynthia/ds"
	"fmt"
)

type Ping struct{}

func (p Ping) Name() string {
	return "ping"
}

func (p Ping) Description() string {
	return "Pong! 🏓"
}

func (p Ping) Handler(client *ds.Client, i *ds.InteractionCreate) {
	msg := fmt.Sprintf("Pong! `%dms`", client.Latency().Milliseconds())
	client.Api.InteractionReplyText(i, msg)
}
