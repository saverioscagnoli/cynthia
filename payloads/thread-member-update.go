package payloads

import "cynthia/dstypes"

type ThreadMemberUpdate struct {
	dstypes.ThreadMember
	GuildID dstypes.Snowflake `json:"guild_id"`
}
