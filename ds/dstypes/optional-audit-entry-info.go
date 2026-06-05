package dstypes

type OptionalAuditEntryInfo struct {
	ApplicationID                 Snowflake `json:"application_id"`
	AutoModerationRuleName        string    `json:"auto_moderation_rule_name"`
	AutoModerationRuleTriggerType string    `json:"auto_moderation_rule_trigger_type"`
	ChannelID                     Snowflake `json:"channel_id"`
	Count                         string    `json:"count"`
	DeleteMemberDays              string    `json:"delete_member_days"`
	ID                            Snowflake `json:"id"`
	MembersRemoved                string    `json:"members_removed"`
	MessageID                     Snowflake `json:"message_id"`
	RoleName                      string    `json:"role_name"`
	Type                          string    `json:"type"`
	IntegrationType               string    `json:"integration_type"`
	Status                        string    `json:"status"`
}
