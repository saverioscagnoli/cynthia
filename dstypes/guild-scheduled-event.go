package dstypes

type GuildScheduledEvent struct {
	ID                 Snowflake                          `json:"id"`
	GuildID            Snowflake                          `json:"guild_id"`
	ChannelID          *Snowflake                         `json:"channel_id"`
	CreatorID          *Snowflake                         `json:"creator_id"`
	Name               string                             `json:"name"`
	Description        *string                            `json:"description"`
	ScheduledStartTime string                             `json:"scheduled_start_time"`
	ScheduledEndTime   string                             `json:"scheduled_end_time"`
	PrivacyLevel       GuildScheduledEventPrivacyLevel    `json:"privacy_level"`
	Status             GuildScheduledEventStatus          `json:"status"`
	EntityType         GuildScheduledEventEntityType      `json:"entity_type"`
	EntityID           *Snowflake                         `json:"entity_id"`
	EntityMetadata     *GuildScheduledEventEntityMetadata `json:"entity_metadata"`
	Creator            *User                              `json:"creator"`
	UserCount          *int                               `json:"user_count"`
	Image              *string                            `json:"image"`
	RecurrenceRule     *GuildScheduledEventRecurrenceRule `json:"recurrence_rule"`
}
