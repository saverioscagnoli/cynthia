package ds

import (
	"strconv"
	"testing"
)

func TestSnowflakeUniqueness(t *testing.T) {
	gen := NewSnowflakeGenerator(1, 0)
	seen := make(map[Snowflake]bool)

	for i := 0; i < 10000; i++ {
		id := gen.Next()
		if seen[id] {
			t.Fatalf("duplicate snowflake: %s", id)
		}
		seen[id] = true
	}
}

func TestSnowflakeMonotonic(t *testing.T) {
	gen := NewSnowflakeGenerator(1, 0)

	prev, _ := strconv.ParseInt(string(gen.Next()), 10, 64)
	for i := 0; i < 1_000_000; i++ {
		curr, _ := strconv.ParseInt(string(gen.Next()), 10, 64)
		if curr <= prev {
			t.Fatalf("snowflake not monotonically increasing: %d <= %d", curr, prev)
		}
		prev = curr
	}
}

func TestSnowflakeWorkerID(t *testing.T) {
	gen := NewSnowflakeGenerator(5, 0)
	id, _ := strconv.ParseInt(string(gen.Next()), 10, 64)
	workerID := (id >> 17) & 0x1F
	if workerID != 5 {
		t.Fatalf("expected workerID 5, got %d", workerID)
	}
}

func TestSnowflakeConcurrent(t *testing.T) {
	gen := NewSnowflakeGenerator(1, 0)
	seen := make(map[Snowflake]bool)
	ch := make(chan Snowflake, 10000)

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				ch <- gen.Next()
			}
		}()
	}

	for i := 0; i < 10000; i++ {
		id := <-ch
		if seen[id] {
			t.Fatalf("concurrent duplicate snowflake: %s", id)
		}
		seen[id] = true
	}
}
