package ds

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"time"

	"github.com/gorilla/websocket"
)

const GatewayURL = "wss://gateway.discord.gg/"

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

type Payload struct {
	Op Op              `json:"op"`
	D  json.RawMessage `json:"d"`
	S  *int            `json:"s"`
	T  *string         `json:"t"`
}

type Hello struct {
	HeartbeatInterval int `json:"heartbeat_interval"`
}

type Identify struct {
	Token      string            `json:"token"`
	Intents    int               `json:"intents"`
	Properties map[string]string `json:"properties"`
}

func startHeartbeat(conn *websocket.Conn, intervalMs int, seq **int) {
	ticker := time.NewTicker(time.Duration(intervalMs) * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		var seqVal any = nil

		if *seq != nil {
			seqVal = **seq
		}

		p, _ := json.Marshal(map[string]any{"op": OpHeartbeat, "d": seqVal})

		if err := conn.WriteMessage(websocket.TextMessage, p); err != nil {
			slog.Error("Failed to send heartbeat payload", "err", err)
			return
		}

		slog.Debug("Heartbeat sent.")
	}
}

func indentify(conn *websocket.Conn, token string) {
	p, _ := json.Marshal(map[string]any{
		"op": OpIdentify,
		"d": Identify{
			Token:   "Bot " + token,
			Intents: 1 << 9,
			Properties: map[string]string{
				"os":      "linux",
				"browser": "cynthia",
				"device":  "cynthia",
			},
		},
	})

	conn.WriteMessage(websocket.TextMessage, p)
	slog.Info("Identification sent.")
}
