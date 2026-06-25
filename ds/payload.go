package ds

import (
	"fmt"
)

type Op int

const (
	OpDispatch         Op = 0
	OpHeartbeat        Op = 1
	OpIdentify         Op = 2
	OpPresenceUpdate   Op = 3
	OpVoiceStateUpdate Op = 4
	OpResume           Op = 6
	OpReconnect        Op = 7
	OpRequestMembers   Op = 8
	OpInvalidSession   Op = 9
	OpHello            Op = 10
	OpHeartbeatACK     Op = 11
)

func (op Op) String() string {
	switch op {
	case OpDispatch:
		return "Dispatch"
	case OpHeartbeat:
		return "Heartbeat"
	case OpIdentify:
		return "Identify"
	case OpPresenceUpdate:
		return "PresenceUpdate"
	case OpVoiceStateUpdate:
		return "VoiceStateUpdate"
	case OpResume:
		return "Resume"
	case OpReconnect:
		return "Reconnect"
	case OpRequestMembers:
		return "RequestMembers"
	case OpInvalidSession:
		return "InvalidSession"
	case OpHello:
		return "Hello"
	case OpHeartbeatACK:
		return "HeartbeatACK"
	default:
		return fmt.Sprintf("Unknown(%d)", int(op))
	}
}

func (o *Op) UnmarshalJSON(b []byte) error {
	return unmarshalNumeric(b, o)
}

type Payload[T any] struct {
	Op Op      `json:"op"`
	D  T       `json:"d"`
	S  *int    `json:"s"`
	T  *string `json:"t"`
}
