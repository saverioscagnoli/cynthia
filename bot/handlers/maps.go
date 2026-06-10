package handlers

import (
	interactioncommands "cynthia/bot/interaction-commands"
	"cynthia/ds"
	"cynthia/ds/payloads"
)

var InteractionCommandHandlers = map[string]func(client *ds.Client, int payloads.InteractionCreate){
	"ping": interactioncommands.Ping,
}
