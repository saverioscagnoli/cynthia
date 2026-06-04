package ds

type Application struct {
	ID                                Snowflake                            `json:"id"`
	Name                              string                               `json:"name"`
	Icon                              *string                              `json:"icon"`
	Description                       string                               `json:"description"`
	RpcOrigins                        *[]string                            `json:"rpc_origins"`
	BotRequireCodeGrant               bool                                 `json:"bot_require_code_grant"`
	Bot                               *User                                `json:"bot"`
	TermsOfServiceURL                 *string                              `json:"terms_of_service_url"`
	PrivacyPolicyURL                  *string                              `json:"privacy_policy_url"`
	Owner                             *User                                `json:"owner"`
	VerifyKey                         string                               `json:"verify_key"`
	Team                              *Team                                `json:"team"`
	GuildID                           *Snowflake                           `json:"guild_id"`
	Guild                             *Guild                               `json:"guild"`
	PrimarySkuID                      *Snowflake                           `json:"primary_sku_id"`
	Slug                              *string                              `json:"slug"`
	CoverImage                        *string                              `json:"cover_image"`
	Flags                             *ApplicationFlags                    `json:"flags"`
	ApproximateGuildCount             *int                                 `json:"approximate_guild_count"`
	ApproximateUserInstallCount       *int                                 `json:"approximate_user_install_count"`
	ApproximateUserAuthorizationcount *int                                 `json:"approximate_user_authorization_count"`
	RedirectURIs                      *[]string                            `json:"redirect_uris"`
	InteractionEndpointsURL           *string                              `json:"interaction_endpoints_url"`
	RoleConnectionsVerificationURL    *string                              `json:"role_connections_verification_url"`
	EventWebhooksURL                  *string                              `json:"event_webhooks_url"`
	EventWebhookStatus                *ApplicationEventWebhookStatus       `json:"event_webhook_status"`
	EventWebhookTypes                 *[]WebhookEventType                  `json:"event_webhook_types"`
	Tags                              *[]string                            `json:"tags"`
	InstallParams                     *InstallParams                       `json:"install_params"`
	IntegrationTypesConfig            map[ApplicationIntegrationTypes]bool `json:"integration_types_config"`
	CustomInstallURL                  *string                              `json:"custom_install_url"`
}
