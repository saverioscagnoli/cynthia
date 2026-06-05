package payloads

import "cynthia/ds/dstypes"

type GuildIntegrationsUpdate struct {
	GuildID dstypes.Snowflake `json:"guild_id"`
}
