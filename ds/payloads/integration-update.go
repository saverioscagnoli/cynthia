package payloads

import "cynthia/dstypes"

type IntegrationUpdate struct {
	dstypes.Integration
	GuildID dstypes.Snowflake `json:"guild_id"`
}
