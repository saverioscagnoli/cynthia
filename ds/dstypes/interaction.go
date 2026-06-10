package dstypes

type Interaction struct {
	ID                           Snowflake                          `json:"id"`
	ApplicationID                Snowflake                          `json:"application_id"`
	Type                         InteractionType                    `json:"type"`
	Data                         *any                               `json:"data"`
	Guild                        *Guild                             `json:"guild"`
	GuildID                      *Snowflake                         `json:"guild_id"`
	Channel                      *Channel                           `json:"channel"`
	ChannelID                    *Snowflake                         `json:"channel_id"`
	Member                       *GuildMember                       `json:"member"`
	User                         *User                              `json:"user"`
	Token                        string                             `json:"token"`
	Version                      int                                `json:"version"`
	Message                      *Message                           `json:"message"`
	AppPermissions               string                             `json:"app_permissions"`
	Locale                       *string                            `json:"locale"`
	GuildLocale                  *string                            `json:"guild_locale"`
	Entitlements                 []Entitlement                      `json:"entitlements"`
	AuthorizingIntegrationOwners map[ApplicationIntegrationType]any `json:"authorizing_integration_owners"`
	Contexts                     *InteractionContextType            `json:"contexts"`
	AttachmentSizeLimit          int                                `json:"attachment_size_limit"`
}
