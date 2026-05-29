package ds

type RoleTags struct {
	BotID                 Snowflake `json:"bot_id"`
	IntegrationID         Snowflake `json:"integration_id"`
	PremiumSubscriber     *bool     `json:"premium_subscriber"`
	SubscriptionListingID Snowflake `json:"subscription_listing_id"`
	AvailableForPurchase  *bool     `json:"available_for_purchase"`
	GuildConnections      *bool     `json:"guild_connections"`
}
