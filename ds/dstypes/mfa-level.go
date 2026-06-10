package dstypes

import "cynthia/util"

type MfaLevel int

const (
	MfaLevelNone     MfaLevel = 0
	MfaLevelElevated MfaLevel = 1
)

func (m *MfaLevel) UnmarshalJSON(data []byte) error {
	return util.UnmarshalNumeric(data, m)
}
