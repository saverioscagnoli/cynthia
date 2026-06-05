package payloads

import "cynthia/dstypes"

type IntegrationCreate struct {
	dstypes.Integration
	GuildID dstypes.Snowflake `json:"guild_id"`
}
