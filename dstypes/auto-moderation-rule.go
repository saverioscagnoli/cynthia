package dstypes

type AutoModerationRule struct {
	ID              Snowflake              `json:"id"`
	GuildID         Snowflake              `json:"guild_id"`
	Name            string                 `json:"name"`
	CreatorID       Snowflake              `json:"creator_id"`
	EventType       RuleEventType          `json:"event_type"`
	TriggerType     RuleTriggerType        `json:"trigger_type"`
	TriggerMetadata TriggerMetadata        `json:"trigger_metadata"`
	Actions         []AutoModerationAction `json:"actions"`
	Enabled         bool                   `json:"enabled"`
	ExemptRoles     []Snowflake            `json:"exempt_roles"`
	ExemptChannels  []Snowflake            `json:"exempt_channel"`
}
