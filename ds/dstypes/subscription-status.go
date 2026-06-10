package dstypes

import "cynthia/util"

type SubscriptionStatus int

const (
	SubscriptionStatusActive   SubscriptionStatus = 0
	SubscriptionStatusEnding   SubscriptionStatus = 1
	SubscriptionStatusInactive SubscriptionStatus = 2
)

func (t *SubscriptionStatus) UnmarshalJSON(data []byte) error {
	return util.UnmarshalNumeric(data, t)
}
