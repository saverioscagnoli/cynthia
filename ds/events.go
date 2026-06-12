package ds

import "encoding/json"

type Event[T any] struct {
	name string
}

var (
	EventReady             = Event[Ready]{"READY"}
	EventMessageCreate     = Event[MessageCreate]{"MESSAGE_CREATE"}
	EventGuildCreate       = Event[GuildCreate]{"GUILD_CREATE"}
	EventInteractionCreate = Event[InteractionCreate]{"INTERACTION_CREATE"}
)

type eventDispatcher func(c *Client, eventName string, data json.RawMessage)

var dispatchers = map[string]eventDispatcher{
	"READY":              makeDispatcher[Ready](),
	"MESSAGE_CREATE":     makeDispatcher[MessageCreate](),
	"GUILD_CREATE":       makeDispatcher[GuildCreate](),
	"INTERACTION_CREATE": makeDispatcher[InteractionCreate](),
}

func makeDispatcher[T any]() eventDispatcher {
	return func(c *Client, eventName string, data json.RawMessage) {
		dispatch[T](c, eventName, data)
	}
}

func dispatch[T any](c *Client, eventName string, data json.RawMessage) {
	handlers, ok := c.handlers[eventName]

	if !ok {
		c.logger.Debug("Unable to find valid handlers for event", "event", eventName)
		return
	}

	var payload T

	if err := json.Unmarshal(data, &payload); err != nil {
		c.logger.Error("Failed to unmarshal event payload", "event", eventName, "err", err)
		return
	}

	for _, h := range handlers {
		handler, ok := h.(func(*Client, T))

		if !ok {
			continue
		}

		go handler(c, payload)
	}
}

func On[T any](c *Client, event Event[T], handler func(*Client, T)) {
	c.handlers[event.name] = append(c.handlers[event.name], handler)
}
