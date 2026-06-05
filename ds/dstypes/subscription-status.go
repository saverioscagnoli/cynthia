package dstypes

type SubscriptionStatus int

const (
	SubscriptionStatusActive   SubscriptionStatus = 0
	SubscriptionStatusEnding   SubscriptionStatus = 1
	SubscriptionStatusInactive SubscriptionStatus = 2
)
