package payloads

import "cynthia/dstypes"

type MessageReactionAdd struct {
	UserID          dstypes.Snowflake    `json:"user_id"`
	ChannelID       dstypes.Snowflake    `json:"channel_id"`
	MessageID       dstypes.Snowflake    `json:"message_id"`
	GuildID         *dstypes.Snowflake   `json:"guild_id"`
	Member          *dstypes.GuildMember `json:"member"`
	Emoji           dstypes.Emoji        `json:"emoji"`
	MessageAuthorID *dstypes.Snowflake   `json:"message_author_id"`
	Burst           bool                 `json:"burst"`
	BurstColors     *[]string            `json:"burst_colors"`
	Type            dstypes.ReactionType `json:"type"`
}
