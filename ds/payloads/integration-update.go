package payloads

import "cynthia/ds/dstypes"

type IntegrationUpdate struct {
	dstypes.Integration
	GuildID dstypes.Snowflake `json:"guild_id"`
}
