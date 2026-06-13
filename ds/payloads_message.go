package ds

import "encoding/json"

type MessageCreate struct {
	Message
	GuildID     *Snowflake          `json:"guild_id"`
	Member      *GuildMember        `json:"member"`
	Mentions    *[]User             `json:"mentions"`
	ChannelType *ChannelType        `json:"channel_type"`
	Components  *[]MessageComponent `json:"components"`
}

func (mc *MessageCreate) UnmarshalJSON(data []byte) error {
	type Alias MessageCreate
	var raw struct {
		Alias
		Components []json.RawMessage `json:"components"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	*mc = MessageCreate(raw.Alias)
	if raw.Components != nil {
		var err error
		mc.Components, err = unmarshalComponentsPtr(raw.Components)
		if err != nil {
			return err
		}
	}
	return nil
}
