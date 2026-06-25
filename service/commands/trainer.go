package commands

import (
	"bytes"
	"camilla/ds"
	"camilla/service/util"
	"camilla/store"
	"context"
	"encoding/json"
	"fmt"
)

type Trainer struct{}

func (t Trainer) Name() string {
	return "trainer"
}

func (t Trainer) Description() string {
	return "Returns a brief description about a trainer."
}

func (t Trainer) Options() []ds.ApplicationCommandOption {
	return []ds.ApplicationCommandOption{
		{
			Name:        "user",
			Description: "The user you want information about. Leave empty to display yourself.",
			Type:        ds.ApplicationCommandOptionTypeUser,
			Required:    util.Ptr(false),
		},
	}
}

func (t Trainer) Handler(client *ds.Client, i *ds.InteractionCreate) error {
	data, err := i.ApplicationCommandData()

	if err != nil {
		return client.Api.InteractionReplyTextEphemeral(i, "Internal error. Please retty.")
	}

	client.Api.InteractionDefer(i)

	var user *ds.User = nil

	if data.Options == nil {
		user = i.Member.User
	} else {
		var userID ds.Snowflake

		for _, opt := range *data.Options {
			if opt.Name == "user" {

				if err := json.Unmarshal(*opt.Value, &userID); err != nil {
					return client.Api.InteractionReplyTextEphemeral(i, "There was an error fetching the user. Please retry.")
				}
			}
		}

		u := (*data.Resolved.Users)[userID]
		user = &u
	}

	u, err := db.GetOrInsertUser(user, context.Background())

	if err != nil {
		return client.Api.InteractionReplyTextEphemeral(i, "Failed to fetch user data. Please retry.")
	}

	d := &ds.InteractionCallbackData{}

	if u.SpriteID != nil {
		img, ok := store.TrainerSprites[*u.SpriteID]

		if ok {
			d.Files = []*ds.MessageFile{{Name: "trainer.png", ContentType: "image/png", Reader: bytes.NewBuffer(*img)}}
		}
	}

	embed := ds.NewEmbed().
		WithTitle(fmt.Sprintf("%s's trainer page", user.Username)).
		WithURL("http://localhost:5173/account").
		WithThumbnail(&ds.EmbedImage{URL: "attachment://trainer.png"}).
		WithAuthor(&ds.EmbedAuthor{Name: &u.DiscordUsername, IconURL: util.Ptr(ds.Routes.AvatarURL(u.ID, *u.AvatarHash))})

	if user.Banner != nil {
		d.Files = append(d.Files, &ds.MessageFile{Name: "banner.jpg", ContentType: "image/jpg", Reader: bytes.NewBuffer(*u.Banner)})

		embed = embed.WithImage(&ds.EmbedImage{URL: "attachment://banner.jpg"})
	}

	d.Embeds = []*ds.Embed{embed}

	return client.Api.InteractionFollowup(i, d)
}
