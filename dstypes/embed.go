package dstypes

type Embed struct {
	Title       *string        `json:"title"`
	Type        *string        `json:"type"`
	Description *string        `json:"description"`
	URL         *string        `json:"url"`
	Timestamp   *string        `json:"timestamp"`
	Color       *int           `json:"color"`
	Footer      *EmbedFooter   `json:"footer"`
	Image       *EmbedImage    `json:"embed_image"`
	Thumbnail   *EmbedImage    `json:"thumbnail"`
	Video       *EmbedVideo    `json:"video"`
	Provider    *EmbedProvider `json:"provider"`
	Author      *EmbedAuthor   `json:"author"`
	Fields      []EmbedField   `json:"fields"`
	Flags       *int           `json:"flags"`
}
