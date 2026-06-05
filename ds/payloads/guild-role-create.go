package payloads

import "cynthia/ds/dstypes"

type GuildRoleCreate struct {
	GuildID dstypes.Snowflake `json:"guild_id"`
	Role    dstypes.Role      `json:"role"`
}
