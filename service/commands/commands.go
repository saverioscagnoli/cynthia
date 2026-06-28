package commands

import (
	"camilla/ds"
	"camilla/service/database"
)

type AppDiscordCommandRegistry interface {
	Register(c *ds.Client, testGuild *ds.Snowflake) error
}

type Registry struct {
	db database.AppDatabase
}

func New(db database.AppDatabase) AppDiscordCommandRegistry {
	return &Registry{db: db}
}

func (r *Registry) Register(c *ds.Client, g *ds.Snowflake) error {
	c.AddCommand(Ping{})
	c.AddCommand(Trainer{db: r.db})
	c.AddCommand(Encounter{})

	if g == nil {
		return c.RegisteGlobalCommands()
	} else {
		return c.RegisterGuildCommands(*g)
	}
}
