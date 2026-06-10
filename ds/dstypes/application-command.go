package dstypes

type ApplicationCommand struct {
	ID                       string                        `json:"id"`
	Type                     *ApplicationCommandType       `json:"type"`
	AppID                    Snowflake                     `json:"app_id"`
	GuildID                  *Snowflake                    `json:"guild_id"`
	Name                     string                        `json:"name"`
	NameLocalizations        *map[string]string            `json:"name_localizations"`
	Description              string                        `json:"description"`
	DescriptionLocalizations *map[string]string            `json:"description_localizations"`
	Options                  *[]ApplicationCommandOption   `json:"options"`
	DefaultMemberPermissions *string                       `json:"default_member_permissions"`
	DMPermission             *bool                         `json:"dm_permission"`
	DefaultPermissions       *bool                         `json:"default_permissions"`
	NSFW                     *bool                         `json:"nsfw"`
	IntegrationTypes         *[]ApplicationIntegrationType `json:"integration_types"`
	Contexts                 *[]InteractionContextType     `json:"contexts"`
	Version                  Snowflake                     `json:"version"`
	Handler                  *EntryPointCommandHandlerType `json:"handler"`
}
