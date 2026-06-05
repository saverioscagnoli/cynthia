package payloads

import "cynthia/dstypes"

type AutoModerationActionExecution struct {
	GuildID              dstypes.Snowflake            `json:"guild_id"`
	Action               dstypes.AutoModerationAction `json:"action"`
	RuleID               dstypes.Snowflake            `json:"rule_id"`
	RuleTriggerType      dstypes.RuleTriggerType      `json:"rule_trigger_type"`
	UserID               dstypes.Snowflake            `json:"user_id"`
	ChannelID            *dstypes.Snowflake           `json:"channel_id"`
	MessageID            *dstypes.Snowflake           `json:"message_id"`
	AlertSystemMessageID *dstypes.Snowflake           `json:"alert_system_message_id"`
	Content              string                       `json:"content"`
	MatchedKeyword       *string                      `json:"matched_keyword"`
	MatchedContent       *string                      `json:"matched_content"`
}
