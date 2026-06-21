package ds

import "cynthia/service/util"

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
