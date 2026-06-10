package dstypes

import "cynthia/util"

type Intents int

const (
	IntentsGuilds                      Intents = 1 << 0
	IntentsGuildMembers                Intents = 1 << 1
	IntentsGuildModeration             Intents = 1 << 2
	IntentsGuildExpression             Intents = 1 << 3
	IntentsGuildIntegrations           Intents = 1 << 4
	IntentsGuildWebHooks               Intents = 1 << 5
	IntentsGuildInvites                Intents = 1 << 6
	IntentsGuildVoiceStates            Intents = 1 << 7
	IntentsGuildPresences              Intents = 1 << 8
	IntentsGuildMessages               Intents = 1 << 9
	IntentsGuildMessageReactions       Intents = 1 << 10
	IntentsGuildMessageTyping          Intents = 1 << 11
	IntentsDirectMessages              Intents = 1 << 12
	IntentsDirectMessageReactions      Intents = 1 << 13
	IntentsDirectMessageTyping         Intents = 1 << 14
	IntentsMessageContent              Intents = 1 << 15
	IntentsGuildScheduledEvents        Intents = 1 << 16
	IntentsAutoModerationConfiguration Intents = 1 << 20
	IntentsAutoModerationExecution     Intents = 1 << 21
	IntentsGuildMessagePolls           Intents = 1 << 24
	IntentsDirectMessagePolls          Intents = 1 << 25
)

func (e *Intents) UnmarshalJSON(b []byte) error {
	return util.UnmarshalNumeric(b, e)
}
