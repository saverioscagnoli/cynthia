package gateway

import (
	"cynthia/ds/dsapi"
	"cynthia/ds/dsevents"
	"cynthia/ds/dstypes"
	"cynthia/ds/payloads"
	"encoding/json"
	"log/slog"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"
)

// DispatchHandler is called for OP 0 (Dispatch) gateway events.
type DispatchHandler func(c *Client, p json.RawMessage)

type Client struct {
	token            string
	appID            dstypes.Snowflake
	mu               sync.RWMutex
	wmu              sync.Mutex
	dispatchHandlers map[dsevents.EventName][]DispatchHandler
	sequence         int
	lastHeartbeat    atomic.Int64
	latency          atomic.Int64
	Api              *dsapi.Client
}

func NewClient(token string, appID dstypes.Snowflake) *Client {
	return &Client{
		token:            token,
		appID:            appID,
		dispatchHandlers: make(map[dsevents.EventName][]DispatchHandler),
		sequence:         0,
		Api:              dsapi.NewApiClient(token, appID),
	}
}

func (c *Client) writeJSON(conn *websocket.Conn, v any) error {
	c.wmu.Lock()
	defer c.wmu.Unlock()
	return conn.WriteJSON(v)
}

func (c *Client) addHandler(event dsevents.EventName, handler DispatchHandler) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.dispatchHandlers[event] = append(c.dispatchHandlers[event], handler)
}

func On[T any](c *Client, event dsevents.EventName, handler func(c *Client, p T)) {
	c.addHandler(event, func(c *Client, raw json.RawMessage) {
		var p T
		if err := json.Unmarshal(raw, &p); err != nil {
			slog.Error("Failed to unmarshal payload", "event", event, "err", err)
			return
		}
		handler(c, p)
	})
}

func (c *Client) emit(p payloads.GenericPayload) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	for _, handler := range c.dispatchHandlers[*p.T] {
		handler(c, p.D)
	}
}

func (c *Client) Start(intents dstypes.Intents) *error {
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
					go c.StartHeartbeat(conn, hello.HeartbeatInterval, &c.sequence, func() {
						c.lastHeartbeat.Store(time.Now().UnixNano())
					})
				}

				c.Identify(conn, c.token, intents)
			}

		case payloads.OpDispatch:
			{
				if payload.T == nil {
					slog.Warn("Dispatch payload received without a type.", "payload", payload)
					continue
				}

				go c.emit(payload)
			}

		case payloads.OpHeartbeatACK:
			{

				latency := time.Now().UnixNano() - c.lastHeartbeat.Load()

				c.latency.Store(latency)
				slog.Debug("Heartbeat ack received.")
			}

		default:
			{
				slog.Warn("Unhandled opcode.", "op", payload.Op)
			}
		}
	}
}

func (c *Client) Latency() time.Duration {
	return time.Duration(c.latency.Load())
}
