package handlers

import (
	messagecommands "cynthia/bot/message-commands"
	"cynthia/ds"
	"cynthia/ds/dstypes"
	"cynthia/ds/payloads"
	"log/slog"
)

func InteractionCreate(client *ds.Client, i payloads.InteractionCreate) {
	if i.Type == dstypes.InteractionTypeApplicationCommand {

		data, ok := (*i.Data).(map[string]any)

		if !ok {
			slog.Error("Couldnt parse interaction data", "data", i.Data)
			return
		}

		name, ok := data["name"].(string)

		if !ok {
			slog.Error("Couldnt parse interaction data", "data", i.Data)
			return
		}

		if command, ok := messagecommands.Registry[name]; ok {
			command.Handler(client, i)
		} else {
			slog.Error("Couldnt find handler for interaction command", "name", name)
		}
	}

}
