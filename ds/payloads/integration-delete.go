package payloads

import "cynthia/dstypes"

type IntegrationDelete struct {
	ID            dstypes.Snowflake  `json:"id"`
	GuildID       dstypes.Snowflake  `json:"guild_id"`
	ApplicationID *dstypes.Snowflake `json:"application_id"`
}
