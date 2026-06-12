package ds

import "cynthia/util"

type MessageComponent interface {
	Type() int
	ID() *int
}

type ComponentBase struct {
	Ctype int  `json:"type"`
	Id    *int `json:"id,omitempty"`
}

func (c *ComponentBase) Type() int { return c.Ctype }
func (c *ComponentBase) ID() *int  { return c.Id }

type ComponentType int

const (
	ComponentTypeActionRow         = 1
	ComponentTypeButton            = 2
	ComponentTypeStringSelect      = 3
	ComponentTypeTextInput         = 4
	ComponentTypeUserSelect        = 5
	ComponentTypeRoleSelect        = 6
	ComponentTypeMentionableSelect = 7
	ComponentTypeChannelSelect     = 8
	ComponentTypeSection           = 9
	ComponentTypeTextDisplay       = 10
	ComponentTypeThumbnail         = 11
	ComponentTypeMediaGallery      = 12
	ComponentTypeFile              = 13
	ComponentTypeSeparator         = 14
	ComponentTypeContainer         = 17
	ComponentTypeLabel             = 18
	ComponentTypeFileUpload        = 19
	ComponentTypeRadioGroup        = 21
	ComponentTypeCheckboxGroup     = 22
	ComponentTypeCheckbox          = 23
)

func (c *ComponentType) UnmarshalJSON(data []byte) error {
	return util.UnmarshalNumeric(data, c)
}

type ActionRow struct {
	ComponentBase
	Components []MessageComponent `json:"components"`
}

type ButtonStyle int

const (
	ButtonStylePrimary   = 1
	ButtonStyleSecondary = 2
	ButtonStyleSucces    = 3
	ButtonStyleDanger    = 4
	ButtonStyleLink      = 5
	ButtonStylePremium   = 6
)

func (b *ButtonStyle) UnmarshalJSON(data []byte) error {
	return util.UnmarshalNumeric(data, b)
}

type Button struct {
	ComponentBase
	Style    ButtonStyle `json:"style"`
	Label    *string     `json:"label"`
	Emoji    *Emoji      `json:"emoji"`
	CustomID *string     `json:"custom_id"`
	SkuID    *Snowflake  `json:"sku_id"`
	URL      *string     `json:"url"`
	Disabled *bool       `json:"disabled"`
}

type SelectDefaultValue struct {
	ID   Snowflake `json:"id"`
	Type string    `json:"type"`
}

type StringSelectOptions struct {
	Label       string  `json:"label"`
	Value       string  `json:"value"`
	Description *string `json:"description"`
	Emoji       *Emoji  `json:"emoji"`
	Default     *bool   `json:"default"`
}

type StringSelect struct {
	ComponentBase
	CustomID    string                `json:"custom_id"`
	Options     []StringSelectOptions `json:"options"`
	Placeholder *string               `json:"placeholder"`
	MinValues   *int                  `json:"min_values"`
	MaxValues   *int                  `json:"max_values"`
	Required    *bool                 `json:"required"`
	Disabled    *bool                 `json:"disabled"`
}

type TextInputStyle int

const (
	TextInputStyleShort     = 1
	TextInputStyleParagraph = 2
)

func (t *TextInputStyle) UnmarshalJSON(data []byte) error {
	return util.UnmarshalNumeric(data, t)
}

type TextInput struct {
	ComponentBase
	CustomID    string         `json:"custom_id"`
	Style       TextInputStyle `json:"style"`
	MinLength   *int           `json:"min_length"`
	MaxLength   *int           `json:"max_length"`
	Required    *bool          `json:"required"`
	Value       *string        `json:"value"`
	Placeholder *string        `json:"placeholder"`
}

type UserSelect struct {
	ComponentBase
	CustomID      string                `json:"custom_id"`
	Placeholder   *string               `json:"placeholder"`
	DefaultValues *[]SelectDefaultValue `json:"default_values"`
	MinValues     *int                  `json:"min_values"`
	MaxValues     *int                  `json:"max_values"`
	Required      *bool                 `json:"required"`
	Disabled      *bool                 `json:"disabled"`
}

type RoleSelect struct {
	ComponentBase
	CustomID      string                `json:"custom_id"`
	Placeholder   *string               `json:"placeholder"`
	DefaultValues *[]SelectDefaultValue `json:"default_values"`
	MinValues     *int                  `json:"min_values"`
	MaxValues     *int                  `json:"max_values"`
	Required      *bool                 `json:"required"`
	Disabled      *bool                 `json:"disabled"`
}

type MentionableSelect struct {
	ComponentBase
	CustomID      string                `json:"custom_id"`
	Placeholder   *string               `json:"placeholder"`
	DefaultValues *[]SelectDefaultValue `json:"default_values"`
	MinValues     *int                  `json:"min_values"`
	MaxValues     *int                  `json:"max_values"`
	Required      *bool                 `json:"required"`
	Disabled      *bool                 `json:"disabled"`
}

type ChannelSelect struct {
	ComponentBase
	CustomID      string                `json:"custom_id"`
	ChannelTypes  *[]ChannelType        `json:"channel_types"`
	Placeholder   *string               `json:"placeholder"`
	DefaultValues *[]SelectDefaultValue `json:"default_values"`
	MinValues     *int                  `json:"min_values"`
	MaxValues     *int                  `json:"max_values"`
	Required      *bool                 `json:"required"`
	Disabled      *bool                 `json:"disabled"`
}

type Section struct {
	ComponentBase
	Components []MessageComponent
	Accessory  MessageComponent
}

type TextDisplay struct {
	ComponentBase
	Content string `json:"content"`
}

type Thumbnail struct {
	ComponentBase
	Media       UnfurledMediaItem `json:"media"`
	Description *string           `json:"description"`
	Spoiler     *bool             `json:"spoiler"`
}

type MediaGalleryItem struct {
	Media       UnfurledMediaItem `json:"media"`
	Description *string           `json:"description"`
	Spoiler     *bool             `json:"spoiler"`
}

type MediaGallery struct {
	ComponentBase
	Items []MediaGalleryItem
}

type File struct {
	ComponentBase
	File    UnfurledMediaItem `json:"file"`
	Spoiler *bool             `json:"spoiler"`
	Name    *string           `json:"name"`
	Size    *int              `json:"size"`
}

type Separator struct {
	ComponentBase
	Divider *bool `json:"divider"`
	Spacing *int  `json:"spacing"`
}

type Container struct {
	ComponentBase
	Components  []MessageComponent `json:"components"`
	AccentColor *int               `json:"accent_color"`
	Spoiler     *bool              `json:"spoiler"`
}

type Label struct {
	ComponentBase
	Label       string           `json:"label"`
	Description *string          `json:"description"`
	Component   MessageComponent `json:"component"`
}

type FileUpload struct {
	ComponentBase
	CustomID  string `json:"custom_id"`
	MinValues *int   `json:"min_values"`
	MaxValues *int   `json:"max_values"`
	Required  *bool  `json:"required"`
}

type RadioGroupOption struct {
	Value       string  `json:"value"`
	Label       string  `json:"label"`
	Description *string `json:"description"`
	Default     *bool   `json:"default"`
}

type RadioGroup struct {
	ComponentBase
	CustomID string             `json:"custom_id"`
	Options  []RadioGroupOption `json:"options"`
	Required *bool              `json:"required"`
}

type CheckboxGroupOption struct {
	Value       string  `json:"value"`
	Label       string  `json:"label"`
	Description *string `json:"description"`
	Default     *bool   `json:"default"`
}

type CheckboxGroup struct {
	ComponentBase
	CustomID  string                `json:"custom_id"`
	Options   []CheckboxGroupOption `json:"options"`
	MinValues *int                  `json:"min_values"`
	MaxValues *int                  `json:"max_values"`
	Required  *bool                 `json:"required"`
}

type Checkbox struct {
	ComponentBase
	CustomID string `json:"custom_id"`
	Default  *bool  `json:"default"`
}

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
