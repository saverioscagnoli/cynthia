package ds

import "cynthia/util"

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

type ChannelMentions struct {
	ID      Snowflake `json:"id"`
	GuildID Snowflake `json:"guild_id"`
	Type    int       `json:"type"`
	Name    string    `json:"name"`
}

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

func (c *ChannelType) UnmarshalJSON(b []byte) error {
	return util.UnmarshalNumeric(b, c)
}

type ForumTag struct {
	ID        Snowflake  `json:"id"`
	Name      string     `json:"name"`
	Moderated bool       `json:"moderated"`
	EmojiID   *Snowflake `json:"emoji_id"`
	EmojiName *string    `json:"emoji_name"`
}

type Overwrite struct {
	ID    Snowflake `json:"id"`
	Type  int       `json:"type"`
	Allow string    `json:"allow"`
	Deny  string    `json:"deny"`
}

type SystemChannelFlags int

const (
	SystemChannelFlagSuppressJoinNotifications                                            SystemChannelFlags = 1 << 0
	SystemChannelFlagSuppressPremiumSubscriptions                                         SystemChannelFlags = 1 << 1
	SystemChannelFlagSuppressGuildReminderNotifications                                   SystemChannelFlags = 1 << 2
	SystemChannelFlagSuppressJoinNotificationReplies                                      SystemChannelFlags = 1 << 3
	SystemChannelFlagSuppressRoleSubscriptionPurchaseNotifications                        SystemChannelFlags = 1 << 4
	SystemChannelFlagSuppressRoleSubscriptionPurchaseNotificationRepliesSystemChannelFlag SystemChannelFlags = 1 << 5
)

func (t *SystemChannelFlags) UnmarshalJSON(data []byte) error {
	return util.UnmarshalNumeric(data, t)
}

type ThreadMetadata struct {
	Archived            bool    `json:"archived"`
	AutoArchiveDuration int     `json:"auto_archive_duration"`
	ArchiveTimestamp    string  `json:"archive_timestamp"`
	Locked              bool    `json:"locked"`
	Invitable           *bool   `json:"invitable"`
	CreateTimestamp     *string `json:"create_timestamp"`
}

type ThreadMember struct {
	ID            *Snowflake   `json:"id"`
	UserID        *Snowflake   `json:"user_id"`
	JoinTimestamp string       `json:"join_timestamp"`
	Flags         int          `json:"flags"`
	Member        *GuildMember `json:"member"`
}

type WelcomeScreenChannel struct {
	ChannelID   Snowflake  `json:"channel_id"`
	Description string     `json:"description"`
	EmojiID     *Snowflake `json:"emoji_id"`
	EmojiName   *string    `json:"emoji_name"`
}
