package payloads

import "cynthia/ds/dstypes"

type GuildRoleDelete struct {
	GuildID dstypes.Snowflake `json:"guild_id"`
	RoleID  dstypes.Snowflake `json:"role_id"`
}
