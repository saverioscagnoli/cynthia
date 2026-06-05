package payloads

import "cynthia/ds/dstypes"

type ThreadMemberUpdate struct {
	dstypes.ThreadMember
	GuildID dstypes.Snowflake `json:"guild_id"`
}
