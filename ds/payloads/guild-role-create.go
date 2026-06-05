package payloads

import "cynthia/dstypes"

type GuildRoleCreate struct {
	GuildID dstypes.Snowflake `json:"guild_id"`
	Role    dstypes.Role      `json:"role"`
}
