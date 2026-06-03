package ds

type StickerType = int

const (
	StickerTypeStandard StickerType = 1
	StickerTypeGuild    StickerType = 2
)

type StickerFormatType = int

const (
	StickerFormatTypePng    StickerFormatType = 1
	StickerFormatTypeAPng   StickerFormatType = 2
	StickerFormatTypeLottie StickerFormatType = 3
	StickerFormatTypeGif    StickerFormatType = 4
)

type Sticker struct {
	ID          Snowflake         `json:"id"`
	PackID      *Snowflake        `json:"pack_id"`
	Name        string            `json:"name"`
	Description *string           `json:"description"`
	Tags        string            `json:"tags"`
	Type        StickerType       `json:"type"`
	FormatType  StickerFormatType `json:"format_type"`
	Available   *bool             `json:"available"`
	GuildID     *Snowflake        `json:"guild_id"`
	User        *User             `json:"user"`
	SortValue   *int              `json:"sort_value"`
}
