package ds

import (
	"cynthia/ds/events"
	"cynthia/ds/payloads"
	"encoding/json"
	"log/slog"
	"sync"

	"github.com/gorilla/websocket"
)

// DispatchHandler is called for OP 0 (Dispatch) gateway events.
type DispatchHandler func(c *Client, p json.RawMessage)

type Client struct {
	mu               sync.RWMutex
	dispatchHandlers map[events.EventName][]DispatchHandler
	sequence         int
}

func NewClient() *Client {
	return &Client{
		dispatchHandlers: make(map[events.EventName][]DispatchHandler),
		sequence:         0,
	}
}

func (c *Client) AddHandler(event events.EventName, handler DispatchHandler) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.dispatchHandlers[event] = append(c.dispatchHandlers[event], handler)
}

func (c *Client) Emit(p payloads.GenericPayload) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	for _, handler := range c.dispatchHandlers[*p.T] {
		handler(c, p.D)
	}
}

func (c *Client) Start(token string) *error {
	conn, _, err := websocket.DefaultDialer.Dial(GatewayURL, nil)

	if err != nil {
		return &err
	}

	defer conn.Close()

	slog.Info("Connected to the discord gateway. Awaiting hello opcode.")

	for {
		_, message, err := conn.ReadMessage()

		if err != nil {
			return &err
		}

		var payload payloads.GenericPayload

		if err := json.Unmarshal(message, &payload); err != nil {
			return &err
		}

		if payload.S != nil {
			c.sequence = *payload.S
		}

		switch payload.Op {
		case payloads.OpHello:
			{
				var hello payloads.Hello

				if err := json.Unmarshal(payload.D, &hello); err != nil {
					return &err
				}

				slog.Info("Hello payload received.", "hello", hello)

				if hello.HeartbeatInterval > 0 {
					go StartHeartbeat(conn, hello.HeartbeatInterval, &c.sequence)
				}

				Identify(conn, token)
			}

		case payloads.OpDispatch:
			{
				if payload.T == nil {
					slog.Warn("Dispatch payload received without a type.", "payload", payload)
					continue
				}

				c.Emit(payload)
			}

		case payloads.OpHeartbeatACK:
			{
				slog.Debug("Heartbeat ack received.")
			}

		default:
			{
				slog.Warn("Unhandled opcode.", "op", payload.Op)
			}
		}

	}

}
