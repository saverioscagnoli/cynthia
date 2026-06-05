package payloads

import "cynthia/ds/dstypes"

type ThreadListSync struct {
	GuildID    dstypes.Snowflake      `json:"guild_id"`
	ChannelIDs *[]dstypes.Snowflake   `json:"channel_ids"`
	Threads    []dstypes.Channel      `json:"threads"`
	Members    []dstypes.ThreadMember `json:"members"`
}
