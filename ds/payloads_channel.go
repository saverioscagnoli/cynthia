package ds

type ChannelCreate = Channel

type ChannelUpdate = Channel

type ChannelDelete = Channel

type ChannelInfo struct {
	GuildID  Snowflake
	Channels []ChannelInfoChannel
}

type ChannelPinsUpdate struct {
	GuildID          *Snowflake `json:"guild_id"`
	ChannelID        Snowflake  `json:"channel_id"`
	LastPinTimestamp *string    `json:"last_pin_timestamp"`
}

type ThreadCreate = Channel

type ThreadUpdate = Channel

type ThreadDelete = Channel

type ThreadListSync struct {
	GuildID    Snowflake      `json:"guild_id"`
	ChannelIDs *[]Snowflake   `json:"channel_ids"`
	Threads    []Channel      `json:"threads"`
	Members    []ThreadMember `json:"members"`
}

type ThreadMemberUpdate struct {
	ThreadMember
	GuildID Snowflake `json:"guild_id"`
}

type ThreadMembersUpdate struct {
	ID               Snowflake       `json:"id"`
	GuildID          Snowflake       `json:"guild_id"`
	MemberCount      int             `json:"member_count"`
	AddedMembers     *[]ThreadMember `json:"added_members"`
	RemovedMemberIDs *[]Snowflake    `json:"removed_member_ids"`
}

type TypingStart struct {
	ChannelID Snowflake    `json:"channel_id"`
	GuildID   *Snowflake   `json:"guild_id"`
	UserID    Snowflake    `json:"user_id"`
	Timestamp int          `json:"timestamp"`
	Member    *GuildMember `json:"member"`
}

type VoiceChannelEffectSend struct {
	ChannelID     Snowflake      `json:"channel_id"`
	GuildID       Snowflake      `json:"guild_id"`
	UserID        Snowflake      `json:"user_id"`
	Emoji         *Emoji         `json:"emoji"`
	AnimationType *AnimationType `json:"animation_type"`
	AnimationID   *int           `json:"animation_id"`
	SoundID       *Snowflake     `json:"sound_id"`
	SoundVolume   float64        `json:"sound_volume"`
}

type VoiceChannelStartTimeUpdate struct {
	ID             Snowflake `json:"id"`
	GuildID        Snowflake `json:"guild_id"`
	VoiceStartTime *int      `json:"voice_start_time"`
}
