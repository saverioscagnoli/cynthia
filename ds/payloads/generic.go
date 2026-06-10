package payloads

import (
	"cynthia/ds/dsevents"
	"encoding/json"
)

type GenericPayload struct {
	Op Op                  `json:"op"`
	D  json.RawMessage     `json:"d"`
	S  *int                `json:"s"`
	T  *dsevents.EventName `json:"t"`
}

type Payload interface {
	EventName() dsevents.EventName
}
