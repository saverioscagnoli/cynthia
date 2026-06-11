package messagecommands

import (
	"cynthia/ds"
	"cynthia/ds/dsapi"
	"cynthia/ds/dstypes"
	"cynthia/ds/payloads"
	"cynthia/util"
	"os"
)

type MessageCommand interface {
	Name() string
	Description() string
	Handler(client *ds.Client, i payloads.InteractionCreate) error
}

type MessageCommandOptions interface {
	Options() []dstypes.ApplicationCommandOption
}

var Registry map[string]MessageCommand = map[string]MessageCommand{
	"ping": PingCommand{},
	"echo": EchoCommand{},
}

func RegisterAll(client *ds.Client) (int, error) {
	bodies := util.NewVector[dsapi.CreateCommandBody]()

	for _, cmd := range Registry {
		body := dsapi.CreateCommandBody{
			Name:        cmd.Name(),
			Description: cmd.Description(),
		}

		if o, ok := cmd.(MessageCommandOptions); ok {
			body.Options = util.Ptr(o.Options())
		}

		bodies.Push(body)
	}

	err := client.Api.BulkOverwriteGuildCommands(os.Getenv("TEST_GUILD"), *bodies.Arr())

	if err != nil {
		return -1, err
	}

	return len(Registry), nil
}
