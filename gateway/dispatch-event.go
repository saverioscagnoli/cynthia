package gateway

import (
	"cynthia/dstypes"
	"cynthia/events"
	"cynthia/payloads"
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
	ReadyEvent                          = DispatchEvent[payloads.Ready]{Name: events.Ready}
	Resumed                             = DispatchEvent[payloads.Resumed]{Name: events.Resumed}
	Reconnect                           = DispatchEvent[payloads.Reconnect]{Name: events.Reconnect}
	InvalidSession                      = DispatchEvent[payloads.InvalidSession]{Name: events.InvalidSession}
	ApplicationCommandPermissionsUpdate = DispatchEvent[dstypes.ApplicationCommandPermissions]{Name: events.ApplicationCommandPermissionsUpdate}
	AutoModerationRuleCreate            = DispatchEvent[dstypes.AutoModerationRule]{Name: events.AutoModerationRuleCreate}
	AutoModerationRuleUpdate            = DispatchEvent[dstypes.AutoModerationRule]{Name: events.AutoModerationRuleUpdate}
	AutoModerationRuleDelete            = DispatchEvent[dstypes.AutoModerationRule]{Name: events.AutoModerationRuleDelete}
	AutoModerationActionExecution       = DispatchEvent[payloads.AutoModerationActionExecution]{Name: events.AutoModerationActionExecution}
	ChannelCreate                       = DispatchEvent[dstypes.Channel]{Name: events.ChannelCreate}
	ChannelUpdate                       = DispatchEvent[dstypes.Channel]{Name: events.ChannelUpdate}
	ChannelDelete                       = DispatchEvent[dstypes.Channel]{Name: events.ChannelDelete}
	ChannelInfo                         = DispatchEvent[payloads.ChannelInfo]{Name: events.ChannelInfo}
	ChannelPinsUpdate                   = DispatchEvent[payloads.ChannelPinsUpdate]{Name: events.ChannelPinsUpdate}
	ThreadCreate                        = DispatchEvent[payloads.ThreadCreate]{Name: events.ThreadCreate}
	ThreadUpdate                        = DispatchEvent[payloads.ThreadUpdate]{Name: events.ThreadUpdate}
	ThreadDelete                        = DispatchEvent[payloads.ThreadDelete]{Name: events.ThreadDelete}
	ThreadListSync                      = DispatchEvent[payloads.ThreadListSync]{Name: events.ThreadListSync}
	ThreadMemberUpdate                  = DispatchEvent[payloads.ThreadMemberUpdate]{Name: events.ThreadMemberUpdate}
	EntitlementCreate                   = DispatchEvent[payloads.EntitlementCreate]{Name: events.EntitlementCreate}
	EntitlementUpdate                   = DispatchEvent[payloads.EntitlementUpdate]{Name: events.EntitlementUpdate}
	EntitlementDelete                   = DispatchEvent[payloads.EntitlementDelete]{Name: events.EntitlementDelete}
	MessageCreate                       = DispatchEvent[payloads.MessageCreate]{Name: events.MessageCreate}
)
