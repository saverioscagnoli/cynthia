package payloads

import "cynthia/ds/dstypes"

type InviteCreate struct {
	ChannelID         dstypes.Snowflake    `json:"channel_id"`
	Code              string               `json:"code"`
	CreatedAt         string               `json:"created_at"`
	GuildID           *dstypes.Snowflake   `json:"guild_id"`
	Inviter           *dstypes.User        `json:"inviter"`
	MaxAge            int                  `json:"max_age"`
	MaxUses           int                  `json:"max_uses"`
	TargetType        dstypes.InviteTarget `json:"target_type"`
	TargetUser        *dstypes.User        `json:"target_user"`
	TargetApplication *dstypes.Application `json:"target_application"`
	Temporary         bool                 `json:"temporary"`
	Uses              int                  `json:"uses"`
	ExpiresAt         *string              `json:"expires_at"`
	RoleIDs           *[]dstypes.Snowflake `json:"role_ids"`
}
