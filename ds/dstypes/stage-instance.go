package dstypes

type StageInstance struct {
	ID                    Snowflake         `json:"id"`
	GuildID               Snowflake         `json:"guild_id"`
	ChannelID             Snowflake         `json:"channel_id"`
	Topic                 string            `json:"topic"`
	PrivacyLevel          StagePrivacyLevel `json:"privacy_level"`
	DiscoverableDisabled  bool              `json:"discoverable_disabled"`
	GuildScheduledEventID *Snowflake        `json:"guild_scheduled_event_id"`
}
