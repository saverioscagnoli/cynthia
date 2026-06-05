package dstypes

type GuildAgeRestrictionLevel int

const (
	GuildAgeRestrictionLevelDefault       GuildAgeRestrictionLevel = 0
	GuildAgeRestrictionLevelExplicit      GuildAgeRestrictionLevel = 1
	GuildAgeRestrictionLevelSafe          GuildAgeRestrictionLevel = 2
	GuildAgeRestructionLevelAgeRestricted GuildAgeRestrictionLevel = 3
)
