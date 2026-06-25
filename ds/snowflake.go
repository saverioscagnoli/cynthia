package ds

import (
	"strconv"
	"sync"
	"time"
)

type SnowflakeGenerator struct {
	mu         sync.Mutex
	epoch      int64
	workerID   int64
	sequence   int64
	lastMillis int64
}

func NewSnowflakeGenerator(workerID, epoch int64) *SnowflakeGenerator {
	return &SnowflakeGenerator{
		workerID: workerID & 0x1F,
		epoch:    epoch,
	}
}

func (g *SnowflakeGenerator) Next() Snowflake {
	g.mu.Lock()

	defer g.mu.Unlock()

	ms := time.Now().UnixMilli() - g.epoch

	if ms == g.lastMillis {
		g.sequence = (g.sequence + 1) & 0xFFF

		if g.sequence == 0 {
			for ms <= g.lastMillis {
				ms = time.Now().UnixMilli() - g.epoch
			}
		}
	} else {
		g.sequence = 0
	}

	g.lastMillis = ms

	return strconv.FormatInt((ms<<22)|(g.workerID<<17)|g.sequence, 10)
}

func (g *SnowflakeGenerator) IsValid(s Snowflake) bool {
	n, err := strconv.ParseInt(string(s), 10, 64)

	if err != nil || n <= 0 {
		return false
	}

	ms := n >> 22
	return ms > 0 && ms <= time.Now().UnixMilli()-g.epoch
}

func IsValidSnowflake(s Snowflake) bool {
	n, err := strconv.ParseInt(string(s), 10, 64)
	return err == nil && n > 0
}
