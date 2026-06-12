package commands

import (
	"cynthia/ds"
	"fmt"
	"net/http"
)

type Ping struct{}

func (p Ping) Name() string {
	return "ping"
}

func (p Ping) Description() string {
	return "Pong! 🏓"
}

func (p Ping) Handler(client *ds.Client, i ds.InteractionCreate) {
	msg := fmt.Sprintf("Pong! `%dms`", client.Latency().Milliseconds())

	client.Api.InteractionDefer(&i)
	client.Api.InteractionFollowupText(&i, msg)

	resp, _ := http.Get("http://localhost:9000/pokemon/sprite/25/front")

	defer resp.Body.Close()

	embed := ds.NewEmbed().
		WithImage(&ds.EmbedImage{URL: "attachment://sprite.png"})

	client.Api.SendMessage(*i.ChannelID, &ds.CreateMessageBody{
		Embeds: []*ds.Embed{embed},
		Files: []*ds.MessageFile{{
			Name:        "sprite.png",
			ContentType: "image/png",
			Reader:      resp.Body,
		}},
	})
}
