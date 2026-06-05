package dstypes

type AuditLogChange struct {
	NewValue *any   `json:"new_value"`
	OldValue *any   `json:"old_value"`
	Key      string `json:"key"`
}
