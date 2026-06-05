package dstypes

type Integration struct {
	ID                Snowflake                  `json:"id"`
	Name              string                     `json:"name"`
	Type              IntegrationType            `json:"type"`
	Enabled           bool                       `json:"enabled"`
	Syncing           *bool                      `json:"syncing"`
	RoleID            *Snowflake                 `json:"role_id"`
	EnableEmoticons   *bool                      `json:"enable_emoticons"`
	ExpireBehavior    *IntegrationExpireBehavior `json:"expire_behavior"`
	ExpireGracePeriod *int                       `json:"expire_grace_period"`
	User              *User                      `json:"user"`
	Account           IntegrationAccount         `json:"account"`
	SyncedAt          *string                    `json:"synced_at"`
	SubscriberCount   *int                       `json:"subscriber_count"`
	Revoked           *bool                      `json:"revoked"`
	Application       *Application               `json:"application"`
	Scopes            *[]OAuth2Scope             `json:"scopes"`
}
