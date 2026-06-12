package ds

import (
	"cynthia/util"
	"io"
	"time"
)

type Message struct {
	ID                   Snowflake                   `json:"id"`
	ChannelID            Snowflake                   `json:"channel_id"`
	Author               User                        `json:"author"`
	Content              string                      `json:"content"`
	Timestamp            string                      `json:"timestamp"`
	EditedTimestamp      string                      `json:"edited_timestamp"`
	TTS                  bool                        `json:"tts"`
	MentionEveryone      bool                        `json:"mention_everyone"`
	Mentions             []User                      `json:"mentions"`
	MentionRoles         []string                    `json:"mention_roles"`
	MentionChannels      []ChannelMentions           `json:"mention_channels"`
	Attachments          []Attachment                `json:"attachments"`
	Embeds               []Embed                     `json:"embeds"`
	Reactions            []Reaction                  `json:"reactions"`
	Nonce                *string                     `json:"nonce"`
	Pinned               bool                        `json:"pinned"`
	WebhookID            *Snowflake                  `json:"webhook_id"`
	Type                 int                         `json:"type"`
	Activity             *MessageActivity            `json:"activity"`
	Application          *Application                `json:"application"`
	ApplicationID        *Snowflake                  `json:"application_id"`
	Flags                *int                        `json:"flags"`
	MessageReference     *MessageReference           `json:"message_reference"`
	MessageSnapshots     []MessageSnapshot           `json:"message_snapshots"`
	ReferencedMessage    *Message                    `json:"referenced_message"`
	InteractionMetadata  *MessageInteractionMetadata `json:"interaction_metadata"`
	Thread               *Channel                    `json:"thread"`
	Components           *[]MessageComponent         `json:"components"`
	StickerItems         *[]StickerItem              `json:"sticker_items"`
	Stickers             *[]Sticker                  `json:"stickers"`
	Position             *int                        `json:"position"`
	RoleSubscriptionData *RoleSubscriptionData       `json:"role_subscription_data"`
	Resolved             *ResolvedData               `json:"resolved"`
	Poll                 *Poll                       `json:"poll"`
	Call                 *MessageCall                `json:"call"`
	SharedClientTheme    *SharedClientTheme          `json:"shared_client_theme"`
}

type MessageActivity struct {
	Type    int        `json:"type"`
	PartyID *Snowflake `json:"party_id"`
}

type MessageCall struct {
	Participants   []Snowflake `json:"participants"`
	EndedTimestamp *string     `json:"ended_timestamp"`
}

type MessageFile struct {
	Name        string
	ContentType string
	Reader      io.Reader
}

type MessageInteractionMetadata struct {
	ID                           Snowflake                          `json:"id"`
	Type                         InteractionType                    `json:"type"`
	User                         User                               `json:"user"`
	AuthorizingIntegrationOwners map[ApplicationIntegrationType]any `json:"authorizing_integration_owners"`
	OriginalResponseMessageID    *Snowflake                         `json:"original_response_message_id"`
	TargetUser                   *User                              `json:"target_user"`
	TargetMessageID              *Snowflake                         `json:"target_message_id"`
}

type MessageReference struct {
	Type            *int       `json:"type"`
	MessageID       *Snowflake `json:"message_id"`
	ChannelID       *Snowflake `json:"channel_id"`
	GuildID         *Snowflake `json:"guild_id"`
	FailIfNotExists *bool      `json:"fail_if_not_exists"`
}

type MessageSnapshot struct {
	Message *Message `json:"message"`
}

type AllowedMentions string

const (
	AllowedMentionsRole     AllowedMentions = "role"
	AllowedMentionsUser     AllowedMentions = "user"
	AllowedMentionsEveryone AllowedMentions = "everyone"
)

func (a *AllowedMentions) UnmarshalJSON(b []byte) error {
	return util.UnmarshalString(b, a)
}

type Attachment struct {
	ID                 Snowflake    `json:"id"`
	Filename           string       `json:"filename"`
	Title              *string      `json:"title"`
	Description        *string      `json:"description"`
	ContentType        string       `json:"content_type"`
	Size               int          `json:"size"`
	URL                string       `json:"url"`
	ProxyURL           string       `json:"proxy_url"`
	Height             *int         `json:"height"`
	Width              *int         `json:"width"`
	Placeholder        *string      `json:"placeholder"`
	PlaceholderVersion *int         `json:"placeholder_version"`
	Ephemeral          *bool        `json:"ephemeral"`
	DurationSecs       *float64     `json:"duration_secs"`
	Waveform           *string      `json:"waveform"`
	Flags              *int         `json:"flags"`
	ClipParticipants   []User       `json:"clip_participants"`
	ClipCreatedAt      *string      `json:"clip_created_at"`
	Application        *Application `json:"application"`
}

type DefaultReaction struct {
	EmojiID   *Snowflake `json:"emoji_id"`
	EmojiName *string    `json:"emoji_name"`
}

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
	Fields      *[]EmbedField  `json:"fields"`
	Flags       *int           `json:"flags"`
}

func NewEmbed() *Embed {
	return &Embed{}
}

func (e *Embed) WithTitle(title string) *Embed {
	e.Title = &title
	return e
}

func (e *Embed) WithDescription(d string) *Embed {
	e.Description = &d
	return e
}

func (e *Embed) WithURL(url string) *Embed {
	e.URL = &url
	return e
}

func (e *Embed) WithTimestamp(t time.Time) *Embed {
	s := t.UTC().Format(time.RFC3339)
	e.Timestamp = &s
	return e
}

func (e *Embed) WithColor(color int) *Embed {
	e.Color = &color
	return e
}

func (e *Embed) WithFooter(f *EmbedFooter) *Embed {
	e.Footer = f
	return e
}

func (e *Embed) WithImage(i *EmbedImage) *Embed {
	e.Image = i
	return e
}

func (e *Embed) WithThumbnail(t *EmbedImage) *Embed {
	e.Thumbnail = t
	return e
}

func (e *Embed) WithVideo(v *EmbedVideo) *Embed {
	e.Video = v
	return e
}

func (e *Embed) WithProvider(p *EmbedProvider) *Embed {
	e.Provider = p
	return e
}

func (e *Embed) WithAuthor(a *EmbedAuthor) *Embed {
	e.Author = a
	return e
}

func (e *Embed) WithFields(f *[]EmbedField) *Embed {
	e.Fields = f
	return e
}

func (e *Embed) WithFlags(f int) *Embed {
	e.Flags = &f
	return e
}

type EmbedAuthor struct {
	Name         *string `json:"name"`
	URL          *string `json:"url"`
	IconURL      *string `json:"icon_url"`
	ProxyIconURL *string `json:"proxy_icon_url"`
}

type EmbedField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline *bool  `json:"inline"`
}

type EmbedFooter struct {
	Text         string  `json:"text"`
	IconURL      *string `json:"icon_url"`
	ProxyIconURL *string `json:"proxy_icon_url"`
}

type EmbedImage struct {
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

type EmbedProvider struct {
	Name *string `json:"name"`
	URL  *string `json:"url"`
}

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

type Emoji struct {
	ID            *Snowflake   `json:"id"`
	Name          *string      `json:"name"`
	Roles         *[]Snowflake `json:"roles"`
	User          *User        `json:"user"`
	RequireColons *bool        `json:"require_colons"`
	Managed       *bool        `json:"managed"`
	Animated      *bool        `json:"animated"`
	Available     *bool        `json:"available"`
}

type Reaction struct {
	Count        int                  `json:"count"`
	CountDetails ReactionCountDetails `json:"count_details"`
	Me           bool                 `json:"me"`
	MeBurst      bool                 `json:"me_burst"`
	Emoji        Emoji                `json:"emoji"`
	BurstColors  []string             `json:"burst_colors"`
}

type ReactionCountDetails struct {
	Burst  int `json:"burst"`
	Normal int `json:"normal"`
}

type Poll struct {
	Question         PollMedia      `json:"question"`
	Answers          []PollAnswer   `json:"poll_answers"`
	Expiry           string         `json:"expiry"`
	AllowMultiselect bool           `json:"allow_multiselect"`
	LayoutType       PollLayoutType `json:"layout_type"`
	Results          *PollResults   `json:"results"`
}

type PollAnswer struct {
	AnswerID  int       `json:"answer_id"`
	PollMedia PollMedia `json:"poll_media"`
}

type PollAnswerCount struct {
	ID      int  `json:"int"`
	Count   int  `json:"count"`
	MeVoted bool `json:"me_voted"`
}

type PollLayoutType int

const (
	PollLayoutTypeDefault PollLayoutType = 1
)

func (l *PollLayoutType) UnmarshalJSON(data []byte) error {
	return util.UnmarshalNumeric(data, l)
}

type PollMedia struct {
	Emoji *Emoji `json:"emoji"`
}

type PollResults struct {
	IsFinalized  bool              `json:"is_finalized"`
	AnswerCounts []PollAnswerCount `json:"answer_counts"`
}

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

type StickerFormatType int

const (
	StickerFormatTypePng    StickerFormatType = 1
	StickerFormatTypeAPng   StickerFormatType = 2
	StickerFormatTypeLottie StickerFormatType = 3
	StickerFormatTypeGif    StickerFormatType = 4
)

func (t *StickerFormatType) UnmarshalJSON(data []byte) error {
	return util.UnmarshalNumeric(data, t)
}

type StickerItem struct {
	ID         Snowflake         `json:"id"`
	Name       string            `json:"name"`
	FormatType StickerFormatType `json:"format_type"`
}

type StickerType int

const (
	StickerTypeStandard StickerType = 1
	StickerTypeGuild    StickerType = 2
)

func (t *StickerType) UnmarshalJSON(data []byte) error {
	return util.UnmarshalNumeric(data, t)
}
