package ds

import "cynthia/util"

type Intent uint64

const (
	IntentGuilds                      Intent = 1 << 0
	IntentGuildMembers                Intent = 1 << 1
	IntentGuildModeration             Intent = 1 << 2
	IntentGuildExpression             Intent = 1 << 3
	IntentGuildIntegrations           Intent = 1 << 4
	IntentGuildWebHooks               Intent = 1 << 5
	IntentGuildInvites                Intent = 1 << 6
	IntentGuildVoiceStates            Intent = 1 << 7
	IntentGuildPresences              Intent = 1 << 8
	IntentGuildMessages               Intent = 1 << 9
	IntentGuildMessageReactions       Intent = 1 << 10
	IntentGuildMessageTyping          Intent = 1 << 11
	IntentDirectMessages              Intent = 1 << 12
	IntentDirectMessageReactions      Intent = 1 << 13
	IntentDirectMessageTyping         Intent = 1 << 14
	IntentMessageContent              Intent = 1 << 15
	IntentGuildScheduledEvents        Intent = 1 << 16
	IntentAutoModerationConfiguration Intent = 1 << 20
	IntentAutoModerationExecution     Intent = 1 << 21
	IntentGuildMessagePolls           Intent = 1 << 24
	IntentDirectMessagePolls          Intent = 1 << 25

	IntentAll Intent = IntentGuilds | IntentGuildMembers | IntentGuildModeration |
		IntentGuildExpression | IntentGuildIntegrations | IntentGuildWebHooks |
		IntentGuildInvites | IntentGuildVoiceStates | IntentGuildPresences |
		IntentGuildMessages | IntentGuildMessageReactions | IntentGuildMessageTyping |
		IntentDirectMessages | IntentDirectMessageReactions | IntentDirectMessageTyping |
		IntentMessageContent | IntentGuildScheduledEvents | IntentAutoModerationConfiguration |
		IntentAutoModerationExecution | IntentGuildMessagePolls | IntentDirectMessagePolls
)

func (i *Intent) UnmarshalJSON(b []byte) error {
	return util.UnmarshalNumeric(b, i)
}

func (i Intent) Has(intent Intent) bool {
	return i&intent == intent
}

func (i Intent) Add(intent Intent) Intent {
	return i | intent
}

func (i Intent) Remove(intent Intent) Intent {
	return i &^ intent
}
