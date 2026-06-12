package commands

import (
	"cynthia/ds"
	"cynthia/util"
	"log/slog"
	"os"
)

type Command interface {
	Name() string
	Description() string
	Handler(client *ds.Client, i ds.InteractionCreate)
}

type CommandOptions interface {
	Options() []ds.ApplicationCommandOption
}

var Registry map[string]Command = map[string]Command{
	"ping": Ping{},
}

func Register(client *ds.Client) {
	bodies := make([]ds.CreateCommandBody, 0, len(Registry))

	for _, cmd := range Registry {
		body := ds.CreateCommandBody{
			Name:        cmd.Name(),
			Description: cmd.Description(),
		}

		if o, ok := cmd.(CommandOptions); ok {
			body.Options = util.Ptr(o.Options())
		}

		bodies = append(bodies, body)
	}

	err := client.Api.BulkOverwriteGuildCommands(os.Getenv("TEST_GUILD"), bodies)

	if err != nil {
		slog.Error("Failed to register guild commands.", "err", err)
	} else {
		slog.Info("Successfully registered guild commands", "count", len(Registry))
	}
}
