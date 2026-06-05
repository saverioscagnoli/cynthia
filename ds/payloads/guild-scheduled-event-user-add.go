package payloads

import "cynthia/dstypes"

type GuildScheduledEventUserAdd struct {
	GuildScheduledEventID dstypes.Snowflake `json:"guild_scheduled_event_id"`
	UserID                dstypes.Snowflake `json:"user_id"`
	GuildID               dstypes.Snowflake `json:"guild_id"`
}
