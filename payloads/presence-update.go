package payloads

import "cynthia/dstypes"

type PresenceUpdate struct {
	User         dstypes.User           `json:"user"`
	GuildID      dstypes.Snowflake      `json:"guild_id"`
	Status       dstypes.PresenceStatus `json:"status"`
	Activities   []dstypes.Activity     `json:"activities"`
	ClientStatus dstypes.ClientStatus   `json:"client_status"`
}
