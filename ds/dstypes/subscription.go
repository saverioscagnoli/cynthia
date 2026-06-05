package dstypes

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
