package dstypes

type InteractionType string

const (
	InteractionTypePing                           InteractionType = "PING"
	InteractionTypeApplicationCommand             InteractionType = "APPLICATION_COMMAND"
	InteractionTypeMessageComponent               InteractionType = "MESSAGE_COMPONENT"
	InteractionTypeApplicationCommandAutocomplete InteractionType = "APPLICATION_COMMAND_AUTOCOMPLETE"
	InteractionTypeModalSubmit                    InteractionType = "MODAL_SUBMIT"
)
