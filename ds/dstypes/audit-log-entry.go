package dstypes

type AuditLogEntry struct {
	TargetID   *string                 `json:"target_id"`
	Changes    *[]AuditLogChange       `json:"changes"`
	UserID     *Snowflake              `json:"user_id"`
	ID         Snowflake               `json:"id"`
	ActionType AuditLogEvent           `json:"action_type"`
	Options    *OptionalAuditEntryInfo `json:"options"`
	Reason     *string                 `json:"reason"`
}
