package ds

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"
)

const GatewayURL = "wss://gateway.discord.gg/"

type Client struct {
	token         string
	appID         Snowflake
	intents       Intent
	handlers      map[string][]any
	conn          *websocket.Conn
	mu            sync.Mutex
	Commands      map[string]SlashCommand
	collectors    *collectorRegistry
	sequence      int
	lastHeartbeat atomic.Int64
	latency       atomic.Int64
	logger        *slog.Logger
	Api           ApiClient
}

type ClientOption func(*Client)

func WithLogger(logger *slog.Logger) ClientOption {
	return func(c *Client) {
		c.logger = logger
	}
}

func WithIntents(intents Intent) ClientOption {
	return func(c *Client) {
		c.intents = intents
	}
}

func NewClient(token string, appID Snowflake, options ...ClientOption) *Client {
	c := &Client{
		token:      token,
		appID:      appID,
		intents:    0,
		handlers:   make(map[string][]any),
		Commands:   make(map[string]SlashCommand),
		collectors: newCollectorRegistry(),
		logger:     slog.Default(),
		Api:        *newApiClient(token, appID),
	}

	for _, opt := range options {
		opt(c)
	}

	if c.intents == 0 {
		c.logger.Warn("No intents set — the bot will receive no events. Use WithIntents() to configure them, e.g. WithIntents(ds.IntentGuilds | ds.IntentGuildMessages)")
	}

	On(c, EventInteractionCreate, func(client *Client, i *InteractionCreate) {
		if i.Type == InteractionTypeMessageComponent {
			c.collectors.dispatch(c, i)
		}
	})

	return c
}

func (c *Client) writeJSON(v any) error {
	if c.conn == nil {
		return fmt.Errorf("Websocket is nil, call client.Login() first.")
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	return c.conn.WriteJSON(v)
}

func (c *Client) startHeartbeat(intervalMs int, seq *int, onSend func()) {
	send := func() {
		onSend()

		var seqVal *int = nil

		if seq != nil {
			seqVal = seq
		}

		if err := c.writeJSON(Payload[*int]{Op: OpHeartbeat, D: seqVal}); err != nil {
			c.logger.Error("Failed to send heartbeat.", "err", err)
			return
		}

		c.logger.Debug("Heartbeat sent.")
	}

	// First send to measure latency
	send()

	ticker := time.NewTicker(time.Duration(intervalMs) * time.Millisecond)

	defer ticker.Stop()

	for range ticker.C {
		send()
	}
}

func (c *Client) identify(token string, intents Intent) {
	payload := &Payload[IdentifyPayload]{
		Op: OpIdentify,
		D: IdentifyPayload{
			Token:   fmt.Sprintf("Bot %s", token),
			Intents: intents,
			Properties: IdentifyProperties{
				Os:      "linux",
				Browser: "chrome",
				Device:  "chrome",
			},
		},
	}

	if err := c.writeJSON(payload); err != nil {
		c.logger.Error("Gateway identification failed.", "err", err)
		return
	}

	c.logger.Info("Identification sent.")
}

func (c *Client) Login() error {
	conn, _, err := websocket.DefaultDialer.Dial(GatewayURL, nil)

	if err != nil {
		return err
	}

	defer conn.Close()

	c.conn = conn
	c.logger.Info("Connected to the discord gateway. Awaiting hello opcode.")

	for {
		_, message, err := conn.ReadMessage()

		if err != nil {
			return err
		}

		var payload Payload[json.RawMessage]

		if err := json.Unmarshal(message, &payload); err != nil {
			c.logger.Error("Failed to unmarshal gateway message", "err", err)
			continue
		}

		if payload.S != nil {
			c.sequence = *payload.S
		}

		switch payload.Op {
		case OpHello:
			{
				var hello HelloPayload

				if err := json.Unmarshal(payload.D, &hello); err != nil {
					return err
				}

				c.logger.Info("Hello event received.", "payload", hello)

				if hello.HeartbeatInterval <= 0 {
					c.logger.Error("Received an impossible heartbeat interval.", "interval", hello.HeartbeatInterval)
					continue
				}

				go c.startHeartbeat(hello.HeartbeatInterval, &c.sequence, func() {
					c.lastHeartbeat.Store(time.Now().UnixNano())
				})

				c.identify(c.token, c.intents)
			}

		case OpDispatch:
			{
				if payload.T == nil {
					c.logger.Error("Received a dispatch event without event name.", "payload", payload)
					continue
				}

				event := *payload.T
				c.logger.Debug("Received dispatch event", "event", event)

				if d, ok := dispatchers[event]; ok {
					d(c, event, payload.D)
				} else {
					c.logger.Warn("Unhandled dispatch event", "event", event)
				}

			}

		case OpHeartbeatACK:
			{
				ping := time.Now().UnixNano() - c.lastHeartbeat.Load()

				c.latency.Store(ping)
				c.logger.Debug("Heartbeat ack received.", "ping", c.Latency())
			}

		default:
			{
				c.logger.Warn("Unhandled opcode", "op", payload.Op)
			}
		}
	}
}

func (c *Client) Latency() time.Duration {
	return time.Duration(c.latency.Load())
}
