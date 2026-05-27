package main

import (
	"cynthia/ds"
	"encoding/json"
	"log/slog"
	"os"

	"github.com/gorilla/websocket"
)

func main() {
	conn, _, err := websocket.DefaultDialer.Dial(ds.GatewayURL, nil)

	if err != nil {
		slog.Error("Failed to connect!", "err", err)
		os.Exit(1)
	}

	defer conn.Close()

	slog.Info("Connected to the gateway")

	for {
		_, msg, err := conn.ReadMessage()

		if err != nil {
			slog.Error("Read error", "err", err)
			return
		}

		var payload ds.Payload

		if err := json.Unmarshal(msg, &payload); err != nil {
			slog.Error("Json decoding error", "err", err)
			continue
		}

		switch payload.Op {
		}
	}

}
