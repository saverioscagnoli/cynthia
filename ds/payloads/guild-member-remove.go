package payloads

import "cynthia/dstypes"

type GuildMemberRemove struct {
	GuildID dstypes.Snowflake `json:"guild_id"`
	User    dstypes.User      `json:"user"`
}
