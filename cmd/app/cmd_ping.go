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
	return "Returns the latency of the websocket"
}

func (p Ping) Handler(c *ds.Client, i *ds.InteractionCreate) error {
	msg := fmt.Sprintf("Pong! `%dms`", c.Latency().Milliseconds())
	err := c.Api.InteractionReplyText(i, msg)

	return err
}
