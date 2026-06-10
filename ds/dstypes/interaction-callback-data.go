package dstypes

type InteractionCallbackData struct {
	TTS             bool                `json:"tts"`
	Content         string              `json:"content"`
	Embeds          []*Embed            `json:"embeds"`
	AllowedMentions *AllowedMentions    `json:"allowed_mentions"`
	Flags           *int                `json:"flags"`
	Components      []*MessageComponent `json:"components"`
	Attachments     []*Attachment       `json:"attachments"`
	Poll            *Poll               `json:"poll"`
}
