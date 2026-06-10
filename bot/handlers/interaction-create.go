package handlers

import (
	"cynthia/ds"
	"cynthia/ds/dstypes"
	"cynthia/ds/payloads"
	"log/slog"
)

func InteractionCreate(client *ds.Client, int payloads.InteractionCreate) {
	if int.Type == dstypes.InteractionTypeApplicationCommand {

		data, ok := (*int.Data).(map[string]any)

		if !ok {
			slog.Warn("Couldnt parse interaction data", "data", int.Data)
			return
		}

		name, ok := data["name"].(string)

		if !ok {
			slog.Warn("Couldnt parse interaction data", "data", int.Data)
			return
		}

		if handler, ok := InteractionCommandHandlers[name]; ok {
			handler(client, int)
		} else {
			slog.Warn("Couldnt find handler for interaction command", "name", name)
		}
	}

}
