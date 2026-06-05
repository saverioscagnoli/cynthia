package dstypes

type EntitlementType int

const (
	EntitlementTypePurchase                EntitlementType = 1
	EntitlementTypePremiumSubscription     EntitlementType = 2
	EntitlementTypeDeveloperGift           EntitlementType = 3
	EntitlementTypeTestModePurchase        EntitlementType = 4
	EntitlementTypeFreePurchase            EntitlementType = 5
	EntitlementTypeUserGift                EntitlementType = 6
	EntitlementTypePremiumPurchase         EntitlementType = 7
	EntitlementTypeApplicationSubscription EntitlementType = 8
)
