package dstypes

import "cynthia/util"

type PremiumTier int

const (
	PremiumTierNone PremiumTier = 0
	PremiumTier1    PremiumTier = 1
	PremiumTier2    PremiumTier = 2
	PremiumTier3    PremiumTier = 3
)

func (t *PremiumTier) UnmarshalJSON(data []byte) error {
	return util.UnmarshalNumeric(data, t)
}
