package dstypes

import "cynthia/util"

type VerificationLevel int

const (
	VerificationLevelNone     VerificationLevel = 0
	VerificationLevelLow      VerificationLevel = 1
	VerificationLevelMedium   VerificationLevel = 2
	VerificationLevelHigh     VerificationLevel = 3
	VerificationLevelVeryHigh VerificationLevel = 4
)

func (t *VerificationLevel) UnmarshalJSON(data []byte) error {
	return util.UnmarshalNumeric(data, t)
}
