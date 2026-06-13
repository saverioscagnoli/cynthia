package commands

import (
	"cynthia/ds"
	"fmt"
	"log/slog"
	"time"
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

	actionRow := ds.NewActionRow()

	button := ds.NewButton().
		WithLabel("Fah!").
		WithStyle(ds.ButtonStyleSecondary).
		WithCustomID("dfdfd")

	actionRow.Add(button)

	client.Api.InteractionReplyText(i, msg)
	sent, err := client.Api.SendMessage(*i.ChannelID, &ds.CreateMessageBody{Content: "a", Components: []ds.MessageComponent{actionRow}})

	if err != nil {
		slog.Error("button", "err", err)
	}

	client.CollectComponents(
		sent.ID,
		time.Duration(100*time.Second),
		2,
		func(c *ds.Client, mi *ds.InteractionCreate) {
			data, _ := mi.MessageComponentData()
			c.Api.InteractionReplyTextEphemeral(mi, "You replied! clicked on "+data.CustomID)
		}, func() {
			button.Disable()
			client.Api.EditMessage(*i.ChannelID, sent.ID, &ds.CreateMessageBody{
				Components: []ds.MessageComponent{actionRow},
			})

		})
}
