package payloads

import "cynthia/dstypes"

type RequestGuildMemberRateLimitMetadata struct {
	GuildID   dstypes.Snowflake    `json:"guild_id"`
	Query     string               `json:"query"`
	Limit     int                  `json:"limit"`
	Presences *bool                `json:"presences"`
	UserIDs   *[]dstypes.Snowflake `json:"user_ids"`
	Nonce     *string              `json:"nonce"`
}
