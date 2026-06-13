package commands

import (
	"cynthia/ds"
	"cynthia/internal/pokemon"
	"fmt"
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

	pkmn, _ := pokemon.DefaultClient.GetPokemon(25)

	embed := ds.NewEmbed().
		WithTitle(pkmn.Name).
		WithImage(&ds.EmbedImage{URL: "attachment://sprite.png"})

	client.Api.SendMessage(*i.ChannelID, &ds.CreateMessageBody{
		Embeds: []*ds.Embed{embed},
	})
}
