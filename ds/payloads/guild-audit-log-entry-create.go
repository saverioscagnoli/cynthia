package payloads

import "cynthia/ds/dstypes"

type GuildAuditLogEntryCreate struct {
	dstypes.AuditLogEntry
	GuildID dstypes.Snowflake `json:"guild_id"`
}
