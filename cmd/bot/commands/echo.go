package commands

import (
	"cynthia/ds"
	"cynthia/util"
)

type Echo struct{}

func (e Echo) Name() string {
	return "echo"
}

func (e Echo) Description() string {
	return "Repeats what you said."
}

func (e Echo) Options() *[]ds.ApplicationCommandOption {
	return &[]ds.ApplicationCommandOption{
		{
			Name:        "sentence",
			Description: "What do you want to say?",
			Type:        ds.ApplicationCommandOptionTypeString,
			Required:    util.Ptr(true),
			MinLength:   util.Ptr(int64(1)),
		},
	}
}

func (e Echo) Handler(client *ds.Client, i *ds.InteractionCreate) {
	data, err := i.ApplicationCommandData()

	if err != nil {
		return
	}

	var sentence string

	for _, opt := range *data.Options {
		if opt.Name == "sentence" {
			sentence = (*opt.Value).(string)
		}
	}

	client.Api.InteractionReplyText(i, sentence)
}
