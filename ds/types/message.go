package ds

type Message struct {
	ID              string            `json:"id"`
	ChannelID       string            `json:"channel_id"`
	Author          User              `json:"author"`
	Content         string            `json:"content"`
	Timestamp       string            `json:"timestamp"`
	EditedTimestamp string            `json:"edited_timestamp"`
	TTS             bool              `json:"tts"`
	MentionEveryone bool              `json:"mention_everyone"`
	Mentions        []User            `json:"mentions"`
	MentionRoles    []string          `json:"mention_roles"`
	MentionChannels []ChannelMentions `json:"mention_channels"`
	Attachments     []Attachment      `json:"attachments"`
	Embeds          []Embed           `json:"embeds"`
	Reactions       []Reaction        `json:"reactions"`
	Nonce           *string           `json:"nonce"`
	Pinned          bool              `json:"pinned"`
	WebhookID       *Snowflake        `json:"webhook_id"`
	Type            int               `json:"type"`
	Activity        *MessageActivity  `json:"activity"`
	// TODO: Application     *Application      `json:"application"`
	ApplicationID       *Snowflake                  `json:"application_id"`
	Flags               *int                        `json:"flags"`
	MessageReference    *MessageReference           `json:"message_reference"`
	MessageSnapshots    []MessageSnapshot           `json:"message_snapshots"`
	ReferencedMessage   *Message                    `json:"referenced_message"`
	InteractionMetadata *MessageInteractionMetadata `json:"interaction_metadata"`
}
