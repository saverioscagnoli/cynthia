package ds

// Payloads defined specifically for `Send Events`, which are events
// that the user application sends to the Discord gateway, on the contrary,
// `Dispatch Events` are the events that the Discord gateway sends to the user application.
//
// Only this file is dedicated to send events, since they are not that many.
// Any other file is dedicated to dispatch events.

type HelloPayload struct {
	HeartbeatInterval int `json:"heartbeat_interval"`
}

type IdentifyProperties struct {
	Os      string `json:"os"`
	Browser string `json:"browser"`
	Device  string `json:"device"`
}

type IdentifyPayload struct {
	Token      string             `json:"token"`
	Intents    Intent             `json:"intents"`
	Properties IdentifyProperties `json:"properties"`
}
