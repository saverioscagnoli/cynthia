package ds

import "camilla/service/util"

type BaseThemeType int

const (
	BaseThemeTypeUnset    BaseThemeType = 0
	BaseThemeTypeDark     BaseThemeType = 1
	BaseThemeTypeLight    BaseThemeType = 2
	BaseThemeTypeDarker   BaseThemeType = 3
	BaseThemeTypeMidnight BaseThemeType = 4
)

func (t *BaseThemeType) UnmarshalJSON(data []byte) error {
	return unmarshalNumeric(data, t)
}

type MembershipState int

const (
	MembershipStateInvited  = 1
	MembershipStateAccepted = 2
)

func (t *MembershipState) UnmarshalJSON(data []byte) error {
	return unmarshalNumeric(data, t)
}

type SharedClientTheme struct {
	Colors        []string       `json:"colors"`
	GradientAngle int            `json:"grandient_angle"`
	BaseMix       int            `json:"base_mix"`
	BaseTheme     *BaseThemeType `json:"base_theme"`
}

type Team struct {
	Icon        *string      `json:"icon"`
	ID          Snowflake    `json:"id"`
	Members     []TeamMember `json:"members"`
	Name        string       `json:"name"`
	OwnerUserID Snowflake    `json:"owner_user_id"`
}

type TeamMember struct {
	MembershipState MembershipState `json:"membership_state"`
	TeamID          Snowflake       `json:"team_id"`
	User            User            `json:"user"`
	Role            string          `json:"role"`
}

type WebhookEventType string

const (
	WebhookEventTypeApplicationAuthorized   WebhookEventType = "APPLICATION_AUTHORIZED"
	WebhookEventTypeApplicationDeauthorized WebhookEventType = "APPLICATION_DEAUTHORIZED"
	WebhookEventTypeEntitlementCreate       WebhookEventType = "ENTITLEMENT_CREATE"
	WebhookEventTypeEntitlementUpdate       WebhookEventType = "ENTITLEMENT_UPDATE"
	WebhookEventTypeEntitlementDelete       WebhookEventType = "ENTITLEMENT_DELETE"
	WebhookEventTypeQuestUserEnrollment     WebhookEventType = "QUEST_USER_ENROLLMENT"
	WebhookEventTypeLobbyMessageCreate      WebhookEventType = "LOBBY_MESSAGE_CREATE"
	WebhookEventTypeLobbyMessageUpdate      WebhookEventType = "LOBBY_MESSAGE_UPDATE"
	WebhookEventTypeLobbyMessageDelete      WebhookEventType = "LOBBY_MESSAGE_DELETE"
	WebhookEventTypeGameDirectMessageCreate WebhookEventType = "GAME_DIRECT_MESSAGE_CREATE"
	WebhookEventTypeGameDirectMessageUpdate WebhookEventType = "GAME_DIRECT_MESSAGE_UPDATE"
	WebhookEventTypeGameDirectMessageDelete WebhookEventType = "GAME_DIRECT_MESSAGE_DELETE"
)

func (e *WebhookEventType) UnmarshalJSON(b []byte) error {
	return util.UnmarshalString(b, e)
}

type StagePrivacyLevel int

const (
	StagePrivacyLevelPublic    StagePrivacyLevel = 1
	StagePrivacyLevelGuildOnly StagePrivacyLevel = 2
)

func (s *StagePrivacyLevel) UnmarshalJSON(b []byte) error {
	return unmarshalNumeric(b, s)
}

type StageInstance struct {
	ID                    Snowflake         `json:"id"`
	GuildID               Snowflake         `json:"guild_id"`
	ChannelID             Snowflake         `json:"channel_id"`
	Topic                 string            `json:"topic"`
	PrivacyLevel          StagePrivacyLevel `json:"privacy_level"`
	DiscoverableDisabled  bool              `json:"discoverable_disabled"`
	GuildScheduledEventID *Snowflake        `json:"guild_scheduled_event_id"`
}

type ActivityType int

const (
	ActivityTypePlaying   ActivityType = 0
	ActivityTypeStreaming ActivityType = 1
	ActivityTypeListening ActivityType = 2
	ActivityTypeWatching  ActivityType = 3
	ActivityTypeCustom    ActivityType = 4
	ActivityTypeCompeting ActivityType = 5
)

func (a *ActivityType) UnmarshalJSON(b []byte) error {
	return unmarshalNumeric(b, a)
}

type ActivityTimestamps struct {
	Start *int `json:"start"`
	End   *int `json:"end"`
}

type StatusDisplayType int

const (
	StatusDisplayTypeName    StatusDisplayType = 0
	StatusDisplayTypeState   StatusDisplayType = 1
	StatusDisplayTypeDetails StatusDisplayType = 2
)

func (s *StatusDisplayType) UnmarshalJSON(b []byte) error {
	return unmarshalNumeric(b, s)
}

type ActivityParty struct {
	ID   *string `json:"id"`
	Size *[]int  `json:"size"`
}

type ActivityAssets struct {
	LargeImage       *string `json:"large_image"`
	LargeText        *string `json:"large_text"`
	LargeURL         *string `json:"large_url"`
	SmallImage       *string `json:"small_image"`
	SmallText        *string `json:"small_text"`
	SmallURL         *string `json:"small_url"`
	InviteCoverImage *string `json:"invite_cover_image"`
}

type ActivitySecrets struct {
	Join     *string `json:"join"`
	Spectate *string `json:"spectate"`
	Match    *string `json:"match"`
}

type ActivityFlags int

const (
	ActivityFlagInstance                 = 1 << 0
	ActivityFlagJoin                     = 1 << 1
	ActivityFlagSpectate                 = 1 << 2
	ActivityFlagJoinRequest              = 1 << 3
	ActivityFlagSync                     = 1 << 4
	ActivityFlagPlay                     = 1 << 5
	ActivityFlagPartyPrivacyFriends      = 1 << 6
	ActivityFlagPartyPrivacyVoiceChannel = 1 << 7
	ActivityFlagEmbedded                 = 1 << 8
)

func (a *ActivityFlags) UnmarshalJSON(b []byte) error {
	return unmarshalNumeric(b, a)
}

type Activity struct {
	Name              string              `json:"name"`
	Type              ActivityType        `json:"type"`
	URL               *string             `json:"url"`
	CreatedAt         int                 `json:"created_at"`
	Timestamps        *ActivityTimestamps `json:"timestamps"`
	ApplicationID     *Snowflake          `json:"application_id"`
	StatusDisplayType StatusDisplayType   `json:"status_display_type"`
	Details           *string             `json:"details"`
	DetailsURL        *string             `json:"details_url"`
	State             *string             `json:"state"`
	StateURL          *string             `json:"state_url"`
	Emoji             *Emoji              `json:"emoji"`
	Party             *ActivityParty      `json:"party"`
	Assets            *ActivityAssets     `json:"assets"`
	Secrets           *ActivitySecrets    `json:"secrets"`
	Instance          *bool               `json:"instance"`
	Flags             *ActivityFlags      `json:"flags"`
	Buttons           *[]Button           `json:"buttons"`
}

type SubscriptionStatus int

const (
	SubscriptionStatusActive   SubscriptionStatus = 0
	SubscriptionStatusEnding   SubscriptionStatus = 1
	SubscriptionStatusInactive SubscriptionStatus = 2
)

func (s *SubscriptionStatus) UnmarshalJSON(b []byte) error {
	return unmarshalNumeric(b, s)
}

type Subscription struct {
	ID                 Snowflake          `json:"id"`
	UserID             Snowflake          `json:"user_id"`
	SkuIDs             []Snowflake        `json:"sku_ids"`
	EntitlementIDs     []Snowflake        `json:"entitlement_ids"`
	RenewalSkuIDs      *[]Snowflake       `json:"renewal_sku_ids"`
	CurrentPeriodStart string             `json:"current_period_start"`
	CurrentPeriodEnd   string             `json:"current_period_end"`
	Status             SubscriptionStatus `json:"status"`
	CanceledAt         *string            `json:"canceled_at"`
	Country            *string            `json:"country"`
}
