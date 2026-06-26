package commands

import (
	"bytes"
	"camilla/ds"
	"camilla/service/util"
	"camilla/store"
	"fmt"
	"log/slog"
	"math/rand/v2"
	"strings"
)

var colorHex = map[string]int{
	"red":    0xFF0000,
	"blue":   0x0000FF,
	"yellow": 0xFFFF00,
	"green":  0x00FF00,
	"black":  0x000000,
	"brown":  0xA52A2A,
	"purple": 0x800080,
	"gray":   0x808080,
	"white":  0xFFFFFF,
	"pink":   0xFFC0CB,
}

type Encounter struct{}

func (Encounter) Name() string {
	return "encounter"
}

func (Encounter) Description() string {
	return "Triggers a Pokemon encounter for the author based on place and weather conditions."
}

func (Encounter) Handler(client *ds.Client, i *ds.InteractionCreate) error {
	err := client.Api.InteractionDefer(i)

	if err != nil {
		return err
	}

	id := store.PokemonIDs[rand.IntN(len(store.PokemonIDs))]

	slog.Debug("Triggering pokemon encounter", "id", id)

	name := store.Pokemons[id].Name
	parts := strings.Split(name, "-")
	name = strings.ToUpper(parts[0][:1]) + parts[0][1:]

	color := store.Pokemons[id].Species.Color
	sprites := store.PokemonSpriteMap[id]

	e := ds.NewEmbed().
		WithTitle(fmt.Sprintf("A wild %s has appeared!", name)).
		WithImage(&ds.EmbedImage{URL: "attachment://sprite.png"}).
		WithColor(colorHex[string(color)])

	var files []*ds.MessageFile

	if sprites != nil && sprites.Front != nil {
		sp, err := util.ScaleSprite(*sprites.Front, 2.0)

		if err != nil {
			sp = *sprites.Front
		}

		e = e.WithImage(&ds.EmbedImage{URL: "attachment://sprite.png"})

		files = []*ds.MessageFile{{
			Name:        "sprite.png",
			ContentType: "image/png",
			Reader:      bytes.NewReader(sp),
		}}
	}

	return client.Api.InteractionFollowup(i, &ds.InteractionCallbackData{
		Embeds: []*ds.Embed{e},
		Files:  files,
	})
}
