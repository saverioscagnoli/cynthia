package payloads

import "cynthia/dstypes"

type ChannelInfo struct {
	GuildID  dstypes.Snowflake
	Channels []dstypes.ChannelInfoChannel
}
