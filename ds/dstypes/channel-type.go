package dstypes

type ChannelType int

const (
	ChannelTypeGuildText          = 0
	ChannelTypeDM                 = 1
	ChannelTypeGuildVoice         = 2
	ChannelTypeGroupDM            = 3
	ChannelTypeGuildCategory      = 4
	ChannelTypeGuildAnnouncement  = 5
	ChannelTypeAnnouncementThread = 10
	ChannelTypePublicThread       = 11
	ChannelTypePrivateThread      = 12
	ChannelTypeGuildstageVoice    = 13
	ChannelTypeGuildDirectory     = 14
	ChannelTypeGuildForum         = 15
	ChannelTypeGuildMedia         = 16
)
