package ds

type ThreadMember struct {
	ID            *Snowflake   `json:"id"`
	UserID        *Snowflake   `json:"user_id"`
	JoinTimestamp string       `json:"join_timestamp"`
	Flags         int          `json:"flags"`
	Member        *GuildMember `json:"member"`
}
