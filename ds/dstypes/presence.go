package dstypes

type Presence struct {
	User         User           `json:"user"`
	GuildID      Snowflake      `json:"guild_id"`
	Status       PresenceStatus `json:"presence_status"`
	Activities   []Activity     `json:"activities"`
	ClientStatus ClientStatus   `json:"client_status"`
}
