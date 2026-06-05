//go:generate go run ../../cmd/gen_handlers/main.go

package gateway

import (
	"cynthia/ds/dstypes"
	"cynthia/ds/payloads"
	"encoding/json"
	"log/slog"
	"time"

	"github.com/gorilla/websocket"
)

const GatewayURL = "wss://gateway.discord.gg/"

func StartHeartbeat(conn *websocket.Conn, intervalMs int, seq *int, onSend func()) {
	send := func() {
		onSend()

		var seqVal any = nil

		if seq != nil {
			seqVal = *seq
		}

		p, _ := json.Marshal(map[string]any{"op": payloads.OpHeartbeat, "d": seqVal})

		if err := conn.WriteMessage(websocket.TextMessage, p); err != nil {
			slog.Error("Failed to send heartbeat payload", "err", err)
			return
		}

		slog.Debug("Heartbeat sent.")
	}

	// Send first to measure latency
	send()

	ticker := time.NewTicker(time.Duration(intervalMs) * time.Millisecond)

	defer ticker.Stop()

	for range ticker.C {
		send()
		slog.Debug("Heartbeat sent.")
	}
}

func Identify(conn *websocket.Conn, token string, intents dstypes.Intents) {
	p, _ := json.Marshal(map[string]any{
		"op": payloads.OpIdentify,
		"d": payloads.Identify{
			Token:   "Bot " + token,
			Intents: intents,
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
