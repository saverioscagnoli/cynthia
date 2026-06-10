package dstypes

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
