package ds

import (
	"sync"
	"sync/atomic"
	"time"
)

const MaxCollectorTimeout = 15 * time.Minute

type ComponentInteractionCollector struct {
	messageID Snowflake
	handler   func(*Client, *InteractionCreate)
	onEnd     func()
	count     atomic.Uint32
	max       uint
	done      chan struct{}
	once      sync.Once
}

func (c *ComponentInteractionCollector) Stop() {
	if c.onEnd != nil {
		c.onEnd()
	}

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

	if c.max > 0 && uint(c.count.Add(1)) >= c.max {
		c.Stop()
	}

	return true
}

type CollectorOptions struct {
	Timeout time.Duration
	Max     uint
	Handler func(*Client, *InteractionCreate)
	OnEnd   func()
}

func (c *Client) CollectComponents(
	messageID Snowflake,
	o CollectorOptions,
) *ComponentInteractionCollector {
	col := &ComponentInteractionCollector{
		messageID: messageID,
		max:       o.Max,
		handler:   o.Handler,
		onEnd:     o.OnEnd,
		done:      make(chan struct{}),
	}

	if o.Timeout <= 0 {
		o.Timeout = 30 * time.Second
	} else if o.Timeout > MaxCollectorTimeout {
		c.logger.Warn("CollectComponents timeout exceeded maximum, clamping", "requested", o.Timeout, "max", MaxCollectorTimeout)
		o.Timeout = MaxCollectorTimeout
	}

	if o.Handler == nil {
		o.Handler = func(*Client, *InteractionCreate) {}
	}

	c.collectors.add(col)

	go func() {
		select {
		case <-time.After(o.Timeout):
			col.Stop()

		case <-col.done:
		}

		c.collectors.remove(messageID)
	}()

	return col
}
