package ds

import "cynthia/service/util"

type BaseThemeType int

const (
	BaseThemeTypeUnset    BaseThemeType = 0
	BaseThemeTypeDark     BaseThemeType = 1
	BaseThemeTypeLight    BaseThemeType = 2
	BaseThemeTypeDarker   BaseThemeType = 3
	BaseThemeTypeMidnight BaseThemeType = 4
)

func (t *BaseThemeType) UnmarshalJSON(data []byte) error {
	return util.UnmarshalNumeric(data, t)
}

type MembershipState int

const (
	MembershipStateInvited  = 1
	MembershipStateAccepted = 2
)

func (t *MembershipState) UnmarshalJSON(data []byte) error {
	return util.UnmarshalNumeric(data, t)
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
