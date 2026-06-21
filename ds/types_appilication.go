package ds

import "cynthia/service/util"

type Application struct {
	ID                                Snowflake                           `json:"id"`
	Name                              string                              `json:"name"`
	Icon                              *string                             `json:"icon"`
	Description                       string                              `json:"description"`
	RpcOrigins                        *[]string                           `json:"rpc_origins"`
	BotRequireCodeGrant               bool                                `json:"bot_require_code_grant"`
	Bot                               *User                               `json:"bot"`
	TermsOfServiceURL                 *string                             `json:"terms_of_service_url"`
	PrivacyPolicyURL                  *string                             `json:"privacy_policy_url"`
	Owner                             *User                               `json:"owner"`
	VerifyKey                         string                              `json:"verify_key"`
	Team                              *Team                               `json:"team"`
	GuildID                           *Snowflake                          `json:"guild_id"`
	Guild                             *Guild                              `json:"guild"`
	PrimarySkuID                      *Snowflake                          `json:"primary_sku_id"`
	Slug                              *string                             `json:"slug"`
	CoverImage                        *string                             `json:"cover_image"`
	Flags                             *ApplicationFlags                   `json:"flags"`
	ApproximateGuildCount             *int                                `json:"approximate_guild_count"`
	ApproximateUserInstallCount       *int                                `json:"approximate_user_install_count"`
	ApproximateUserAuthorizationcount *int                                `json:"approximate_user_authorization_count"`
	RedirectURIs                      *[]string                           `json:"redirect_uris"`
	InteractionEndpointsURL           *string                             `json:"interaction_endpoints_url"`
	RoleConnectionsVerificationURL    *string                             `json:"role_connections_verification_url"`
	EventWebhooksURL                  *string                             `json:"event_webhooks_url"`
	EventWebhookStatus                *ApplicationEventWebhookStatus      `json:"event_webhook_status"`
	EventWebhookTypes                 *[]WebhookEventType                 `json:"event_webhook_types"`
	Tags                              *[]string                           `json:"tags"`
	InstallParams                     *InstallParams                      `json:"install_params"`
	IntegrationTypesConfig            map[ApplicationIntegrationType]bool `json:"integration_types_config"`
	CustomInstallURL                  *string                             `json:"custom_install_url"`
}

type ApplicationFlags int

const (
	ApplicationFlagApplicationAutoModerationRuleCreateBadge ApplicationFlags = 1 << 6
	ApplicationFlagGatewayPresence                          ApplicationFlags = 1 << 12
	ApplicationFlagGatewayPresenceLimited                   ApplicationFlags = 1 << 13
	ApplicationFlagGatewayGuildMembers                      ApplicationFlags = 1 << 14
	ApplicationFlagGatewayGuildMembersLimited               ApplicationFlags = 1 << 15
	ApplicationFlagVerificationPendingGuildLimit            ApplicationFlags = 1 << 16
	ApplicationFlagEmbedded                                 ApplicationFlags = 1 << 17
	ApplicationFlagGatewayMessageContent                    ApplicationFlags = 1 << 18
	ApplicationFlagGatewayMessageContentLimited             ApplicationFlags = 1 << 19
	ApplicationFlagApplicationCommandBadge                  ApplicationFlags = 1 << 23
)

func (a *ApplicationFlags) UnmarshalJSON(b []byte) error {
	return util.UnmarshalNumeric(b, a)
}

type ApplicationEventWebhookStatus int

const (
	ApplicationEventWebhookStatusDisabled          ApplicationEventWebhookStatus = 1
	ApplicationEventWebhookStatusEnabled           ApplicationEventWebhookStatus = 2
	ApplicationEventWebhookStatusDisabledByDiscord ApplicationEventWebhookStatus = 3
)

func (a *ApplicationEventWebhookStatus) UnmarshalJSON(b []byte) error {
	return util.UnmarshalNumeric(b, a)
}

type ApplicationIntegrationType int

const (
	ApplicationIntegrationTypesGuildInstall ApplicationIntegrationType = 0
	ApplicationIntegrationTypesUserInstall  ApplicationIntegrationType = 1
)

func (a *ApplicationIntegrationType) UnmarshalJSON(b []byte) error {
	return util.UnmarshalNumeric(b, a)
}

type InstallParams struct {
	Scopes      []OAuth2Scope `json:"scopes"`
	Permissions string        `json:"permissions"`
}

type OAuth2Scope string

const (
	OAuth2ScopeActivitiesRead                        OAuth2Scope = "activities.read"
	OAuth2ScopeActivitiesWrite                       OAuth2Scope = "activities.write"
	OAuth2ScopeApplicationsBuildsRead                OAuth2Scope = "applications.builds.read"
	OAuth2ScopeApplicationsBuildsUpload              OAuth2Scope = "applications.builds.upload"
	OAuth2ScopeApplicationsCommands                  OAuth2Scope = "applications.commands"
	OAuth2ScopeApplicationsCommandsUpdate            OAuth2Scope = "applications.commands.update"
	OAuth2ScopeApplicationsCommandsPermissionsUpdate OAuth2Scope = "applications.commands.permissions.update"
	OAuth2ScopeApplicationsEntitlements              OAuth2Scope = "applications.entitlements"
	OAuth2ScopeApplicationsStoreUpdate               OAuth2Scope = "applications.store.update"
	OAuth2ScopeBot                                   OAuth2Scope = "bot"
	OAuth2ScopeConnections                           OAuth2Scope = "connections"
	OAuth2ScopeDmChannelsRead                        OAuth2Scope = "dm_channels.read"
	OAuth2ScopeEmail                                 OAuth2Scope = "email"
	OAuth2ScopeGdmJoin                               OAuth2Scope = "gdm.join"
	OAuth2ScopeGuilds                                OAuth2Scope = "guilds"
	OAuth2ScopeGuildsJoin                            OAuth2Scope = "guilds.join"
	OAuth2ScopeGuildsMembersRead                     OAuth2Scope = "guilds.members.read"
	OAuth2ScopeIdentify                              OAuth2Scope = "identify"
	OAuth2ScopeIdentifyPremium                       OAuth2Scope = "identify.premium"
	OAuth2ScopeMessagesRead                          OAuth2Scope = "messages.read"
	OAuth2ScopeRelationshipsRead                     OAuth2Scope = "relationships.read"
	OAuth2ScopeRoleConnectionsWrite                  OAuth2Scope = "role_connections.write"
	OAuth2ScopeRpc                                   OAuth2Scope = "rpc"
	OAuth2ScopeRpcActivitiesWrite                    OAuth2Scope = "rpc.activities.write"
	OAuth2ScopeRpcNotificationsRead                  OAuth2Scope = "rpc.notifications.read"
	OAuth2ScopeRpcVoiceRead                          OAuth2Scope = "rpc.voice.read"
	OAuth2ScopeRpcVoiceWrite                         OAuth2Scope = "rpc.voice.write"
	OAuth2ScopeVoice                                 OAuth2Scope = "voice"
	OAuth2ScopeWebhookIncoming                       OAuth2Scope = "webhook.incoming"
)

func (s *OAuth2Scope) UnmarshalJSON(data []byte) error {
	return util.UnmarshalString(data, s)
}
