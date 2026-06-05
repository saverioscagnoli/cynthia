package dstypes

type ExplicitContentFilterLevel int

const (
	ExplicitContentFilterLevelDisabled            = 0
	ExplicitContentFilterLevelMembersWithoutRoles = 1
	ExplicitContentFilterLevelAllMembers          = 2
)
