package gateway

import (
	"bytes"
	"cynthia/api"
	"cynthia/dstypes"
	"cynthia/events"
	"cynthia/payloads"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

// DispatchHandler is called for OP 0 (Dispatch) gateway events.
type DispatchHandler func(c *Client, p json.RawMessage)

type Client struct {
	mu               sync.RWMutex
	dispatchHandlers map[events.EventName][]DispatchHandler
	sequence         int
	token            string
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

func (c *Client) Start(token string, intents dstypes.Intents) *error {
	c.token = token
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

				Identify(conn, token, intents)
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

func (c *Client) SendMessage(channelID dstypes.Snowflake, content string) error {
	url := api.EndpointCreateMessage(channelID)
	body, err := json.Marshal(map[string]any{"content": content})

	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))

	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bot "+c.token)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("Unexpected status when sending a message: %s", res.Status)
	}

	return nil
}
