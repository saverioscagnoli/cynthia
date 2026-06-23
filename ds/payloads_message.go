package ds

import "encoding/json"

type MessageCreate struct {
	Message
	GuildID     *Snowflake          `json:"guild_id"`
	Member      *GuildMember        `json:"member"`
	Mentions    *[]User             `json:"mentions"`
	ChannelType *ChannelType        `json:"channel_type"`
	Components  *[]MessageComponent `json:"components"`
}

func (mc *MessageCreate) UnmarshalJSON(data []byte) error {
	type Alias MessageCreate

	var raw struct {
		Alias
		Components []json.RawMessage `json:"components"`
	}

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	*mc = MessageCreate(raw.Alias)

	if raw.Components != nil {
		var err error

		mc.Components, err = unmarshalComponentsPtr(raw.Components)

		if err != nil {
			return err
		}
	}

	return nil
}

type MessageUpdate = Message

type MessageDelete struct {
	ID        Snowflake  `json:"id"`
	ChannelID Snowflake  `json:"channel_id"`
	GuildID   *Snowflake `json:"guild_id"`
}

type MessageDeleteBulk struct {
	IDs       []Snowflake `json:"ids"`
	ChannelID Snowflake   `json:"channel_id"`
	GuildID   *Snowflake  `json:"guild_id"`
}

type MessageReactionAdd struct {
	UserID          Snowflake    `json:"user_id"`
	ChannelID       Snowflake    `json:"channel_id"`
	MessageID       Snowflake    `json:"message_id"`
	GuildID         *Snowflake   `json:"guild_id"`
	Member          *GuildMember `json:"member"`
	Emoji           Emoji        `json:"emoji"`
	MessageAuthorID *Snowflake   `json:"message_author_id"`
	Burst           bool         `json:"burst"`
	BurstColors     *[]string    `json:"burst_colors"`
	Type            ReactionType `json:"type"`
}

type MessageReactionRemove struct {
	UserID    Snowflake    `json:"user_id"`
	ChannelID Snowflake    `json:"channel_id"`
	MessageID Snowflake    `json:"message_id"`
	GuildID   *Snowflake   `json:"guild_id"`
	Emoji     *Emoji       `json:"emoji"`
	Burst     bool         `json:"burst"`
	Type      ReactionType `json:"type"`
}

type MessageReactionRemoveAll struct {
	ChannelID Snowflake  `json:"channel_id"`
	MessageID Snowflake  `json:"message_id"`
	GuildID   *Snowflake `json:"guild_id"`
}

type MessageReactionRemoveEmoji struct {
	ChannelID Snowflake  `json:"channel_id"`
	GuildID   *Snowflake `json:"guild_id"`
	MessageID Snowflake  `json:"message_id"`
	Emoji     Emoji      `json:"emoji"`
}

type MessagePollVoteAdd struct {
	UserID    Snowflake  `json:"user_id"`
	ChannelID Snowflake  `json:"channel_id"`
	MessageID Snowflake  `json:"message_id"`
	GuildID   *Snowflake `json:"guild_id"`
	AnswerID  int        `json:"answer_id"`
}

type MessagePollVoteRemove struct {
	UserID    Snowflake  `json:"user_id"`
	ChannelID Snowflake  `json:"channel_id"`
	MessageID Snowflake  `json:"message_id"`
	GuildID   *Snowflake `json:"guild_id"`
	AnswerID  int
}
