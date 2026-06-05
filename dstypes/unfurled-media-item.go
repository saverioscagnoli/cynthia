package dstypes

type UnfurledMediaItem struct {
	URL                string     `json:"url"`
	ProxyURL           *string    `json:"proxy_url"`
	Height             *int       `json:"height"`
	Width              *int       `json:"width"`
	Placeholder        *string    `json:"placeholder"`
	PlaceholderVersion *int       `json:"placeholder_version"`
	ContentType        *string    `json:"content_type"`
	Flags              *int       `json:"flags"`
	AttachmentID       *Snowflake `json:"attachment_id"`
}
