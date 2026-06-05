package payloads

import "cynthia/dstypes"

type MessageCreate struct {
	dstypes.Message
	GuildID     *dstypes.Snowflake
	Member      *dstypes.GuildMember
	Mentions    *[]dstypes.User
	ChannelType *dstypes.ChannelType
}
