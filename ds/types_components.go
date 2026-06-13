package ds

import (
	"cynthia/util"
	"encoding/json"
)

type MessageComponent interface {
	Type() ComponentType
	ID() *int
}

type ComponentBase struct {
	Ctype ComponentType `json:"type"`
	Id    *int          `json:"id,omitempty"`
}

func (c *ComponentBase) Type() ComponentType { return c.Ctype }
func (c *ComponentBase) ID() *int            { return c.Id }

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

func UnmarshalMessageComponent(data []byte) (MessageComponent, error) {
	var base ComponentBase
	if err := json.Unmarshal(data, &base); err != nil {
		return nil, err
	}

	switch base.Ctype {
	case ComponentTypeActionRow:
		var c ActionRow
		if err := json.Unmarshal(data, &c); err != nil {
			return nil, err
		}
		return &c, nil
	case ComponentTypeButton:
		var c Button
		if err := json.Unmarshal(data, &c); err != nil {
			return nil, err
		}
		return &c, nil
	case ComponentTypeStringSelect:
		var c StringSelect
		if err := json.Unmarshal(data, &c); err != nil {
			return nil, err
		}
		return &c, nil
	case ComponentTypeTextInput:
		var c TextInput
		if err := json.Unmarshal(data, &c); err != nil {
			return nil, err
		}
		return &c, nil
	case ComponentTypeUserSelect:
		var c UserSelect
		if err := json.Unmarshal(data, &c); err != nil {
			return nil, err
		}
		return &c, nil
	case ComponentTypeRoleSelect:
		var c RoleSelect
		if err := json.Unmarshal(data, &c); err != nil {
			return nil, err
		}
		return &c, nil
	case ComponentTypeMentionableSelect:
		var c MentionableSelect
		if err := json.Unmarshal(data, &c); err != nil {
			return nil, err
		}
		return &c, nil
	case ComponentTypeChannelSelect:
		var c ChannelSelect
		if err := json.Unmarshal(data, &c); err != nil {
			return nil, err
		}
		return &c, nil
	case ComponentTypeSection:
		var c Section
		if err := json.Unmarshal(data, &c); err != nil {
			return nil, err
		}
		return &c, nil
	case ComponentTypeTextDisplay:
		var c TextDisplay
		if err := json.Unmarshal(data, &c); err != nil {
			return nil, err
		}
		return &c, nil
	case ComponentTypeThumbnail:
		var c Thumbnail
		if err := json.Unmarshal(data, &c); err != nil {
			return nil, err
		}
		return &c, nil
	case ComponentTypeMediaGallery:
		var c MediaGallery
		if err := json.Unmarshal(data, &c); err != nil {
			return nil, err
		}
		return &c, nil
	case ComponentTypeFile:
		var c File
		if err := json.Unmarshal(data, &c); err != nil {
			return nil, err
		}
		return &c, nil
	case ComponentTypeSeparator:
		var c Separator
		if err := json.Unmarshal(data, &c); err != nil {
			return nil, err
		}
		return &c, nil
	case ComponentTypeContainer:
		var c Container
		if err := json.Unmarshal(data, &c); err != nil {
			return nil, err
		}
		return &c, nil
	case ComponentTypeLabel:
		var c Label
		if err := json.Unmarshal(data, &c); err != nil {
			return nil, err
		}
		return &c, nil
	case ComponentTypeFileUpload:
		var c FileUpload
		if err := json.Unmarshal(data, &c); err != nil {
			return nil, err
		}
		return &c, nil
	case ComponentTypeRadioGroup:
		var c RadioGroup
		if err := json.Unmarshal(data, &c); err != nil {
			return nil, err
		}
		return &c, nil
	case ComponentTypeCheckboxGroup:
		var c CheckboxGroup
		if err := json.Unmarshal(data, &c); err != nil {
			return nil, err
		}
		return &c, nil
	case ComponentTypeCheckbox:
		var c Checkbox
		if err := json.Unmarshal(data, &c); err != nil {
			return nil, err
		}
		return &c, nil
	default:
		return &base, nil
	}
}

func unmarshalComponents(raws []json.RawMessage) ([]MessageComponent, error) {
	components := make([]MessageComponent, 0, len(raws))
	for _, r := range raws {
		c, err := UnmarshalMessageComponent(r)
		if err != nil {
			return nil, err
		}
		components = append(components, c)
	}
	return components, nil
}

type ActionRow struct {
	ComponentBase
	Components []MessageComponent `json:"components"`
}

func NewActionRow() *ActionRow {
	return &ActionRow{ComponentBase: ComponentBase{Ctype: ComponentTypeActionRow}}
}

func (a *ActionRow) Add(c MessageComponent) {
	a.Components = append(a.Components, c)
}

func (a *ActionRow) WithComponent(c MessageComponent) *ActionRow {
	a.Add(c)
	return a
}

func (a *ActionRow) UnmarshalJSON(data []byte) error {
	var raw struct {
		ComponentBase
		Components []json.RawMessage `json:"components"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	a.ComponentBase = raw.ComponentBase
	var err error
	a.Components, err = unmarshalComponents(raw.Components)
	return err
}

func (a *ActionRow) MarshalJSON() ([]byte, error) {
	components := make([]json.RawMessage, len(a.Components))
	for i, c := range a.Components {
		data, err := json.Marshal(c)
		if err != nil {
			return nil, err
		}
		components[i] = data
	}
	return json.Marshal(&struct {
		Type       ComponentType     `json:"type"`
		ID         *int              `json:"id,omitempty"`
		Components []json.RawMessage `json:"components"`
	}{
		Type:       a.Ctype,
		ID:         a.Id,
		Components: components,
	})
}

type Section struct {
	ComponentBase
	Components []MessageComponent `json:"components"`
	Accessory  MessageComponent   `json:"accessory"`
}

func (s *Section) UnmarshalJSON(data []byte) error {
	var raw struct {
		ComponentBase
		Components []json.RawMessage `json:"components"`
		Accessory  json.RawMessage   `json:"accessory"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	s.ComponentBase = raw.ComponentBase
	var err error
	s.Components, err = unmarshalComponents(raw.Components)
	if err != nil {
		return err
	}
	if raw.Accessory != nil {
		s.Accessory, err = UnmarshalMessageComponent(raw.Accessory)
		if err != nil {
			return err
		}
	}
	return nil
}

type Container struct {
	ComponentBase
	Components  []MessageComponent `json:"components"`
	AccentColor *int               `json:"accent_color"`
	Spoiler     *bool              `json:"spoiler"`
}

func (c *Container) UnmarshalJSON(data []byte) error {
	var raw struct {
		ComponentBase
		Components  []json.RawMessage `json:"components"`
		AccentColor *int              `json:"accent_color"`
		Spoiler     *bool             `json:"spoiler"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	c.ComponentBase = raw.ComponentBase
	c.AccentColor = raw.AccentColor
	c.Spoiler = raw.Spoiler
	var err error
	c.Components, err = unmarshalComponents(raw.Components)
	return err
}

type Label struct {
	ComponentBase
	LabelText   string           `json:"label"`
	Description *string          `json:"description"`
	Component   MessageComponent `json:"component"`
}

func (l *Label) UnmarshalJSON(data []byte) error {
	var raw struct {
		ComponentBase
		LabelText   string          `json:"label"`
		Description *string         `json:"description"`
		Component   json.RawMessage `json:"component"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	l.ComponentBase = raw.ComponentBase
	l.LabelText = raw.LabelText
	l.Description = raw.Description
	if raw.Component != nil {
		var err error
		l.Component, err = UnmarshalMessageComponent(raw.Component)
		if err != nil {
			return err
		}
	}
	return nil
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

func NewButton() *Button {
	return &Button{ComponentBase: ComponentBase{Ctype: ComponentTypeButton}}
}

func (b *Button) WithLabel(label string) *Button {
	b.Label = &label
	return b
}

func (b *Button) WithStyle(style ButtonStyle) *Button {
	b.Style = style
	return b
}

func (b *Button) WithEmoji(emoji *Emoji) *Button {
	b.Emoji = emoji
	return b
}

func (b *Button) WithCustomID(id string) *Button {
	b.CustomID = &id
	return b
}

func (b *Button) WithURL(url string) *Button {
	b.URL = &url
	return b
}

func (b *Button) WithDisabled(d bool) *Button {
	b.Disabled = &d
	return b
}

func (b *Button) Disable() {
	b.Disabled = util.Ptr(true)
}

func (b *Button) Enable() {
	b.Disabled = util.Ptr(false)
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
	Items []MediaGalleryItem `json:"items"`
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
