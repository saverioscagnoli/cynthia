package payloads

import "cynthia/ds/dstypes"

type GuildBanRemove struct {
	GuildID dstypes.Snowflake `json:"guild_id"`
	User    dstypes.User      `json:"user"`
}
