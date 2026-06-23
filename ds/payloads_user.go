package ds

type EntitlementCreate = Entitlement

type EntitlementUpdate = Entitlement

type EntitlementDelete = Entitlement

type PresenceUpdate struct {
	User         User           `json:"user"`
	GuildID      Snowflake      `json:"guild_id"`
	Status       PresenceStatus `json:"status"`
	Activities   []Activity     `json:"activities"`
	ClientStatus ClientStatus   `json:"client_status"`
}

type UserUpdate = User
