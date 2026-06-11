package dsapi

import (
	"cynthia/ds/dstypes"
	"fmt"
)

// Endpoint for getting the application command objects for all commands of all guilds.
//
// METHOD: GET
func EndpointGetGlobalCommands(appID dstypes.Snowflake) string {
	return fmt.Sprintf("%s/applications/%s/commands", ApiURL, appID)
}

// Endpoint for creating global commands, affecting all guilds.
// The system needs some time to sync the commands for all guilds, so it will not be instant.
//
// For testing purposes, it's better to create a 'guild' command, which is essentially the same
// thing, but is instant for a single specified guild.
//
// METHOD: POST
func EndpointCreateGlobalCommand(appID dstypes.Snowflake) string {
	return fmt.Sprintf("%s/applications/%s/commands", ApiURL, appID)
}

func EndpointDeleteGlobalCommand(appID dstypes.Snowflake, commandID dstypes.Snowflake) string {
	return fmt.Sprintf("%s/applications/%s/commands/%s", ApiURL, appID, commandID)
}

func EndpointBulkOverwriteGlobalCommands(appID dstypes.Snowflake) string {
	return fmt.Sprintf("%s/applications/%s/commands", ApiURL, appID)
}

// Endpoint for getting all application commands within a guild.
//
// METHOD: GET
func EndpointGetGuildCommands(appID dstypes.Snowflake, guildID dstypes.Snowflake) string {
	return fmt.Sprintf("%s/applications/%s/guilds/%s/commands", ApiURL, appID, guildID)
}

// Endpoint for creating application command within a single guild. Unlike global commands,
// guild commands are created / overwritten instantly.
//
// METHOD: POST
func EndpointCreateGuildCommand(appID dstypes.Snowflake, guildID dstypes.Snowflake) string {
	return fmt.Sprintf("%s/applications/%s/guilds/%s/commands", ApiURL, appID, guildID)
}

func EndpointDeleteGuildCommand(appID dstypes.Snowflake, guildID dstypes.Snowflake, commandID dstypes.Snowflake) string {
	return fmt.Sprintf("%s/applications/%s/guilds/%s/commands/%s", ApiURL, appID, guildID, commandID)
}

func EndpointBulkOverwriteGuildCommands(appID dstypes.Snowflake, guildID dstypes.Snowflake) string {
	return fmt.Sprintf("%s/applications/%s/guilds/%s/commands", ApiURL, appID, guildID)
}

type CreateCommandBody struct {
	Name                     string                              `json:"name"`
	NameLocalizations        map[string]string                   `json:"name_localizations"`
	Description              string                              `json:"description"`
	DescriptionLocalizations map[string]string                   `json:"description_localizations"`
	Options                  *[]dstypes.ApplicationCommandOption `json:"options"`
	DefaultMemberPermissions *string                             `json:"default_member_permissions"`
	DefaultPermissions       *bool                               `json:"default_permissions"`
	Type                     *dstypes.ApplicationCommandType     `json:"type"`
	NSFW                     *bool                               `json:"nsfw"`
}

func (c *Client) GetGlobalCommands() ([]dstypes.ApplicationCommand, error) {
	return Request[[]dstypes.ApplicationCommand](c, "GET", EndpointGetGlobalCommands(c.appID), nil)
}

func (c *Client) CreateGlobalCommand(body CreateCommandBody) error {
	_, err := Request[any](c, "POST", EndpointCreateGlobalCommand(c.appID), body)
	return err
}

func (c *Client) DeleteGlobalCommand(commandID dstypes.Snowflake) error {
	_, err := Request[any](c, "DELETE", EndpointDeleteGlobalCommand(c.appID, commandID), nil)
	return err
}

func (c *Client) BulkOverwriteGlobalCommands(body []CreateCommandBody) error {
	_, err := Request[any](c, "PUT", EndpointBulkOverwriteGlobalCommands(c.appID), body)
	return err
}

func (c *Client) DeleteGlobalCommands() error {
	commands, err := c.GetGlobalCommands()

	if err != nil {
		return err
	}

	for _, command := range commands {
		err = c.DeleteGlobalCommand(command.ID)

		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Client) GetGuildCommands(guildID dstypes.Snowflake) ([]dstypes.ApplicationCommand, error) {
	return Request[[]dstypes.ApplicationCommand](c, "GET", EndpointGetGuildCommands(c.appID, guildID), nil)

}

func (c *Client) CreateGuildCommand(guildID dstypes.Snowflake, body CreateCommandBody) error {
	_, err := Request[any](c, "POST", EndpointCreateGuildCommand(c.appID, guildID), body)
	return err
}

func (c *Client) DeleteGuildCommand(guildID dstypes.Snowflake, commandID dstypes.Snowflake) error {
	_, err := Request[any](c, "DELETE", EndpointDeleteGuildCommand(c.appID, guildID, commandID), nil)
	return err

}

func (c *Client) BulkOverwriteGuildCommands(guildID dstypes.Snowflake, body []CreateCommandBody) error {
	_, err := Request[any](c, "PUT", EndpointBulkOverwriteGuildCommands(c.appID, guildID), body)
	return err
}

func (c *Client) ClearGuildCommands(guildID dstypes.Snowflake) error {
	_, err := Request[any](c, "PUT", EndpointBulkOverwriteGuildCommands(c.appID, guildID), []CreateCommandBody{})
	return err
}

func (c *Client) DeleteGuildCommands(guildID dstypes.Snowflake) error {
	commands, err := c.GetGuildCommands(guildID)

	if err != nil {
		return err
	}

	for _, command := range commands {
		err = c.DeleteGuildCommand(guildID, command.ID)

		if err != nil {
			return err
		}
	}

	return nil
}
