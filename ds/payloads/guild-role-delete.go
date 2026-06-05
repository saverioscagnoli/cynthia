package payloads

import "cynthia/dstypes"

type GuildRoleDelete struct {
	GuildID dstypes.Snowflake `json:"guild_id"`
	RoleID  dstypes.Snowflake `json:"role_id"`
}
