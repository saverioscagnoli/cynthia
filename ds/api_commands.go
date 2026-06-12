package ds

import (
	"fmt"
	"net/http"
)

func (r *routes) GetGlobalCommands(appID Snowflake) (string, string) {
	return http.MethodGet, fmt.Sprintf("/applications/%s/commands", appID)
}

func (r *routes) CreateGlobalCommand(appID Snowflake) (string, string) {
	return http.MethodPost, fmt.Sprintf("/applications/%s/commands", appID)
}

func (r *routes) DeleteGlobalCommand(appID Snowflake, commandID Snowflake) (string, string) {
	return http.MethodDelete, fmt.Sprintf("/applications/%s/commands/%s", appID, commandID)
}

func (r *routes) BulkOverwriteGlobalCommands(appID Snowflake) (string, string) {
	return http.MethodPut, fmt.Sprintf("/applications/%s/commands", appID)
}

func (r *routes) GetGuildCommands(appID Snowflake, guildID Snowflake) (string, string) {
	return http.MethodGet, fmt.Sprintf("/applications/%s/guilds/%s/commands", appID, guildID)
}

func (r *routes) CreateGuildCommand(appID Snowflake, guildID Snowflake) (string, string) {
	return http.MethodPost, fmt.Sprintf("/applications/%s/guilds/%s/commands", appID, guildID)
}

func (r *routes) DeleteGuildCommand(appID Snowflake, guildID Snowflake, commandID Snowflake) (string, string) {
	return http.MethodDelete, fmt.Sprintf("/applications/%s/guilds/%s/commands/%s", appID, guildID, commandID)
}

func (r *routes) BulkOverwriteGuildCommands(appID Snowflake, guildID Snowflake) (string, string) {
	return http.MethodPut, fmt.Sprintf("/applications/%s/guilds/%s/commands", appID, guildID)
}

type CreateCommandBody struct {
	Name                     string                      `json:"name"`
	NameLocalizations        map[string]string           `json:"name_localizations"`
	Description              string                      `json:"description"`
	DescriptionLocalizations map[string]string           `json:"description_localizations"`
	Options                  *[]ApplicationCommandOption `json:"options"`
	DefaultMemberPermissions *string                     `json:"default_member_permissions"`
	DefaultPermissions       *bool                       `json:"default_permissions"`
	Type                     *ApplicationCommandType     `json:"type"`
	NSFW                     *bool                       `json:"nsfw"`
}

func (c *ApiClient) BulkOverwriteGuildCommands(guildID Snowflake, body []CreateCommandBody) error {
	method, endpoint := Routes.BulkOverwriteGuildCommands(c.appID, guildID)
	_, err := c.request(method, endpoint, body)

	return err
}
