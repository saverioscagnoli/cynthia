package events

type EventName string

const (
	Hello                               EventName = "HELLO"
	Ready                               EventName = "READY"
	Resumed                             EventName = "RESUMED"
	Reconnect                           EventName = "RECONNECT"
	RateLimited                         EventName = "RATE_LIMITED"
	InvalidSession                      EventName = "INVALID_SESSION"
	ApplicationCommandPermissionsUpdate EventName = "APPLICATION_COMMAND_PERMISSIONS_UPDATE"
	AutoModerationRuleCreate            EventName = "AUTO_MODERATION_RULE_CREATE"
	AutoModerationRuleUpdate            EventName = "AUTO_MODERATION_RULE_UPDATE"
	AutoModerationRuleDelete
	AutoModerationActionExecution
	ChannelCreate
	ChannelUpdate
	ChannelDelete
	ChannelInfo
	ChannelPinsUpdate
	ThreadCreate
	ThreadUpdate
	ThreadDelete
	ThreadListSync
	ThreadMemberUpdate
	ThreadMembersUpdate
	EntitlementCreate
	EntitlementUpdate
	EntitlementDelete
)
