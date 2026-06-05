package dstypes

type IntegrationType string

const (
	IntegrationTypeDiscord           IntegrationType = "discord"
	IntegrationTypeGuildSubscription IntegrationType = "guild_subscription"
	IntegrationTypeTwitch            IntegrationType = "twitch"
	IntegrationTypeYoutube           IntegrationType = "youtupe"
)
