package ds

import "encoding/json"

type Event[T any] struct {
	name string
}

var (
	EventReady                               = Event[*Ready]{"READY"}
	EventResumed                             = Event[*Resumed]{"RESUMED"}
	EventReconnect                           = Event[*Reconnect]{"RECONNECT"}
	EventRateLimited                         = Event[*RateLimited]{"RATE_LIMITED"}
	EventInvalidSession                      = Event[*InvalidSession]{"INVALID_SESSION"}
	EventApplicationCommandPermissionsUpdate = Event[*ApplicationCommandPermissionsUpdate]{"APPLICATION_COMMAND_PERMISSIONS_UPDATE"}
	EventAutoModerationRuleCreate            = Event[*AutoModerationRuleCreate]{"AUTO_MODERATION_RULE_CREATE"}
	EventAutoModerationRuleUpdate            = Event[*AutoModerationRuleUpdate]{"AUTO_MODERATION_RULE_UPDATE"}
	EventAutoModerationRuleDelete            = Event[*AutoModerationRuleDelete]{"AUTO_MODERATION_RULE_DELETE"}
	EventAutoModerationActionExecution       = Event[*AutoModerationActionExecution]{"AUTO_MODERATION_ACTION_EXECUTION"}
	EventChannelCreate                       = Event[*ChannelCreate]{"CHANNEL_CREATE"}
	EventChannelUpdate                       = Event[*ChannelUpdate]{"CHANNEL_UPDATE"}
	EventChannelDelete                       = Event[*ChannelDelete]{"CHANNEL_DELETE"}
	EventChannelInfo                         = Event[*ChannelInfo]{"CHANNEL_INFO"}
	EventChannelPinsUpdate                   = Event[*ChannelPinsUpdate]{"CHANNEL_PINS_UPDATE"}
	EventThreadCreate                        = Event[*ThreadCreate]{"THREAD_CREATE"}
	EventThreadUpdate                        = Event[*ThreadUpdate]{"THREAD_UPDATE"}
	EventThreadDelete                        = Event[*ThreadDelete]{"THREAD_DELETE"}
	EventThreadListSync                      = Event[*ThreadListSync]{"THREAD_LIST_SYNC"}
	EventThreadMemberUpdate                  = Event[*ThreadMemberUpdate]{"THREAD_MEMBER_UPDATE"}
	EventThreadMembersUpdate                 = Event[*ThreadMembersUpdate]{"THREAD_MEMBERS_UPDATE"}
	EventEntitlementCreate                   = Event[*EntitlementCreate]{"ENTITLEMENT_CREATE"}
	EventEntitlementUpdate                   = Event[*EntitlementUpdate]{"ENTITLEMENT_UPDATE"}
	EventEntitlementDelete                   = Event[*EntitlementDelete]{"ENTITLEMENT_DELETE"}
	EventGuildCreate                         = Event[*GuildCreate]{"GUILD_CREATE"}
	EventGuildUpdate                         = Event[*GuildUpdate]{"GUILD_UPDATE"}
	EventGuildDelete                         = Event[*GuildDelete]{"GUILD_DELETE"}
	EventGuildAuditLogEntryCreate            = Event[*GuildAuditLogEntryCreate]{"GUILD_AUDIT_LOG_ENTRY_CREATE"}
	EventGuildBanAdd                         = Event[*GuildBanAdd]{"GUILD_BAN_ADD"}
	EventGuildBanRemove                      = Event[*GuildBanRemove]{"GUILD_BAN_REMOVE"}
	EventGuildEmojisUpdate                   = Event[*GuildEmojisUpdate]{"GUILD_EMOJIS_UPDATE"}
	EventGuildStickersUpdate                 = Event[*GuildStickersUpdate]{"GUILD_STICKERS_UPDATE"}
	EventGuildIntegrationsUpdate             = Event[*GuildIntegrationsUpdate]{"GUILD_INTEGRATIONS_UPDATE"}
	EventGuildMemberAdd                      = Event[*GuildMemberAdd]{"GUILD_MEMBER_ADD"}
	EventGuildMemberRemove                   = Event[*GuildMemberRemove]{"GUILD_MEMBER_REMOVE"}
	EventGuildMemberUpdate                   = Event[*GuildMemberUpdate]{"GUILD_MEMBER_UPDATE"}
	EventGuildMembersChunk                   = Event[*GuildMembersChunk]{"GUILD_MEMBERS_CHUNK"}
	EventGuildRoleCreate                     = Event[*GuildRoleCreate]{"GUILD_ROLE_CREATE"}
	EventGuildRoleUpdate                     = Event[*GuildRoleUpdate]{"GUILD_ROLE_UPDATE"}
	EventGuildRoleDelete                     = Event[*GuildRoleDelete]{"GUILD_ROLE_DELETE"}
	EventGuildScheduledEventCreate           = Event[*GuildScheduledEventCreate]{"GUILD_SCHEDULED_EVENT_CREATE"}
	EventGuildScheduledEventUpdate           = Event[*GuildScheduledEventUpdate]{"GUILD_SCHEDULED_EVENT_UPDATE"}
	EventGuildScheduledEventDelete           = Event[*GuildScheduledEventDelete]{"GUILD_SCHEDULED_EVENT_DELETE"}
	EventGuildScheduledEventUserAdd          = Event[*GuildScheduledEventUserAdd]{"GUILD_SCHEDULED_EVENT_USER_ADD"}
	EventGuildScheduledEventUserRemove       = Event[*GuildScheduledEventUserRemove]{"GUILD_SCHEDULED_EVENT_USER_REMOVE"}
	EventGuildSoundboardSoundCreate          = Event[*GuildSoundboardSoundCreate]{"GUILD_SOUNDBOARD_SOUND_CREATE"}
	EventGuildSoundboardSoundUpdate          = Event[*GuildSoundboardSoundUpdate]{"GUILD_SOUNDBOARD_SOUND_UPDATE"}
	EventGuildSoundboardSoundDelete          = Event[*GuildSoundboardSoundDelete]{"GUILD_SOUNDBOARD_SOUND_DELETE"}
	EventGuildSoundboardSoundsUpdate         = Event[*GuildSoundboardSoundsUpdate]{"GUILD_SOUNDBOARD_SOUNDS_UPDATE"}
	EventSoundboardSounds                    = Event[*SoundboardSounds]{"SOUNDBOARD_SOUNDS"}
	EventIntegrationCreate                   = Event[*IntegrationCreate]{"INTEGRATION_CREATE"}
	EventIntegrationUpdate                   = Event[*IntegrationUpdate]{"INTEGRATION_UPDATE"}
	EventIntegrationDelete                   = Event[*IntegrationDelete]{"INTEGRATION_DELETE"}
	EventInteractionCreate                   = Event[*InteractionCreate]{"INTERACTION_CREATE"}
	EventInviteCreate                        = Event[*InviteCreate]{"INVITE_CREATE"}
	EventInviteDelete                        = Event[*InviteDelete]{"INVITE_DELETE"}
	EventMessageCreate                       = Event[*MessageCreate]{"MESSAGE_CREATE"}
	EventMessageUpdate                       = Event[*MessageUpdate]{"MESSAGE_UPDATE"}
	EventMessageDelete                       = Event[*MessageDelete]{"MESSAGE_DELETE"}
	EventMessageDeleteBulk                   = Event[*MessageDeleteBulk]{"MESSAGE_DELETE_BULK"}
	EventMessageReactionAdd                  = Event[*MessageReactionAdd]{"MESSAGE_REACTION_ADD"}
	EventMessageReactionRemove               = Event[*MessageReactionRemove]{"MESSAGE_REACTION_REMOVE"}
	EventMessageReactionRemoveAll            = Event[*MessageReactionRemoveAll]{"MESSAGE_REACTION_REMOVE_ALL"}
	EventMessageReactionRemoveEmoji          = Event[*MessageReactionRemoveEmoji]{"MESSAGE_REACTION_REMOVE_EMOJI"}
	EventPresenceUpdate                      = Event[*PresenceUpdate]{"PRESENCE_UPDATE"}
	EventStageInstanceCreate                 = Event[*StageInstanceCreate]{"STAGE_INSTANCE_CREATE"}
	EventStageInstanceUpdate                 = Event[*StageInstanceUpdate]{"STAGE_INSTANCE_UPDATE"}
	EventStageInstanceDelete                 = Event[*StageInstanceDelete]{"STAGE_INSTANCE_DELETE"}
	EventSubscriptionCreate                  = Event[*SubscriptionCreate]{"SUBSCRIPTION_CREATE"}
	EventSubscriptionUpdate                  = Event[*SubscriptionUpdate]{"SUBSCRIPTION_UPDATE"}
	EventSubscriptionDelete                  = Event[*SubscriptionDelete]{"SUBSCRIPTION_DELETE"}
	EventTypingStart                         = Event[*TypingStart]{"TYPING_START"}
	EventUserUpdate                          = Event[*UserUpdate]{"USER_UPDATE"}
	EventVoiceChannelEffectSend              = Event[*VoiceChannelEffectSend]{"VOICE_CHANNEL_EFFECT_SEND"}
	EventVoiceChannelStartTimeUpdate         = Event[*VoiceChannelStartTimeUpdate]{"VOICE_CHANNEL_START_TIME_UPDATE"}
	EventVoiceStateUpdate                    = Event[*VoiceStateUpdate]{"VOICE_STATE_UPDATE"}
	EventVoiceServerUpdate                   = Event[*VoiceServerUpdate]{"VOICE_SERVER_UPDATE"}
	EventWebhooksUpdate                      = Event[*WebhooksUpdate]{"WEBHOOKS_UPDATE"}
	EventMessagePollVoteAdd                  = Event[*MessagePollVoteAdd]{"MESSAGE_POLL_VOTE_ADD"}
	EventMessagePollVoteRemove               = Event[*MessagePollVoteRemove]{"MESSAGE_POLL_VOTE_REMOVE"}
)

type eventDispatcher func(c *Client, eventName string, data json.RawMessage)

var dispatchers = map[string]eventDispatcher{
	"READY":                                  makeDispatcher[*Ready](),
	"RESUMED":                                makeDispatcher[*Resumed](),
	"RECONNECT":                              makeDispatcher[*Reconnect](),
	"RATE_LIMITED":                           makeDispatcher[*RateLimited](),
	"INVALID_SESSION":                        makeDispatcher[*InvalidSession](),
	"APPLICATION_COMMAND_PERMISSIONS_UPDATE": makeDispatcher[*ApplicationCommandPermissionsUpdate](),
	"AUTO_MODERATION_RULE_CREATE":            makeDispatcher[*AutoModerationRuleCreate](),
	"AUTO_MODERATION_RULE_UPDATE":            makeDispatcher[*AutoModerationRuleUpdate](),
	"AUTO_MODERATION_RULE_DELETE":            makeDispatcher[*AutoModerationRuleDelete](),
	"AUTO_MODERATION_ACTION_EXECUTION":       makeDispatcher[*AutoModerationActionExecution](),
	"CHANNEL_CREATE":                         makeDispatcher[*ChannelCreate](),
	"CHANNEL_UPDATE":                         makeDispatcher[*ChannelUpdate](),
	"CHANNEL_DELETE":                         makeDispatcher[*ChannelDelete](),
	"CHANNEL_INFO":                           makeDispatcher[*ChannelInfo](),
	"CHANNEL_PINS_UPDATE":                    makeDispatcher[*ChannelPinsUpdate](),
	"THREAD_CREATE":                          makeDispatcher[*ThreadCreate](),
	"THREAD_UPDATE":                          makeDispatcher[*ThreadUpdate](),
	"THREAD_DELETE":                          makeDispatcher[*ThreadDelete](),
	"THREAD_LIST_SYNC":                       makeDispatcher[*ThreadListSync](),
	"THREAD_MEMBER_UPDATE":                   makeDispatcher[*ThreadMemberUpdate](),
	"THREAD_MEMBERS_UPDATE":                  makeDispatcher[*ThreadMembersUpdate](),
	"ENTITLEMENT_CREATE":                     makeDispatcher[*EntitlementCreate](),
	"ENTITLEMENT_UPDATE":                     makeDispatcher[*EntitlementUpdate](),
	"ENTITLEMENT_DELETE":                     makeDispatcher[*EntitlementDelete](),
	"GUILD_CREATE":                           makeDispatcher[*GuildCreate](),
	"GUILD_UPDATE":                           makeDispatcher[*GuildUpdate](),
	"GUILD_DELETE":                           makeDispatcher[*GuildDelete](),
	"GUILD_AUDIT_LOG_ENTRY_CREATE":           makeDispatcher[*GuildAuditLogEntryCreate](),
	"GUILD_BAN_ADD":                          makeDispatcher[*GuildBanAdd](),
	"GUILD_BAN_REMOVE":                       makeDispatcher[*GuildBanRemove](),
	"GUILD_EMOJIS_UPDATE":                    makeDispatcher[*GuildEmojisUpdate](),
	"GUILD_STICKERS_UPDATE":                  makeDispatcher[*GuildStickersUpdate](),
	"GUILD_INTEGRATIONS_UPDATE":              makeDispatcher[*GuildIntegrationsUpdate](),
	"GUILD_MEMBER_ADD":                       makeDispatcher[*GuildMemberAdd](),
	"GUILD_MEMBER_REMOVE":                    makeDispatcher[*GuildMemberRemove](),
	"GUILD_MEMBER_UPDATE":                    makeDispatcher[*GuildMemberUpdate](),
	"GUILD_MEMBERS_CHUNK":                    makeDispatcher[*GuildMembersChunk](),
	"GUILD_ROLE_CREATE":                      makeDispatcher[*GuildRoleCreate](),
	"GUILD_ROLE_UPDATE":                      makeDispatcher[*GuildRoleUpdate](),
	"GUILD_ROLE_DELETE":                      makeDispatcher[*GuildRoleDelete](),
	"GUILD_SCHEDULED_EVENT_CREATE":           makeDispatcher[*GuildScheduledEventCreate](),
	"GUILD_SCHEDULED_EVENT_UPDATE":           makeDispatcher[*GuildScheduledEventUpdate](),
	"GUILD_SCHEDULED_EVENT_DELETE":           makeDispatcher[*GuildScheduledEventDelete](),
	"GUILD_SCHEDULED_EVENT_USER_ADD":         makeDispatcher[*GuildScheduledEventUserAdd](),
	"GUILD_SCHEDULED_EVENT_USER_REMOVE":      makeDispatcher[*GuildScheduledEventUserRemove](),
	"GUILD_SOUNDBOARD_SOUND_CREATE":          makeDispatcher[*GuildSoundboardSoundCreate](),
	"GUILD_SOUNDBOARD_SOUND_UPDATE":          makeDispatcher[*GuildSoundboardSoundUpdate](),
	"GUILD_SOUNDBOARD_SOUND_DELETE":          makeDispatcher[*GuildSoundboardSoundDelete](),
	"GUILD_SOUNDBOARD_SOUNDS_UPDATE":         makeDispatcher[*GuildSoundboardSoundsUpdate](),
	"SOUNDBOARD_SOUNDS":                      makeDispatcher[*SoundboardSounds](),
	"INTEGRATION_CREATE":                     makeDispatcher[*IntegrationCreate](),
	"INTEGRATION_UPDATE":                     makeDispatcher[*IntegrationUpdate](),
	"INTEGRATION_DELETE":                     makeDispatcher[*IntegrationDelete](),
	"INTERACTION_CREATE":                     makeDispatcher[*InteractionCreate](),
	"INVITE_CREATE":                          makeDispatcher[*InviteCreate](),
	"INVITE_DELETE":                          makeDispatcher[*InviteDelete](),
	"MESSAGE_CREATE":                         makeDispatcher[*MessageCreate](),
	"MESSAGE_UPDATE":                         makeDispatcher[*MessageUpdate](),
	"MESSAGE_DELETE":                         makeDispatcher[*MessageDelete](),
	"MESSAGE_DELETE_BULK":                    makeDispatcher[*MessageDeleteBulk](),
	"MESSAGE_REACTION_ADD":                   makeDispatcher[*MessageReactionAdd](),
	"MESSAGE_REACTION_REMOVE":                makeDispatcher[*MessageReactionRemove](),
	"MESSAGE_REACTION_REMOVE_ALL":            makeDispatcher[*MessageReactionRemoveAll](),
	"MESSAGE_REACTION_REMOVE_EMOJI":          makeDispatcher[*MessageReactionRemoveEmoji](),
	"PRESENCE_UPDATE":                        makeDispatcher[*PresenceUpdate](),
	"STAGE_INSTANCE_CREATE":                  makeDispatcher[*StageInstanceCreate](),
	"STAGE_INSTANCE_UPDATE":                  makeDispatcher[*StageInstanceUpdate](),
	"STAGE_INSTANCE_DELETE":                  makeDispatcher[*StageInstanceDelete](),
	"SUBSCRIPTION_CREATE":                    makeDispatcher[*SubscriptionCreate](),
	"SUBSCRIPTION_UPDATE":                    makeDispatcher[*SubscriptionUpdate](),
	"SUBSCRIPTION_DELETE":                    makeDispatcher[*SubscriptionDelete](),
	"TYPING_START":                           makeDispatcher[*TypingStart](),
	"USER_UPDATE":                            makeDispatcher[*UserUpdate](),
	"VOICE_CHANNEL_EFFECT_SEND":              makeDispatcher[*VoiceChannelEffectSend](),
	"VOICE_CHANNEL_START_TIME_UPDATE":        makeDispatcher[*VoiceChannelStartTimeUpdate](),
	"VOICE_STATE_UPDATE":                     makeDispatcher[*VoiceStateUpdate](),
	"VOICE_SERVER_UPDATE":                    makeDispatcher[*VoiceServerUpdate](),
	"WEBHOOKS_UPDATE":                        makeDispatcher[*WebhooksUpdate](),
	"MESSAGE_POLL_VOTE_ADD":                  makeDispatcher[*MessagePollVoteAdd](),
	"MESSAGE_POLL_VOTE_REMOVE":               makeDispatcher[*MessagePollVoteRemove](),
}

func makeDispatcher[T any]() eventDispatcher {
	return func(c *Client, eventName string, data json.RawMessage) {
		dispatch[T](c, eventName, data)
	}
}

func dispatch[T any](c *Client, eventName string, data json.RawMessage) {
	handlers, ok := c.handlers[eventName]

	if !ok {
		c.logger.Debug("Unable to find valid handlers for event", "event", eventName)
		return
	}

	var payload T

	if err := json.Unmarshal(data, &payload); err != nil {
		c.logger.Error("Failed to unmarshal event payload", "event", eventName, "err", err)
		return
	}

	for _, h := range handlers {
		handler, ok := h.(func(*Client, T))

		if !ok {
			continue
		}

		go handler(c, payload)
	}
}

func On[T any](c *Client, event Event[T], handler func(*Client, T)) {
	c.handlers[event.name] = append(c.handlers[event.name], handler)
}
