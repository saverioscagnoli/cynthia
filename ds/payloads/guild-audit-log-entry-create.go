package payloads

import "cynthia/dstypes"

type GuildAuditLogEntryCreate struct {
	dstypes.AuditLogEntry
	GuildID dstypes.Snowflake `json:"guild_id"`
}
