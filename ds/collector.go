package ds

import (
	"sync"
	"sync/atomic"
	"time"
)

type ComponentInteractionCollector struct {
	messageID Snowflake
	handler   func(*Client, *InteractionCreate)
	count     atomic.Uint32
	max       uint
	done      chan struct{}
	once      sync.Once
}

func (c *ComponentInteractionCollector) Stop() {
	c.once.Do(func() { close(c.done) })
}

type collectorRegistry struct {
	mu         sync.RWMutex
	collectors map[Snowflake]*ComponentInteractionCollector
}

func newCollectorRegistry() *collectorRegistry {
	return &collectorRegistry{
		collectors: map[Snowflake]*ComponentInteractionCollector{},
	}
}

func (r *collectorRegistry) add(c *ComponentInteractionCollector) {
	r.mu.Lock()
	r.collectors[c.messageID] = c
	r.mu.Unlock()
}

func (r *collectorRegistry) remove(messageID Snowflake) {
	r.mu.Lock()
	delete(r.collectors, messageID)
	r.mu.Unlock()
}

func (r *collectorRegistry) dispatch(client *Client, i *InteractionCreate) bool {
	if i.Message == nil {
		return false
	}

	r.mu.RLock()
	c, ok := r.collectors[i.Message.ID]

	if !ok {
		return false
	}

	c.handler(client, i)

	if c.max > 0 {
		if uint(c.count.Add(1)) >= c.max {
			c.Stop()
		}
	}

	return true
}

func (c *Client) CollectComponents(
	messageID Snowflake,
	timeout time.Duration,
	max uint,
	handler func(*Client, *InteractionCreate),
	onExpire func(),
) *ComponentInteractionCollector {
	col := &ComponentInteractionCollector{
		messageID: messageID,
		max:       max,
		handler:   handler,
		done:      make(chan struct{}),
	}

	c.collectors.add(col)

	go func() {
		select {
		case <-time.After(timeout):
			col.Stop()

			if onExpire != nil {
				onExpire()
			}

		case <-col.done:
		}

		c.collectors.remove(messageID)
	}()

	return col
}
