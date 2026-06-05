package payloads

import (
	"cynthia/ds/events"
	"encoding/json"
)

type GenericPayload struct {
	Op Op                `json:"op"`
	D  json.RawMessage   `json:"d"`
	S  *int              `json:"s"`
	T  *events.EventName `json:"t"`
}

type Payload interface {
	EventName() events.EventName
}
