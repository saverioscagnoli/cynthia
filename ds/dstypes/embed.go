package dstypes

type Embed struct {
	Title       *string        `json:"title"`
	Type        *string        `json:"type"`
	Description *string        `json:"description"`
	URL         *string        `json:"url"`
	Timestamp   *string        `json:"timestamp"`
	Color       *int           `json:"color"`
	Footer      *EmbedFooter   `json:"footer"`
	Image       *EmbedImage    `json:"image"`
	Thumbnail   *EmbedImage    `json:"thumbnail"`
	Video       *EmbedVideo    `json:"video"`
	Provider    *EmbedProvider `json:"provider"`
	Author      *EmbedAuthor   `json:"author"`
	Fields      []EmbedField   `json:"fields"`
	Flags       *int           `json:"flags"`
}

type EmbedBuilder struct {
	embed Embed
}

func strPtr(s string) *string { return &s }
func intPtr(i int) *int       { return &i }

func NewEmbed() *EmbedBuilder {
	return &EmbedBuilder{
		embed: Embed{},
	}
}

func (b *EmbedBuilder) Title(t string) *EmbedBuilder {
	b.embed.Title = strPtr(t)
	return b
}

func (b *EmbedBuilder) Description(d string) *EmbedBuilder {
	b.embed.Description = strPtr(d)
	return b
}

func (b *EmbedBuilder) URL(u string) *EmbedBuilder {
	b.embed.URL = strPtr(u)
	return b
}

func (b *EmbedBuilder) Color(c int) *EmbedBuilder {
	b.embed.Color = intPtr(c)
	return b
}

func (b *EmbedBuilder) Image(url string) *EmbedBuilder {
	b.embed.Image = &EmbedImage{URL: url}
	return b
}

func (b *EmbedBuilder) Thumbnail(url string) *EmbedBuilder {
	b.embed.Thumbnail = &EmbedImage{URL: url}
	return b
}

func (b *EmbedBuilder) Video(url string) *EmbedBuilder {
	b.embed.Video = &EmbedVideo{URL: url}
	return b
}

func (b *EmbedBuilder) Build() Embed {
	return b.embed
}
