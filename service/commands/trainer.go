package commands

import (
	"bytes"
	"context"
	"cynthia/ds"
	"cynthia/service/util"
	"cynthia/store"
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

	var userID string

	if data.Options == nil {

		userID = i.Member.User.ID
	} else {
		for _, opt := range *data.Options {
			if opt.Name == "user" {

				if err := json.Unmarshal(*opt.Value, &userID); err != nil {
					return client.Api.InteractionReplyTextEphemeral(i, "There was an error fetching the user. Please retry.")
				}
			}
		}
	}

	user, err := db.GetUser(userID, context.Background())

	if err != nil {
		return client.Api.InteractionReplyTextEphemeral(i, "Failed to fetch user data. Please retry.")
	}

	d := &ds.InteractionCallbackData{}

	if user.SpriteID != nil {
		img, ok := store.TrainerSprites[*user.SpriteID]

		if ok {
			d.Files = []*ds.MessageFile{{Name: "trainer.png", ContentType: "image/png", Reader: bytes.NewBuffer(*img)}}
		}
	}

	embed := ds.NewEmbed().
		WithTitle(fmt.Sprintf("%s's trainer page", user.Username)).
		WithURL("http://localhost:5173/account").
		WithThumbnail(&ds.EmbedImage{URL: "attachment://trainer.png"}).
		WithAuthor(&ds.EmbedAuthor{Name: &user.DiscordUsername, IconURL: util.Ptr(ds.Routes.AvatarURL(userID, *user.AvatarHash))})

	if user.Banner != nil {
		d.Files = append(d.Files, &ds.MessageFile{Name: "banner.jpg", ContentType: "image/jpg", Reader: bytes.NewBuffer(*user.Banner)})

		embed = embed.WithImage(&ds.EmbedImage{URL: "attachment://banner.jpg"})
	}

	d.Embeds = []*ds.Embed{embed}

	return client.Api.InteractionFollowup(i, d)
}
