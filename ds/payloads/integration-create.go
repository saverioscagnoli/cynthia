package payloads

import "cynthia/ds/dstypes"

type IntegrationCreate struct {
	dstypes.Integration
	GuildID dstypes.Snowflake `json:"guild_id"`
}
