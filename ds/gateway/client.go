package gateway

import (
	"bytes"
	"cynthia/ds/dsapi"
	"cynthia/ds/dstypes"
	"cynthia/ds/events"
	"cynthia/ds/payloads"
	"cynthia/util"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
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
	dispatchHandlers map[events.EventName][]DispatchHandler
	sequence         int
	lastHeartbeat    atomic.Int64
	latency          atomic.Int64
	http             *http.Client
}

func NewClient(token string, appID dstypes.Snowflake) *Client {
	return &Client{
		token:            token,
		appID:            appID,
		dispatchHandlers: make(map[events.EventName][]DispatchHandler),
		sequence:         0,
		http:             &http.Client{},
	}
}

func (c *Client) do(method, url string, body any) ([]byte, error) {
	var buf io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		buf = bytes.NewBuffer(b)
	}

	req, err := http.NewRequest(method, url, buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bot "+c.token)
	req.Header.Set("Content-Type", "application/json")

	res, err := c.http.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return nil, fmt.Errorf("discord API error %d: %s", res.StatusCode, string(data))
	}

	return data, nil
}

func (c *Client) addHandler(event events.EventName, handler DispatchHandler) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.dispatchHandlers[event] = append(c.dispatchHandlers[event], handler)
}

func on[T any](c *Client, event events.EventName, handler func(c *Client, p T)) {
	c.addHandler(event, func(c *Client, raw json.RawMessage) {
		var p T
		if err := json.Unmarshal(raw, &p); err != nil {
			slog.Error("Failed to unmarshal payload", "event", event, "err", err)
			return
		}
		handler(c, p)
	})
}

func (c *Client) Emit(p payloads.GenericPayload) {
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
					go StartHeartbeat(conn, hello.HeartbeatInterval, &c.sequence, func() {
						c.lastHeartbeat.Store(time.Now().UnixNano())
					})
				}

				Identify(conn, c.token, intents)
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

func (c *Client) SendMessage(channelID dstypes.Snowflake, content string) error {
	return dsapi.SendMessageContent(c.token, channelID, content)
}

func (c *Client) SendEmbed(channelID dstypes.Snowflake, embed dstypes.Embed) error {
	e := util.NewVector[dstypes.Embed]()
	e.Push(embed)

	return dsapi.SendMessageEmbeds(c.token, channelID, e)
}

func (c *Client) SendEmbeds(channelID dstypes.Snowflake, embeds *util.Vector[dstypes.Embed]) error {
	return dsapi.SendMessageEmbeds(c.token, channelID, embeds)
}

func (c *Client) CreateGuildCommand(body dsapi.CreateGuildCommandBody, guildID dstypes.Snowflake) error {
	_, err := c.do("POST", dsapi.EndpointRegisterGuildCommand(c.appID, guildID), body)
	return err
}
