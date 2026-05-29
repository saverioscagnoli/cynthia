package ds

import (
	"cynthia/ds/events"
	"cynthia/ds/payloads"
	"encoding/json"
	"log/slog"
)

// DispatchEvent binds a Discord gateway dispatch event name (e.g. "READY")
// to the Go type used to unmarshal its payload.
//
// This lets you register type-safe handlers without having to manually
// json.Unmarshal in every handler.
//
// Usage:
//
//	client.AddTypedHandler(ds.ReadyEvent, func(c *ds.Client, p payloads.Ready) { ... })
//
// Internally handlers are still stored by event name, but the generic type
// parameter ensures your handler receives the right payload type.
type DispatchEvent[P any] struct {
	Name events.EventName
}

// AddHandler registers a type-safe handler for this dispatch event.
//
// Note: Go doesn't allow methods with their own type parameters, so the
// generic type parameter lives on DispatchEvent[P], not on Client.
func (e DispatchEvent[P]) AddHandler(c *Client, handler func(c *Client, p P)) {
	c.AddHandler(e.Name, func(c *Client, raw json.RawMessage) {
		var p P
		if err := json.Unmarshal(raw, &p); err != nil {
			slog.Error("Failed to unmarshal dispatch payload", "event", e.Name, "err", err)
			return
		}

		handler(c, p)
	})
}

var (
	ReadyEvent = DispatchEvent[payloads.Ready]{Name: events.Ready}
)
