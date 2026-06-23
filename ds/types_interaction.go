package ds

import (
	"cynthia/util"
	"encoding/json"
	"fmt"
)

type Interaction struct {
	ID                           Snowflake                          `json:"id"`
	ApplicationID                Snowflake                          `json:"application_id"`
	Type                         InteractionType                    `json:"type"`
	Data                         *json.RawMessage                   `json:"data"`
	Guild                        *Guild                             `json:"guild"`
	GuildID                      *Snowflake                         `json:"guild_id"`
	Channel                      *Channel                           `json:"channel"`
	ChannelID                    *Snowflake                         `json:"channel_id"`
	Member                       *GuildMember                       `json:"member"`
	User                         *User                              `json:"user"`
	Token                        string                             `json:"token"`
	Version                      int                                `json:"version"`
	Message                      *Message                           `json:"message"`
	AppPermissions               string                             `json:"app_permissions"`
	Locale                       *string                            `json:"locale"`
	GuildLocale                  *string                            `json:"guild_locale"`
	Entitlements                 []Entitlement                      `json:"entitlements"`
	AuthorizingIntegrationOwners map[ApplicationIntegrationType]any `json:"authorizing_integration_owners"`
	Contexts                     *InteractionContextType            `json:"contexts"`
	AttachmentSizeLimit          int                                `json:"attachment_size_limit"`
}

func (i *Interaction) ApplicationCommandData() (*ApplicationCommandData, error) {
	if i.Data == nil {
		return nil, fmt.Errorf("interaction data is nil")
	}

	if i.Type != InteractionTypeApplicationCommand && i.Type != InteractionTypeApplicationCommandAutocomplete {
		return nil, fmt.Errorf("interaction type is %v, not ApplicationCommand nor AppplicationCommandAutocomplete", i.Type)
	}

	var d ApplicationCommandData

	if err := json.Unmarshal(*i.Data, &d); err != nil {
		return nil, err
	}

	return &d, nil
}

func (i *Interaction) MessageComponentData() (*MessageComponentData, error) {
	if i.Data == nil {
		return nil, fmt.Errorf("interaction data is nil")
	}

	if i.Type != InteractionTypeMessageComponent {
		return nil, fmt.Errorf("interaction type is %v, not MessageComponent", i.Type)
	}

	var d MessageComponentData

	if err := json.Unmarshal(*i.Data, &d); err != nil {
		return nil, err
	}

	return &d, nil
}

func (i *Interaction) ModalSubmitData() (*ModalSubmitData, error) {
	if i.Data == nil {
		return nil, fmt.Errorf("interaction data is nil")
	}

	if i.Type != InteractionTypeModalSubmit {
		return nil, fmt.Errorf("interaction type is %v, not ModalSubmit", i.Type)
	}

	var d ModalSubmitData

	if err := json.Unmarshal(*i.Data, &d); err != nil {
		return nil, err
	}

	return &d, nil
}

type InteractionCallbackData struct {
	TTS             bool               `json:"tts"`
	Content         string             `json:"content"`
	Embeds          []*Embed           `json:"embeds"`
	AllowedMentions *AllowedMentions   `json:"allowed_mentions"`
	Flags           *int               `json:"flags"`
	Components      []MessageComponent `json:"components"`
	Attachments     []*Attachment      `json:"attachments"`
	Poll            *Poll              `json:"poll"`
}

type InteractionCallbackType int

const (
	InteractionCallbackTypePong                                 InteractionCallbackType = 1
	InteractionCallbackTypeChannelMessageWithSource             InteractionCallbackType = 4
	InteractionCallbackTypeDeferredChannelMessageWithSource     InteractionCallbackType = 5
	InteractionCallbackTypeDeferredUpdateMessage                InteractionCallbackType = 6
	InteractionCallbackTypeUpdateMessage                        InteractionCallbackType = 7
	InteractionCallbackTypeApplicationCommandAutoCompleteResult InteractionCallbackType = 8
	InteractionCallbackTypeModal                                InteractionCallbackType = 9
	InteractionCallbackTypePremiumRequired                      InteractionCallbackType = 10
	InteractionCallbackTypeLaunchActivity                       InteractionCallbackType = 11
)

func (t *InteractionCallbackType) UnmarshalJSON(data []byte) error {
	return util.UnmarshalNumeric(data, t)
}

type InteractionContextType int

const (
	InteractionContextTypeGuild          InteractionContextType = 0
	InteractionContextTypeBotDM          InteractionContextType = 1
	InteractionContextTypePrivateChannel InteractionContextType = 2
)

func (e *InteractionContextType) UnmarshalJSON(b []byte) error {
	return util.UnmarshalNumeric(b, e)
}

type InteractionResponse struct {
	Type InteractionCallbackType  `json:"type"`
	Data *InteractionCallbackData `json:"data"`
}

type InteractionType int

const (
	InteractionTypePing                           InteractionType = 1
	InteractionTypeApplicationCommand             InteractionType = 2
	InteractionTypeMessageComponent               InteractionType = 3
	InteractionTypeApplicationCommandAutocomplete InteractionType = 4
	InteractionTypeModalSubmit                    InteractionType = 5
)

func (i *InteractionType) UnmarshalJSON(data []byte) error {
	return util.UnmarshalNumeric(data, i)
}

type ApplicationCommandData struct {
	ID       Snowflake                                  `json:"id"`
	Name     string                                     `json:"name"`
	Type     ApplicationCommandType                     `json:"type"`
	Resolved *ResolvedData                              `json:"resolved"`
	Options  *[]ApplicationCommandInteractionDataOption `json:"options"`
	GuildID  *Snowflake                                 `json:"guild_id"`
	TargetID *Snowflake                                 `json:"target_id"`
}

type ApplicationCommandOption struct {
	Type                     ApplicationCommandOptionType      `json:"type"`
	Name                     string                            `json:"name"`
	NameLocalizations        *map[string]string                `json:"name_localizations"`
	Description              string                            `json:"description"`
	DescriptionLocalizations *map[string]string                `json:"description_localizations"`
	Required                 *bool                             `json:"required"`
	Choices                  *[]ApplicationCommandOptionChoice `json:"choices"`
	Options                  *[]ApplicationCommandOption       `json:"options"`
	ChannelType              *[]ChannelType                    `json:"channel_types"`
	MinValue                 *int64                            `json:"min_value"`
	MaxValue                 *int64                            `json:"max_value"`
	MinLength                *int64                            `json:"min_length"`
	MaxLength                *int64                            `json:"max_length"`
	Autocomplete             *bool                             `json:"autocomplete"`
}

type ApplicationCommandOptionChoice struct {
	Name              string             `json:"name"`
	NameLocalizations *map[string]string `json:"name_localizations"`
	Value             any                `json:"value"`
}

type ApplicationCommandOptionType int

const (
	ApplicationCommandOptionTypeSubCommand      ApplicationCommandOptionType = 1
	ApplicationCommandOptionTypeSubCommandGroup ApplicationCommandOptionType = 2
	ApplicationCommandOptionTypeString          ApplicationCommandOptionType = 3
	ApplicationCommandOptionTypeInteger         ApplicationCommandOptionType = 4
	ApplicationCommandOptionTypeBoolean         ApplicationCommandOptionType = 5
	ApplicationCommandOptionTypeUser            ApplicationCommandOptionType = 6
	ApplicationCommandOptionTypeChannel         ApplicationCommandOptionType = 7
	ApplicationCommandOptionTypeRole            ApplicationCommandOptionType = 8
	ApplicationCommandOptionTypeMentionable     ApplicationCommandOptionType = 9
	ApplicationCommandOptionTypeNumber          ApplicationCommandOptionType = 10
	ApplicationCommandOptionTypeAttachment      ApplicationCommandOptionType = 11
)

func (a *ApplicationCommandOptionType) UnmarshalJSON(b []byte) error {
	return util.UnmarshalNumeric(b, a)
}

type ApplicationCommandType int

const (
	ApplicationCommandTypeChatInput         ApplicationCommandType = 1
	ApplicationCommandTypeUser              ApplicationCommandType = 2
	ApplicationCommandTypeMessage           ApplicationCommandType = 3
	ApplicationCommandTypePrimaryEntryPoint ApplicationCommandType = 4
)

func (a *ApplicationCommandType) UnmarshalJSON(b []byte) error {
	return util.UnmarshalNumeric(b, a)
}

type ApplicationCommandInteractionDataOption struct {
	Name    string                       `json:"name"`
	Type    ApplicationCommandOptionType `json:"type"`
	Value   *any
	Options *[]ApplicationCommandInteractionDataOption
	Focused *bool
}

type MessageComponentData struct {
	CustomID      string        `json:"custom_id"`
	ComponentType ComponentType `json:"component_type"`
	Values        *[]string     `json:"values"`
	Resolved      *ResolvedData `json:"resolved"`
}

type ModalSubmitData struct {
	CustomID   string             `json:"custom_id"`
	Components []MessageComponent `json:"component"`
	Resolved   *ResolvedData      `json:"resolved"`
}

type ResolvedData struct {
	Users      *map[Snowflake]User
	Members    *map[Snowflake]GuildMember
	Roles      *map[Snowflake]Role
	Channels   *map[Snowflake]Channel
	Messages   *map[Snowflake]Message
	Attachment *map[Snowflake]Attachment
}
