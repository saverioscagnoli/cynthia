package payloads

import "cynthia/ds/dstypes"

type GuildMemberAdd struct {
	dstypes.GuildMember
	GuildId dstypes.Snowflake `json:"guild_id"`
}
