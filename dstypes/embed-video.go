package dstypes

type EmbedVideo struct {
	URL                string  `json:"url"`
	ProxyURL           *string `json:"proxy_url"`
	Height             *int    `json:"height"`
	Width              *int    `json:"width"`
	ContentType        *string `json:"content_type"`
	Placeholder        *string `json:"placeholder"`
	PlaceholderVersion *int    `json:"placeholder_version"`
	Description        *string `json:"description"`
	Flags              *int    `json:"flags"`
}
