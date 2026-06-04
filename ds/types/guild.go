package ds

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
