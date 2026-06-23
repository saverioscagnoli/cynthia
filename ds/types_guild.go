package ds

import "cynthia/util"

type Guild struct {
	ID                        Snowflake                  `json:"id"`
	Name                      string                     `json:"name"`
	Icon                      *string                    `json:"icon"`
	IconHash                  *string                    `json:"icon_hash"`
	Splash                    *string                    `json:"splash"`
	DiscoverySplash           *string                    `json:"discovery_splash"`
	Owner                     *bool                      `json:"owner"`
	OwnerID                   Snowflake                  `json:"owner_id"`
	Permissions               *string                    `json:"permissions"`
	Region                    *string                    `json:"region"`
	AFKChannelID              *Snowflake                 `json:"afk_channel_id"`
	AFKTimeout                int                        `json:"afk_timeout"`
	WidgetEnabled             *bool                      `json:"widget_enabled"`
	WidgetChannelID           *Snowflake                 `json:"widget_channel_id"`
	VerificationLevel         VerificationLevel          `json:"verification_level"`
	ExplicitContentFilter     ExplicitContentFilterLevel `json:"explicit_content_filter"`
	Roles                     []Role                     `json:"roles"`
	Emojis                    []Emoji                    `json:"emojis"`
	Features                  []GuildFeature             `json:"features"`
	MfaLevel                  MfaLevel                   `json:"mfa_level"`
	ApplicationID             *Snowflake                 `json:"application_id"`
	SystemChannelID           *Snowflake                 `json:"system_channel_id"`
	SystemChannelFlags        SystemChannelFlags         `json:"system_channel_flags"`
	RulesChannelID            *Snowflake                 `json:"rules_channel_id"`
	MaxPresences              *int                       `json:"max_presences"`
	MaxMembers                *int                       `json:"max_members"`
	VanityURLCode             *string                    `json:"vanity_url_code"`
	Description               *string                    `json:"description"`
	Banner                    *string                    `json:"banner"`
	PremiumTier               PremiumTier                `json:"premium_tier"`
	PremiumSubscriptionCount  *int                       `json:"premium_subscription_count"`
	PreferredLocale           string                     `json:"preferred_locale"`
	PublicUpdatesChannelID    *Snowflake                 `json:"public_updates_channel_id"`
	MaxVideoChannelUsers      *int                       `json:"max_video_channel_users"`
	MaxStageVideoChannelUsers *int                       `json:"max_stage_video_channel_users"`
	ApproximateMemberCount    *int                       `json:"approximate_member_count"`
	ApproximatePresencecount  *int                       `json:"approximate_presence_count"`
	WelcomeScreen             *WelcomeScreen             `json:"welcome_screen"`
	NsfwLevel                 GuildAgeRestrictionLevel   `json:"nsfw_level"`
	Stickers                  *[]Sticker                 `json:"stickers"`
	PremiumProgressBarEnabled bool                       `json:"premium_progress_bar_enabled"`
	SafetyAlertsChannelID     *Snowflake                 `json:"safety_alerts_channel_id"`
	IncidentsData             *IncidentsData             `json:"incidents_data"`
}

type GuildAgeRestrictionLevel int

const (
	GuildAgeRestrictionLevelDefault       GuildAgeRestrictionLevel = 0
	GuildAgeRestrictionLevelExplicit      GuildAgeRestrictionLevel = 1
	GuildAgeRestrictionLevelSafe          GuildAgeRestrictionLevel = 2
	GuildAgeRestructionLevelAgeRestricted GuildAgeRestrictionLevel = 3
)

func (e *GuildAgeRestrictionLevel) UnmarshalJSON(b []byte) error {
	return util.UnmarshalNumeric(b, e)
}

type GuildFeature string

const (
	GuildFeatureAnimatedBanner                        GuildFeature = "ANIMATED_BANNER"
	GuildFeatureAnimatedIcon                          GuildFeature = "ANIMATED_ICON"
	GuildFeatureApplicationCommandPermissionsV2       GuildFeature = "APPLICATION_COMMAND_PERMISSIONS_V2"
	GuildFeatureAutoModeration                        GuildFeature = "AUTO_MODERATION"
	GuildFeatureBanner                                GuildFeature = "BANNER"
	GuildFeatureCommunity                             GuildFeature = "COMMUNITY"
	GuildFeatureCreatorMonetizableProvisional         GuildFeature = "CREATOR_MONETIZABLE_PROVISIONAL"
	GuildFeatureCreatorStorePage                      GuildFeature = "CREATOR_STORE_PAGE"
	GuildFeatureDeveloperSupportServer                GuildFeature = "DEVELOPER_SUPPORT_SERVER"
	GuildFeatureDiscoverable                          GuildFeature = "DISCOVERABLE"
	GuildFeatureFeaturable                            GuildFeature = "FEATURABLE"
	GuildFeatureInvitesDisabled                       GuildFeature = "INVITES_DISABLED"
	GuildFeatureInviteSplash                          GuildFeature = "INVITE_SPLASH"
	GuildFeatureMemberVerificationGateEnabled         GuildFeature = "MEMBER_VERIFICATION_GATE_ENABLED"
	GuildFeatureMoreSoundboard                        GuildFeature = "MORE_SOUNDBOARD"
	GuildFeatureMoreStickers                          GuildFeature = "MORE_STICKERS"
	GuildFeatureNews                                  GuildFeature = "NEWS"
	GuildFeaturePartnered                             GuildFeature = "PARTNERED"
	GuildFeaturePreviewEnabled                        GuildFeature = "PREVIEW_ENABLED"
	GuildFeatureRaidAlertsDisabled                    GuildFeature = "RAID_ALERTS_DISABLED"
	GuildFeatureRoleIcons                             GuildFeature = "ROLE_ICONS"
	GuildFeatureRoleSubscriptionsAvailableForPurchase GuildFeature = "ROLE_SUBSCRIPTIONS_AVAILABLE_FOR_PURCHASE"
	GuildFeatureRoleSubscriptionsEnabled              GuildFeature = "ROLE_SUBSCRIPTIONS_ENABLED"
	GuildFeatureSoundboard                            GuildFeature = "SOUNDBOARD"
	GuildFeatureTicketedEventsEnabled                 GuildFeature = "TICKETED_EVENTS_ENABLED"
	GuildFeatureVanityURL                             GuildFeature = "VANITY_URL"
	GuildFeatureVerified                              GuildFeature = "VERIFIED"
	GuildFeatureVipRegions                            GuildFeature = "VIP_REGIONS"
	GuildFeatureWelcomeScreenEnabled                  GuildFeature = "WELCOME_SCREEN_ENABLED"
	GuildFeatureGuestsEnabled                         GuildFeature = "GUESTS_ENABLED"
	GuildFeatureGuildTags                             GuildFeature = "GUILD_TAGS"
	GuildFeatureEnhancedRoleColors                    GuildFeature = "ENHANCED_ROLE_COLORS"
)

func (e *GuildFeature) UnmarshalJSON(b []byte) error {
	return util.UnmarshalString(b, e)
}

type GuildMember struct {
	User                       *User                 `json:"user"`
	Nick                       *string               `json:"nick"`
	Avatar                     *string               `json:"avatar"`
	Banner                     *string               `json:"banner"`
	Roles                      []Snowflake           `json:"roles"`
	JoinedAt                   string                `json:"joined_at"`
	PremiumSince               *string               `json:"premium_since"`
	Deaf                       bool                  `json:"deaf"`
	Mute                       bool                  `json:"mute"`
	Flags                      int                   `json:"flags"`
	Pending                    *bool                 `json:"pending"`
	Permissions                *string               `json:"permissions"`
	CommunicationDisabledUntil *string               `json:"communication_disabled_until"`
	AvatarDecorationData       *AvatarDecorationData `json:"avatar_decoration_data"`
	Collectibles               *Collectibles         `json:"collectibles"`
}

type ExplicitContentFilterLevel int

const (
	ExplicitContentFilterLevelDisabled            = 0
	ExplicitContentFilterLevelMembersWithoutRoles = 1
	ExplicitContentFilterLevelAllMembers          = 2
)

func (e *ExplicitContentFilterLevel) UnmarshalJSON(b []byte) error {
	return util.UnmarshalNumeric(b, e)
}

type IncidentsData struct {
	InvitesDisabledUntil string `json:"invites_disabled_until"`
	DmsDisabledUntil     string `json:"dms_disabled_until"`
	DmSpamDetectedAt     string `json:"dm_spam_detected_at"`
	RaidDetectedAt       string `json:"raid_detected_at"`
}

type MfaLevel int

const (
	MfaLevelNone     MfaLevel = 0
	MfaLevelElevated MfaLevel = 1
)

func (m *MfaLevel) UnmarshalJSON(data []byte) error {
	return util.UnmarshalNumeric(data, m)
}

type PremiumTier int

const (
	PremiumTierNone PremiumTier = 0
	PremiumTier1    PremiumTier = 1
	PremiumTier2    PremiumTier = 2
	PremiumTier3    PremiumTier = 3
)

func (t *PremiumTier) UnmarshalJSON(data []byte) error {
	return util.UnmarshalNumeric(data, t)
}

type Role struct {
	ID           string     `json:"id"`
	Name         string     `json:"name"`
	Colors       RoleColors `json:"colors"`
	Hoist        bool       `json:"hoist"`
	Icon         *string    `json:"icon"`
	UnicodeEmoji *string    `json:"unicode_emoji"`
	Position     int        `json:"position"`
	Permissions  string     `json:"permissions"`
	Managed      bool       `json:"managed"`
	Mentionable  bool       `json:"mentionable"`
	Tags         RoleTags   `json:"tags"`
	Flags        int        `json:"flags"`
}

type RoleColors struct {
	PrimaryColor   int  `json:"primary_color"`
	SecondaryColor *int `json:"secondary_color"`
	TertiaryColor  *int `json:"tertiary_color"`
}

type RoleSubscriptionData struct {
	RoleSubscriptionListingID Snowflake `json:"role_subscription_listing_id"`
	TierName                  string    `json:"tier_name"`
	TotalMonthsSubscribed     int       `json:"total_months_subscribed"`
	IsRenewal                 bool      `json:"is_renewal"`
}

type RoleTags struct {
	BotID                 Snowflake `json:"bot_id"`
	IntegrationID         Snowflake `json:"integration_id"`
	PremiumSubscriber     *bool     `json:"premium_subscriber"`
	SubscriptionListingID Snowflake `json:"subscription_listing_id"`
	AvailableForPurchase  *bool     `json:"available_for_purchase"`
	GuildConnections      *bool     `json:"guild_connections"`
}

type VerificationLevel int

const (
	VerificationLevelNone     VerificationLevel = 0
	VerificationLevelLow      VerificationLevel = 1
	VerificationLevelMedium   VerificationLevel = 2
	VerificationLevelHigh     VerificationLevel = 3
	VerificationLevelVeryHigh VerificationLevel = 4
)

func (t *VerificationLevel) UnmarshalJSON(data []byte) error {
	return util.UnmarshalNumeric(data, t)
}

type WelcomeScreen struct {
	Description     *string                `json:"description"`
	WelcomeChannels []WelcomeScreenChannel `json:"welcome_channels"`
}

type RequestGuildMemberRateLimitMetadata struct {
	GuildID   Snowflake    `json:"guild_id"`
	Query     string       `json:"query"`
	Limit     int          `json:"limit"`
	Presences *bool        `json:"presences"`
	UserIDs   *[]Snowflake `json:"user_ids"`
	Nonce     *string      `json:"nonce"`
}

type RuleEventType = int

const (
	RuleEventTypeMessageSend  RuleEventType = 1
	RuleEventTypeMemberUpdate RuleEventType = 2
)

type RuleTriggerType int

const (
	RuleTriggerTypeKeyword       RuleTriggerType = 1
	RuleTriggerTypeSpam          RuleTriggerType = 3
	RuleTriggerTypeKeywordPreset RuleTriggerType = 4
	RuleTriggerTypeMentionSpam   RuleTriggerType = 5
	RuleTriggerTypeMemberProfile RuleTriggerType = 6
)

type KeywordPresetTypes int

const (
	KeywordPresetTypesProfanity     KeywordPresetTypes = 1
	KeywordPresetTypesSexualContent KeywordPresetTypes = 2
	KeywordPresetTypesSlurs         KeywordPresetTypes = 3
)

type TriggerMetadata struct {
	KeywordFilter                []string             `json:"keyword_filter"`
	RegexPatterns                []string             `json:"regex_patterns"`
	Presets                      []KeywordPresetTypes `json:"presets"`
	AllowList                    []string             `json:"allow_list"`
	MentionTotalLimit            int                  `json:"mention_total_limit"`
	MentionRaidProtectionEnabled bool                 `json:"mention_raid_protection_enabled"`
}

type AutoModerationActionType int

const (
	AutoModerationActionTypeBlockMessage           AutoModerationActionType = 1
	AutoModerationActionTypeSendAlertMessage       AutoModerationActionType = 2
	AutoModerationActionTypeTimeout                AutoModerationActionType = 3
	AutoModerationActionTypeBlockMemberInteraction AutoModerationActionType = 4
)

type AutoModerationActionMetadata struct {
	ChannelID       Snowflake `json:"channel_id"`
	DurationSeconds int       `json:"duration_seconds"`
	CustomMessage   *string   `json:"custom_message"`
}

type AutoModerationAction struct {
	Type     AutoModerationActionType      `json:"type"`
	Metadata *AutoModerationActionMetadata `json:"metadata"`
}

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

type GuildScheduledEventPrivacyLevel int

const (
	GuildScheduledEventPrivacyLevelGuildOnly GuildScheduledEventPrivacyLevel = 2
)

func (g *GuildScheduledEventPrivacyLevel) UnmarshalJSON(b []byte) error {
	return util.UnmarshalNumeric(b, g)
}

type GuildScheduledEventStatus int

const (
	GuildScheduledEventStatusScheduled GuildScheduledEventStatus = 1
	GuildScheduledEventStatusActive    GuildScheduledEventStatus = 2
	GuildScheduledEventCompleted       GuildScheduledEventStatus = 3
	GuildScheduledEventCanceled        GuildScheduledEventStatus = 4
)

func (g *GuildScheduledEventStatus) UnmarshalJSON(b []byte) error {
	return util.UnmarshalNumeric(b, g)
}

type GuildScheduledEventEntityType int

const (
	GuildScheduledEventEntityTypeStageInstace GuildScheduledEventEntityType = 1
	GuildScheduledEventEntityTypeVoice        GuildScheduledEventEntityType = 2
	GuildScheduledEventEntityTypeExternal     GuildScheduledEventEntityType = 3
)

func (g *GuildScheduledEventEntityType) UnmarshalJSON(b []byte) error {
	return util.UnmarshalNumeric(b, g)
}

type GuildScheduledEventEntityMetadata struct {
	Location *string `json:"location"`
}

type GuildScheduledEventRecurrenceRuleFrequency int

const (
	GuildScheduledEventRecurrenceRuleFrequencyYearly  GuildScheduledEventRecurrenceRuleFrequency = 0
	GuildScheduledEventRecurrenceRuleFrequencyMonthly GuildScheduledEventRecurrenceRuleFrequency = 1
	GuildScheduledEventRecurrenceRuleFrequencyWeekly  GuildScheduledEventRecurrenceRuleFrequency = 2
	GuildScheduledEventRecurrenceRuleFrequencyDaily   GuildScheduledEventRecurrenceRuleFrequency = 3
)

func (g *GuildScheduledEventRecurrenceRuleFrequency) UnmarshalJSON(b []byte) error {
	return util.UnmarshalNumeric(b, g)
}

type GuildScheduledEventRecurrenceRuleWeekday int

const (
	GuildScheduledEventRecurrenceRuleWeekdayMonday    GuildScheduledEventRecurrenceRuleWeekday = 0
	GuildScheduledEventRecurrenceRuleWeekdayTuesday   GuildScheduledEventRecurrenceRuleWeekday = 1
	GuildScheduledEventRecurrenceRuleWeekdayWednesday GuildScheduledEventRecurrenceRuleWeekday = 2
	GuildScheduledEventRecurrenceRuleWeekdayThursday  GuildScheduledEventRecurrenceRuleWeekday = 3
	GuildScheduledEventRecurrenceRuleWeekdayFriday    GuildScheduledEventRecurrenceRuleWeekday = 4
	GuildScheduledEventRecurrenceRuleWeekdaySaturday  GuildScheduledEventRecurrenceRuleWeekday = 5
	GuildScheduledEventRecurrenceRuleWeekdaySunday    GuildScheduledEventRecurrenceRuleWeekday = 6
)

func (g *GuildScheduledEventRecurrenceRuleWeekday) UnmarshalJSON(b []byte) error {
	return util.UnmarshalNumeric(b, g)
}

type GuildScheduledEventRecurrenceRuleNWeekday struct {
	N   int                                      `json:"n"`
	Day GuildScheduledEventRecurrenceRuleWeekday `json:"day"`
}

type GuildScheduledEventRecurrenceRuleMonth int

const (
	GuildScheduledEventRecurrenceRuleMonthJanuary   GuildScheduledEventRecurrenceRuleMonth = 1
	GuildScheduledEventRecurrenceRuleMonthFebruary  GuildScheduledEventRecurrenceRuleMonth = 2
	GuildScheduledEventRecurrenceRuleMonthMarch     GuildScheduledEventRecurrenceRuleMonth = 3
	GuildScheduledEventRecurrenceRuleMonthApril     GuildScheduledEventRecurrenceRuleMonth = 4
	GuildScheduledEventRecurrenceRuleMonthMay       GuildScheduledEventRecurrenceRuleMonth = 5
	GuildScheduledEventRecurrenceRuleMonthJune      GuildScheduledEventRecurrenceRuleMonth = 6
	GuildScheduledEventRecurrenceRuleMonthJuly      GuildScheduledEventRecurrenceRuleMonth = 7
	GuildScheduledEventRecurrenceRuleMonthAugust    GuildScheduledEventRecurrenceRuleMonth = 8
	GuildScheduledEventRecurrenceRuleMonthSeptember GuildScheduledEventRecurrenceRuleMonth = 9
	GuildScheduledEventRecurrenceRuleMonthOctober   GuildScheduledEventRecurrenceRuleMonth = 10
	GuildScheduledEventRecurrenceRuleMonthNovember  GuildScheduledEventRecurrenceRuleMonth = 11
	GuildScheduledEventRecurrenceRuleMonthDecember  GuildScheduledEventRecurrenceRuleMonth = 12
)

func (g *GuildScheduledEventRecurrenceRuleMonth) UnmarshalJSON(b []byte) error {
	return util.UnmarshalNumeric(b, g)
}

type GuildScheduledEventRecurrenceRule struct {
	Start      string                                       `json:"start"`
	End        string                                       `json:"end"`
	Frequency  GuildScheduledEventRecurrenceRuleFrequency   `json:"frequency"`
	Interval   int                                          `json:"interval"`
	ByWeekday  *[]GuildScheduledEventRecurrenceRuleWeekday  `json:"by_weekday"`
	ByNWeekday *[]GuildScheduledEventRecurrenceRuleNWeekday `json:"by_n_weekday"`
	ByMonth    *[]GuildScheduledEventRecurrenceRuleMonth    `json:"by_month"`
	ByMonthDay *[]int                                       `json:"by_month_day"`
	ByYearDay  *[]int                                       `json:"by_year_day"`
	Count      *int                                         `json:"count"`
}

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

type SoundboardSound struct {
	Name      string     `json:"name"`
	SoundID   Snowflake  `json:"sound_id"`
	Volume    float64    `json:"volume"`
	EmojiID   *Snowflake `json:"emoji_id"`
	EmojiName *string    `json:"emoji_name"`
	GuildID   *Snowflake `json:"guild_id"`
	Available bool       `json:"available"`
	User      *User      `json:"user"`
}

type AuditLogChange struct {
	NewValue *any   `json:"new_value"`
	OldValue *any   `json:"old_value"`
	Key      string `json:"key"`
}

type AuditLogEvent int

const (
	AuditLogEventGuildUpdate AuditLogEvent = 1

	AuditLogEventChannelCreate          AuditLogEvent = 10
	AuditLogEventChannelUpdate          AuditLogEvent = 11
	AuditLogEventChannelDelete          AuditLogEvent = 12
	AuditLogEventChannelOverwriteCreate AuditLogEvent = 13
	AuditLogEventChannelOverwriteUpdate AuditLogEvent = 14
	AuditLogEventChannelOverwriteDelete AuditLogEvent = 15

	AuditLogEventMemberKick       AuditLogEvent = 20
	AuditLogEventMemberPrune      AuditLogEvent = 21
	AuditLogEventMemberBanAdd     AuditLogEvent = 22
	AuditLogEventMemberBanRemove  AuditLogEvent = 23
	AuditLogEventMemberUpdate     AuditLogEvent = 24
	AuditLogEventMemberRoleUpdate AuditLogEvent = 25
	AuditLogEventMemberMove       AuditLogEvent = 26
	AuditLogEventMemberDisconnect AuditLogEvent = 27
	AuditLogEventBotAdd           AuditLogEvent = 28

	AuditLogEventRoleCreate AuditLogEvent = 30
	AuditLogEventRoleUpdate AuditLogEvent = 31
	AuditLogEventRoleDelete AuditLogEvent = 32

	AuditLogEventInviteCreate AuditLogEvent = 40
	AuditLogEventInviteUpdate AuditLogEvent = 41
	AuditLogEventInviteDelete AuditLogEvent = 42

	AuditLogEventWebhookCreate AuditLogEvent = 50
	AuditLogEventWebhookUpdate AuditLogEvent = 51
	AuditLogEventWebhookDelete AuditLogEvent = 52

	AuditLogEventEmojiCreate AuditLogEvent = 60
	AuditLogEventEmojiUpdate AuditLogEvent = 61
	AuditLogEventEmojiDelete AuditLogEvent = 62

	AuditLogEventMessageDelete     AuditLogEvent = 72
	AuditLogEventMessageBulkDelete AuditLogEvent = 73
	AuditLogEventMessagePin        AuditLogEvent = 74
	AuditLogEventMessageUnpin      AuditLogEvent = 75

	AuditLogEventIntegrationCreate AuditLogEvent = 80
	AuditLogEventIntegrationUpdate AuditLogEvent = 81
	AuditLogEventIntegrationDelete AuditLogEvent = 82

	AuditLogEventStageInstanceCreate AuditLogEvent = 83
	AuditLogEventStageInstanceUpdate AuditLogEvent = 84
	AuditLogEventStageInstanceDelete AuditLogEvent = 85

	AuditLogEventStickerCreate AuditLogEvent = 90
	AuditLogEventStickerUpdate AuditLogEvent = 91
	AuditLogEventStickerDelete AuditLogEvent = 92

	AuditLogEventGuildScheduledEventCreate AuditLogEvent = 100
	AuditLogEventGuildScheduledEventUpdate AuditLogEvent = 101
	AuditLogEventGuildScheduledEventDelete AuditLogEvent = 102

	AuditLogEventThreadCreate AuditLogEvent = 110
	AuditLogEventThreadUpdate AuditLogEvent = 111
	AuditLogEventThreadDelete AuditLogEvent = 112

	AuditLogEventApplicationCommandPermissionUpdate AuditLogEvent = 121

	AuditLogEventSoundboardSoundCreate AuditLogEvent = 130
	AuditLogEventSoundboardSoundUpdate AuditLogEvent = 131
	AuditLogEventSoundboardSoundDelete AuditLogEvent = 132

	AuditLogEventAutoModerationRuleCreate                AuditLogEvent = 140
	AuditLogEventAutoModerationRuleUpdate                AuditLogEvent = 141
	AuditLogEventAutoModerationRuleDelete                AuditLogEvent = 142
	AuditLogEventAutoModerationBlockMessage              AuditLogEvent = 143
	AuditLogEventAutoModerationFlagToChannel             AuditLogEvent = 144
	AuditLogEventAutoModerationUserCommunicationDisabled AuditLogEvent = 145
	AuditLogEventAutoModerationQuarantineUser            AuditLogEvent = 146

	AuditLogEventCreatorMonetizationRequestCreated AuditLogEvent = 150
	AuditLogEventCreatorMonetizationTermsAccepted  AuditLogEvent = 151

	AuditLogEventOnboardingPromptCreate AuditLogEvent = 163
	AuditLogEventOnboardingPromptUpdate AuditLogEvent = 164
	AuditLogEventOnboardingPromptDelete AuditLogEvent = 165
	AuditLogEventOnboardingCreate       AuditLogEvent = 166
	AuditLogEventOnboardingUpdate       AuditLogEvent = 167

	AuditLogEventHomeSettingsCreate AuditLogEvent = 190
	AuditLogEventHomeSettingsUpdate AuditLogEvent = 191

	AuditLogEventVoiceChannelStatusUpdate AuditLogEvent = 192
	AuditLogEventVoiceChannelStatusDelete AuditLogEvent = 193
)

func (a *AuditLogEvent) UnmarshalJSON(b []byte) error {
	return util.UnmarshalNumeric(b, a)
}

type OptionalAuditEntryInfo struct {
	ApplicationID                 Snowflake `json:"application_id"`
	AutoModerationRuleName        string    `json:"auto_moderation_rule_name"`
	AutoModerationRuleTriggerType string    `json:"auto_moderation_rule_trigger_type"`
	ChannelID                     Snowflake `json:"channel_id"`
	Count                         string    `json:"count"`
	DeleteMemberDays              string    `json:"delete_member_days"`
	ID                            Snowflake `json:"id"`
	MembersRemoved                string    `json:"members_removed"`
	MessageID                     Snowflake `json:"message_id"`
	RoleName                      string    `json:"role_name"`
	Type                          string    `json:"type"`
	IntegrationType               string    `json:"integration_type"`
	Status                        string    `json:"status"`
}

type AuditLogEntry struct {
	TargetID   *string                 `json:"target_id"`
	Changes    *[]AuditLogChange       `json:"changes"`
	UserID     *Snowflake              `json:"user_id"`
	ID         Snowflake               `json:"id"`
	ActionType AuditLogEvent           `json:"action_type"`
	Options    *OptionalAuditEntryInfo `json:"options"`
	Reason     *string                 `json:"reason"`
}

type IntegrationType string

const (
	IntegrationTypeDiscord           IntegrationType = "discord"
	IntegrationTypeGuildSubscription IntegrationType = "guild_subscription"
	IntegrationTypeTwitch            IntegrationType = "twitch"
	IntegrationTypeYoutube           IntegrationType = "youtupe"
)

func (i *IntegrationType) UnmarshalJSON(b []byte) error {
	return util.UnmarshalString(b, i)
}

type IntegrationExpireBehavior int

const (
	IntegrationExpireBehaviorRemoveRole IntegrationExpireBehavior = 0
	IntegrationExpireBehaviorKick       IntegrationExpireBehavior = 1
)

func (i *IntegrationExpireBehavior) UnmarshalJSON(b []byte) error {
	return util.UnmarshalNumeric(b, i)
}

type IntegrationAccount struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

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

type InviteTarget int

const (
	InviteTargetStream              InviteTarget = 1
	InviteTargetEmbeddedApplication InviteTarget = 2
)

func (i *InviteTarget) UnmarshalJSON(b []byte) error {
	return util.UnmarshalNumeric(b, i)
}
