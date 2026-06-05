package payloads

import "cynthia/ds/dstypes"

type ChannelInfo struct {
	GuildID  dstypes.Snowflake
	Channels []dstypes.ChannelInfoChannel
}
