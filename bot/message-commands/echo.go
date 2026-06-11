package messagecommands

import (
	"cynthia/ds"
	"cynthia/ds/dstypes"
	"cynthia/ds/payloads"
	"cynthia/util"
)

type EchoCommand struct{}

func (c EchoCommand) Name() string {
	return "echo"
}

func (c EchoCommand) Description() string {
	return "Repeats what you say."
}

func (c EchoCommand) Options() []dstypes.ApplicationCommandOption {
	return []dstypes.ApplicationCommandOption{
		dstypes.ApplicationCommandOption{Name: "sentence", Description: "What do you want to say?", MinLength: util.Ptr(int64(1)), Required: util.Ptr(true), Type: dstypes.ApplicationCommandOptionTypeString},
	}
}

func (c EchoCommand) Handler(client *ds.Client, i payloads.InteractionCreate) error {
	data := (*i.Data).(map[string]interface{})
	options := data["options"].([]interface{})

	for _, opt := range options {
		o := opt.(map[string]interface{})
		if o["name"].(string) == "sentence" {
			sentence := o["value"].(string)
			client.Api.InteractionReplyText(&i, sentence)
		}
	}

	return nil
}
