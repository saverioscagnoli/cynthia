package dstypes

type ActivityFlags int

const (
	ActivityFlagInstance                 = 1 << 0
	ActivityFlagJoin                     = 1 << 1
	ActivityFlagSpectate                 = 1 << 2
	ActivityFlagJoinRequest              = 1 << 3
	ActivityFlagSync                     = 1 << 4
	ActivityFlagPlay                     = 1 << 5
	ActivityFlagPartyPrivacyFriends      = 1 << 6
	ActivityFlagPartyPrivacyVoiceChannel = 1 << 7
	ActivityFlagEmbedded                 = 1 << 8
)
