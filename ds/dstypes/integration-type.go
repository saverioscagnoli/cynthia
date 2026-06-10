package dstypes

import "cynthia/util"

type IntegrationType string

const (
	IntegrationTypeDiscord           IntegrationType = "discord"
	IntegrationTypeGuildSubscription IntegrationType = "guild_subscription"
	IntegrationTypeTwitch            IntegrationType = "twitch"
	IntegrationTypeYoutube           IntegrationType = "youtupe"
)

func (e *IntegrationType) UnmarshalJSON(b []byte) error {
	return util.UnmarshalString(b, e)
}
