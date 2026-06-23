package ds

type AutoModerationRuleCreate = AutoModerationRule

type AutoModerationRuleUpdate = AutoModerationRule

type AutoModerationRuleDelete = AutoModerationRule

type AutoModerationActionExecution struct {
	GuildID              Snowflake            `json:"guild_id"`
	Action               AutoModerationAction `json:"action"`
	RuleID               Snowflake            `json:"rule_id"`
	RuleTriggerType      RuleTriggerType      `json:"rule_trigger_type"`
	UserID               Snowflake            `json:"user_id"`
	ChannelID            *Snowflake           `json:"channel_id"`
	MessageID            *Snowflake           `json:"message_id"`
	AlertSystemMessageID *Snowflake           `json:"alert_system_message_id"`
	Content              string               `json:"content"`
	MatchedKeyword       *string              `json:"matched_keyword"`
	MatchedContent       *string              `json:"matched_content"`
}

type GuildCreate struct {
	Guild
	JoinedAt             string                `json:"joined_at"`
	Unavailable          *bool                 `json:"unavailable"`
	MemberCount          int                   `json:"member_count"`
	VoiceStates          []VoiceState          `json:"voice_states"`
	Members              []GuildMember         `json:"members"`
	Channels             []Channel             `json:"channels"`
	Threads              []Channel             `json:"threads"`
	Presences            []PresenceUpdate      `json:"presences"`
	StageInstances       []StageInstance       `json:"stage_instances"`
	GuildScheduledEvents []GuildScheduledEvent `json:"guild_scheduled_events"`
	SoundboardSounts     []SoundboardSound     `json:"soundboard_sounds"`
}

type GuildUpdate = Guild

type GuildDelete = Guild

type GuildAuditLogEntryCreate struct {
	AuditLogEntry
	GuildID Snowflake `json:"guild_id"`
}

type GuildBanAdd struct {
	GuildID Snowflake `json:"guild_id"`
	User    User      `json:"user"`
}

type GuildBanRemove struct {
	GuildID Snowflake `json:"guild_id"`
	User    User      `json:"user"`
}

type GuildEmojisUpdate struct {
	GuildID Snowflake `json:"guild_id"`
	Emojis  []Emoji   `json:"emojis"`
}

type GuildStickersUpdate struct {
	GuildID  Snowflake `json:"guild_id"`
	Stickers []Sticker `json:"stickers"`
}

type GuildIntegrationsUpdate struct {
	GuildID Snowflake `json:"guild_id"`
}

type GuildMemberAdd struct {
	GuildMember
	GuildId Snowflake `json:"guild_id"`
}

type GuildMemberRemove struct {
	GuildID Snowflake `json:"guild_id"`
	User    User      `json:"user"`
}

type GuildMemberUpdate struct {
	GuildID                    Snowflake             `json:"guild_id"`
	Roles                      []Snowflake           `json:"roles"`
	User                       User                  `json:"user"`
	Nick                       *string               `json:"nick"`
	Avatar                     *string               `json:"avatar"`
	Banner                     *string               `json:"banner"`
	JoinedAt                   *string               `json:"joined_at"`
	PremiumSince               *string               `json:"premium_since"`
	Deaf                       *bool                 `json:"deaf"`
	Mute                       *bool                 `json:"mute"`
	Pending                    *bool                 `json:"pending"`
	CommunicationDisabledUntil *string               `json:"communication_disabled_until"`
	AvatarDecorationData       *AvatarDecorationData `json:"avatar_decoration_data"`
	Collectibles               *Collectibles         `json:"collectibles"`
}

type GuildMembersChunk struct {
	GuildID    Snowflake     `json:"guild_id"`
	Members    []GuildMember `json:"members"`
	ChunkIndex int           `json:"chunk_index"`
	ChunkCount int           `json:"chunk_count"`
	NotFound   *[]any        `json:"not_found"`
	Presences  *[]Presence   `json:"presences"`
	Nonce      *string       `json:"nonce"`
}

type GuildRoleCreate struct {
	GuildID Snowflake `json:"guild_id"`
	Role    Role      `json:"role"`
}

type GuildRoleUpdate struct {
	GuildID Snowflake `json:"guild_id"`
	Role    Role      `json:"role"`
}

type GuildRoleDelete struct {
	GuildID Snowflake `json:"guild_id"`
	RoleID  Snowflake `json:"role_id"`
}

type GuildScheduledEventCreate = GuildScheduledEvent

type GuildScheduledEventUpdate = GuildScheduledEvent

type GuildScheduledEventDelete = GuildScheduledEvent

type GuildScheduledEventUserAdd struct {
	GuildScheduledEventID Snowflake `json:"guild_scheduled_event_id"`
	UserID                Snowflake `json:"user_id"`
	GuildID               Snowflake `json:"guild_id"`
}

type GuildScheduledEventUserRemove struct {
	GuildScheduledEventUserID Snowflake `json:"guild_scheduled_event_id"`
	UserID                    Snowflake `json:"user_id"`
	GuildID                   Snowflake `json:"guild_id"`
}

type GuildSoundboardSoundCreate = SoundboardSound

type GuildSoundboardSoundUpdate = SoundboardSound

type GuildSoundboardSoundDelete struct {
	SoundID Snowflake `json:"sound_id"`
	GuildID Snowflake `json:"guild_id"`
}

type GuildSoundboardSoundsUpdate struct {
	SoundboardSounds []SoundboardSound `json:"soundboard_sounds"`
	GuildID          Snowflake         `json:"guild_id"`
}

type IntegrationCreate struct {
	Integration
	GuildID Snowflake `json:"guild_id"`
}

type IntegrationUpdate struct {
	Integration
	GuildID Snowflake `json:"guild_id"`
}

type IntegrationDelete struct {
	ID            Snowflake  `json:"id"`
	GuildID       Snowflake  `json:"guild_id"`
	ApplicationID *Snowflake `json:"application_id"`
}

type InviteCreate struct {
	ChannelID         Snowflake    `json:"channel_id"`
	Code              string       `json:"code"`
	CreatedAt         string       `json:"created_at"`
	GuildID           *Snowflake   `json:"guild_id"`
	Inviter           *User        `json:"inviter"`
	MaxAge            int          `json:"max_age"`
	MaxUses           int          `json:"max_uses"`
	TargetType        InviteTarget `json:"target_type"`
	TargetUser        *User        `json:"target_user"`
	TargetApplication *Application `json:"target_application"`
	Temporary         bool         `json:"temporary"`
	Uses              int          `json:"uses"`
	ExpiresAt         *string      `json:"expires_at"`
	RoleIDs           *[]Snowflake `json:"role_ids"`
}

type InviteDelete struct {
	ChannelID Snowflake  `json:"channel_id"`
	GuildID   *Snowflake `json:"guild_id"`
	Code      string     `json:"code"`
}
