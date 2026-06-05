package payloads

import "cynthia/dstypes"

type GuildMemberUpdate struct {
	GuildID                    dstypes.Snowflake             `json:"guild_id"`
	Roles                      []dstypes.Snowflake           `json:"roles"`
	User                       dstypes.User                  `json:"user"`
	Nick                       *string                       `json:"nick"`
	Avatar                     *string                       `json:"avatar"`
	Banner                     *string                       `json:"banner"`
	JoinedAt                   *string                       `json:"joined_at"`
	PremiumSince               *string                       `json:"premium_since"`
	Deaf                       *bool                         `json:"deaf"`
	Mute                       *bool                         `json:"mute"`
	Pending                    *bool                         `json:"pending"`
	CommunicationDisabledUntil *string                       `json:"communication_disabled_until"`
	AvatarDecorationData       *dstypes.AvatarDecorationData `json:"avatar_decoration_data"`
	Collectibles               *dstypes.Collectibles         `json:"collectibles"`
}
