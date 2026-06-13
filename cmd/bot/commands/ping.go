package commands

import (
	"cynthia/ds"
	"cynthia/util"
	"fmt"
	"log/slog"
	"strconv"
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
		WithLabel("0").
		WithStyle(ds.ButtonStyleSecondary).
		WithCustomID("testbutton")

	actionRow.Add(button)

	client.Api.InteractionReplyText(i, msg)
	sent, err := client.Api.SendMessage(*i.ChannelID, &ds.CreateMessageBody{Components: []ds.MessageComponent{actionRow}})

	if err != nil {
		slog.Error("button", "err", err)
		return
	}

	count := 0
	client.CollectComponents(
		sent.ID,
		ds.CollectorOptions{
			Timeout: 1 * time.Minute,
			Max:     10,
			Handler: func(c *ds.Client, mi *ds.InteractionCreate) {
				c.Api.InteractionDeferUpdate(mi)
				count++
				button.Label = util.Ptr(strconv.Itoa(count))

				c.Api.EditMessage(sent.ChannelID, sent.ID, &ds.CreateMessageBody{
					Components: []ds.MessageComponent{actionRow},
				})
			},
			OnEnd: func() {
				button.Disable()
				client.Api.EditMessage(sent.ChannelID, sent.ID, &ds.CreateMessageBody{
					Components: []ds.MessageComponent{actionRow},
				})
			},
		})

}
