package dstypes

import "cynthia/util"

type StagePrivacyLevel int

const (
	StagePrivacyLevelPublic    StagePrivacyLevel = 1
	StagePrivacyLevelGuildOnly StagePrivacyLevel = 2
)

func (t *StagePrivacyLevel) UnmarshalJSON(data []byte) error {
	return util.UnmarshalNumeric(data, t)
}
