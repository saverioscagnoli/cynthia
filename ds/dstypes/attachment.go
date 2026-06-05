package dstypes

type Attachment struct {
	ID                 Snowflake `json:"id"`
	Filename           string    `json:"filename"`
	Title              *string   `json:"title"`
	Description        *string   `json:"description"`
	ContentType        string    `json:"content_type"`
	Size               int       `json:"size"`
	URL                string    `json:"url"`
	ProxyURL           string    `json:"proxy_url"`
	Height             *int      `json:"height"`
	Width              *int      `json:"width"`
	Placeholder        *string   `json:"placeholder"`
	PlaceholderVersion *int      `json:"placeholder_version"`
	Ephemeral          *bool     `json:"ephemeral"`
	DurationSecs       *float64  `json:"duration_secs"`
	Waveform           *string   `json:"waveform"`
	Flags              *int      `json:"flags"`
	ClipParticipants   []User    `json:"clip_participants"`
	ClipCreatedAt      *string   `json:"clip_created_at"`
	// TODO: Application        *Application `json:"application"`
}
