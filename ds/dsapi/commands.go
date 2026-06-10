package dsapi

import (
	"cynthia/ds/dstypes"
	"fmt"
)

func EndpointCreateGuildCommand(ApplicationID dstypes.Snowflake, GuildID dstypes.Snowflake) string {
	return fmt.Sprintf("%s/applications/%s/guilds/%s/commands", ApiURL, ApplicationID, GuildID)
}

func EndpointBulkOverwriteGuildCommands(ApplicationID dstypes.Snowflake, GuildID dstypes.Snowflake) string {
	return fmt.Sprintf("%s/applications/%s/guilds/%s/commands", ApiURL, ApplicationID, GuildID)
}

type CreateGuildCommandBody struct {
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

func (c *Client) CreateGuildCommand(guildID dstypes.Snowflake, body CreateGuildCommandBody) error {
	_, err := c.Request("POST", EndpointCreateGuildCommand(c.appID, guildID), body)
	return err
}

func (c *Client) BulkOverwriteGuildCommands(guildID dstypes.Snowflake, body []CreateGuildCommandBody) error {
	_, err := c.Request("PUT", EndpointBulkOverwriteGuildCommands(c.appID, guildID), body)
	return err
}

func (c *Client) ClearGuildCommands(guildID dstypes.Snowflake) error {
	_, err := c.Request("DELETE", EndpointBulkOverwriteGuildCommands(c.appID, guildID), []CreateGuildCommandBody{})
	return err
}
