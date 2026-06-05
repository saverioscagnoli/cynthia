package payloads

import "cynthia/dstypes"

type GuildMemberAdd struct {
	dstypes.GuildMember
	GuildId dstypes.Snowflake `json:"guild_id"`
}
