package payloads

import "cynthia/dstypes"

type Identify struct {
	Token      string            `json:"token"`
	Intents    dstypes.Intents   `json:"intents"`
	Properties map[string]string `json:"properties"`
}
