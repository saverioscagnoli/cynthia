package payloads

import "cynthia/dstypes"

type GuildRoleUpdate struct {
	GuildID dstypes.Snowflake `json:"guild_id"`
	Role    dstypes.Role      `json:"role"`
}
