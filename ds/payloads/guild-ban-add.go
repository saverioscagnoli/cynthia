package payloads

import "cynthia/ds/dstypes"

type GuildBanAdd struct {
	GuildID dstypes.Snowflake `json:"guild_id"`
	User    dstypes.User      `json:"user"`
}
