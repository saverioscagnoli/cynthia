package dstypes

import "cynthia/util"

type GuildAgeRestrictionLevel int

const (
	GuildAgeRestrictionLevelDefault       GuildAgeRestrictionLevel = 0
	GuildAgeRestrictionLevelExplicit      GuildAgeRestrictionLevel = 1
	GuildAgeRestrictionLevelSafe          GuildAgeRestrictionLevel = 2
	GuildAgeRestructionLevelAgeRestricted GuildAgeRestrictionLevel = 3
)

func (e *GuildAgeRestrictionLevel) UnmarshalJSON(b []byte) error {
	return util.UnmarshalNumeric(b, e)
}
