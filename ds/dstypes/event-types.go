package dstypes

import "cynthia/util"

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
