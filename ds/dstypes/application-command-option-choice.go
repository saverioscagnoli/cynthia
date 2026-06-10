package dstypes

type ApplicationCommandOptionChoice struct {
	Name              string             `json:"name"`
	NameLocalizations *map[string]string `json:"name_localizations"`
	Value             any                `json:"value"`
}
