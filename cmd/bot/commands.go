package main

import (
	"cynthia/ds"
	"log/slog"
	"os"
)

type Command interface {
	Name() string
	Description() string
	Handler(client *ds.Client, i *ds.InteractionCreate)
}

type CommandOptions interface {
	Options() *[]ds.ApplicationCommandOption
}

var CommandsRegistry map[string]Command = map[string]Command{}

func RegisterCommands(client *ds.Client, DB *DB) {
	CommandsRegistry["trainer"] = &Trainer{DB: DB}
	CommandsRegistry["ping"] = &Ping{}

	bodies := make([]ds.CreateCommandBody, 0, len(CommandsRegistry))

	for _, cmd := range CommandsRegistry {
		body := ds.CreateCommandBody{
			Name:        cmd.Name(),
			Description: cmd.Description(),
		}

		if o, ok := cmd.(CommandOptions); ok {
			body.Options = o.Options()
		}

		bodies = append(bodies, body)
	}

	err := client.Api.BulkOverwriteGuildCommands(os.Getenv("TEST_GUILD"), bodies)

	if err != nil {
		slog.Error("Failed to register guild commands.", "err", err)
	} else {
		slog.Info("Successfully registered guild commands", "count", len(CommandsRegistry))
	}
}
