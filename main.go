package main

import (
	"cynthia/ds"
	"cynthia/ds/events"
	"cynthia/ds/payloads"
	"encoding/json"
	"flag"
	"log/slog"
	"os"

	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
)

func main() {
	logLevel := flag.String("level", "info", "log level: debug, info, warn, error")
	flag.Parse()

	var level slog.Level
	if err := level.UnmarshalText([]byte(*logLevel)); err != nil {
		slog.Error("invalid log level", "err", err)
		os.Exit(1)
	}

	slog.SetDefault(slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: level,
		}),
	))

	godotenv.Load()
	conn, _, err := websocket.DefaultDialer.Dial(ds.GatewayURL, nil)

	if err != nil {
		slog.Error("Failed to connect!", "err", err)
		os.Exit(1)
	}

	defer conn.Close()

	slog.Info("Connected to the gateway")

	var sequence *int

	for {
		_, msg, err := conn.ReadMessage()

		if err != nil {
			slog.Error("Read error", "err", err)
			return
		}

		var payload payloads.GenericPayload

		if err := json.Unmarshal(msg, &payload); err != nil {
			slog.Error("Json decoding error", "err", err)
			continue
		}

		switch payload.Op {
		case payloads.OpHello:
			var hello payloads.Hello

			if err := json.Unmarshal(payload.D, &hello); err != nil {
				slog.Error("Json decoding error", "err", err)
				continue
			}
			slog.Info("Hello received", "hello", hello)

			if hello.HeartbeatInterval > 0 {
				go ds.StartHeartbeat(conn, hello.HeartbeatInterval, &sequence)
			}

			ds.Identify(conn, os.Getenv("TOKEN"))

		case payloads.OpDispatch:
			if payload.T == nil {
				slog.Error("Received dispatch event with null T field.")
				continue
			}

			switch *payload.T {
			case events.Ready:
				var ready payloads.Ready

				if err := json.Unmarshal(payload.D, &ready); err != nil {
					slog.Error("Json decoding error", "err", err)
					continue
				}
				slog.Info("Ready received", "ready", ready)

			default:
				slog.Warn("Unknown event", "event", *payload.T)
			}

		case payloads.OpHeartbeatACK:
			slog.Debug("Heartbeat ACK received")

		default:
			slog.Warn("Unknown op", "op", payload.Op)
		}
	}

}
