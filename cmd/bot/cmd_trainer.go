package main

import (
	"cynthia/cmd/bot/models"
	"cynthia/ds"
	"cynthia/util"
	"fmt"
	"log/slog"
)

type Trainer struct {
	DB *DB
}

func (t Trainer) Name() string {
	return "trainer"
}

func (t Trainer) Description() string {
	return "Displays information about a trainer. Leave the user black to display yourself."
}

func (t Trainer) Options() *[]ds.ApplicationCommandOption {
	return &[]ds.ApplicationCommandOption{
		{
			Name:        "user",
			Description: "The user you want to display information for. Leave blank to display your info.",
			Type:        ds.ApplicationCommandOptionTypeUser,
			Required:    util.Ptr(false),
		},
	}
}

func (t Trainer) Handler(client *ds.Client, i *ds.InteractionCreate) {
	client.Api.InteractionDefer(i)

	data, _ := i.ApplicationCommandData()

	var user *ds.User = nil

	for _, opt := range util.Deref(data.Options) {
		if opt.Name == "user" {
			if opt.Value != nil {
				uid, ok := (*opt.Value).(ds.Snowflake)
				u, err := client.Api.GetUser(uid)

				if !ok || err != nil {
					err := client.Api.InteractionReplyTextEphemeral(i, "There was an error fetching the user. Please retry.")

					if err != nil {
						slog.Error("Failed to reply to interaction", "err", err)
					}

					return
				}

				if u.Bot != nil && *u.Bot {
					err := client.Api.InteractionReplyTextEphemeral(i, "This user is a bot! Please enter a valid user.")

					if err != nil {
						slog.Error("Failed to reply to interaction", "err", err)
					}

					return
				}

				user = u
			}
		}
	}

	if user == nil {
		user = i.Member.User
	}

	trainer, err := models.GetOrCreateTrainer(t.DB.Context, t.DB.Pool, user.ID)

	if err != nil {
		err := client.Api.InteractionReplyTextEphemeral(i, "There was an error fetching the user. Please retry.")

		if err != nil {
			slog.Error("Failed to reply to interaction", "err", err)
		}
	}

	embed := ds.NewEmbed().
		WithTitle(fmt.Sprintf("Info about %s", user.Username)).
		WithColor(0xff00ff).
		WithDescription(fmt.Sprintf("Money: %d", trainer.Money))

	client.Api.InteractionFollowup(i, ds.InteractionCallbackData{
		Embeds: []*ds.Embed{embed},
	})
}
