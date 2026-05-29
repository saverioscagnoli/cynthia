package ds

type MessageInteractionMetadata struct {
	ID                           Snowflake          `json:"id"`
	Type                         InteractionType    `json:"type"`
	User                         User               `json:"user"`
	AuthorizingIntegrationOwners map[Snowflake]bool `json:"authorizing_integration_owners"`
	OriginalResponseMessageID    *Snowflake         `json:"original_response_message_id"`
	TargetUser                   *User              `json:"target_user"`
	TargetMessageID              *Snowflake         `json:"target_message_id"`
}
