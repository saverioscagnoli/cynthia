package payloads

import "cynthia/dstypes"

type GuildIntegrationsUpdate struct {
	GuildID dstypes.Snowflake `json:"guild_id"`
}
