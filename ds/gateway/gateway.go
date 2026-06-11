//go:generate go run ../../cmd/gen_handlers/main.go

package gateway

import (
	"cynthia/ds/dstypes"
	"cynthia/ds/payloads"
	"log/slog"
	"time"

	"github.com/gorilla/websocket"
)

const GatewayURL = "wss://gateway.discord.gg/"

func (c *Client) StartHeartbeat(conn *websocket.Conn, intervalMs int, seq *int, onSend func()) {
	send := func() {
		onSend()

		var seqVal any = nil

		if seq != nil {
			seqVal = *seq
		}

		if err := c.writeJSON(conn, map[string]any{"op": payloads.OpHeartbeat, "d": seqVal}); err != nil {
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

func (c *Client) Identify(conn *websocket.Conn, token string, intents dstypes.Intents) {
	if err := c.writeJSON(conn, map[string]any{
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
	}); err != nil {
		slog.Error("Failed to identiy", "err", err)
		return
	}

	slog.Info("Identification sent.")
}
