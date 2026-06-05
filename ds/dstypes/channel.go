package dstypes

type Channel struct {
	ID                            Snowflake        `json:"id"`
	Type                          ChannelType      `json:"type"`
	GuildID                       *Snowflake       `json:"guild_id"`
	Position                      *int             `json:"position"`
	PermissionOverwrites          *[]Overwrite     `json:"permission_overwrites"`
	Name                          *string          `json:"name"`
	Topic                         *string          `json:"topic"`
	NSFW                          *bool            `json:"nsfw"`
	LastMessageID                 *Snowflake       `json:"last_message_id"`
	Bitrate                       *int             `json:"bitrate"`
	UserLimit                     *int             `json:"user_limit"`
	RateLimitPerUser              *int             `json:"rate_limit_per_user"`
	Recipients                    *[]User          `json:"recipients"`
	Icon                          *string          `json:"icon"`
	OwnerID                       *Snowflake       `json:"owner_id"`
	ApplicationID                 *Snowflake       `json:"application_id"`
	Managed                       *bool            `json:"managed"`
	ParentID                      *Snowflake       `json:"parent_id"`
	LastPinTimestamp              *string          `json:"last_pin_timestamp"`
	RtcRegion                     *string          `json:"rtc_region"`
	VideoQualityMode              *int             `json:"video_quality_mode"`
	MessageCount                  *int             `json:"message_count"`
	MemberCount                   *int             `json:"member_count"`
	ThreadMetadata                *ThreadMetadata  `json:"thread_metadata"`
	Member                        *ThreadMember    `json:"member"`
	DefaultAutoArchiveDuration    *int             `json:"default_auto_archive_duration"`
	Permissions                   *string          `json:"permissions"`
	Flags                         *int             `json:"flags"`
	TotalMessageSent              *int             `json:"total_message_sent"`
	AvailableTags                 *[]ForumTag      `json:"available_tags"`
	AppliedTags                   *[]Snowflake     `json:"applied_tags"`
	DefaultReactionEmoji          *DefaultReaction `json:"default_reaction_emoji"`
	DefaultThreadRateLimitPerUser *int             `json:"default_thread_rate_limit_per_user"`
	DefaultSortOrder              *int             `json:"default_sort_order"`
	DefaultForumLayout            *int             `json:"default_forum_layout"`
}
