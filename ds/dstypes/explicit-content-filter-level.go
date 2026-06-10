package dstypes

import "cynthia/util"

type ExplicitContentFilterLevel int

const (
	ExplicitContentFilterLevelDisabled            = 0
	ExplicitContentFilterLevelMembersWithoutRoles = 1
	ExplicitContentFilterLevelAllMembers          = 2
)

func (e *ExplicitContentFilterLevel) UnmarshalJSON(b []byte) error {
	return util.UnmarshalNumeric(b, e)
}
