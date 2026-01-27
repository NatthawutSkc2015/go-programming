package main

import (
	"fmt"
	"sync"
)

// SafeCounter is a thread-safe counter
type SafeCounter struct {
	mu    sync.Mutex
	count int
}

// Inc increments the counter by 1 (thread-safe)
func (c *SafeCounter) Inc() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

// Value returns the current count value (thread-safe)
func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func main() {
	fmt.Println("=== Thread-Safe Counter Example ===")

	counter := &SafeCounter{}
	var wg sync.WaitGroup

	// Create 1000 goroutines to increment concurrently
	numGoroutines := 1000
	for range numGoroutines {
		wg.Go(func() {
			counter.Inc()
		})
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Check final value
	fmt.Printf("Expected: %d, Got: %d\n", numGoroutines, counter.Value())
	if counter.Value() == numGoroutines {
		fmt.Println("✓ No race condition - counter is accurate!")
	} else {
		fmt.Println("✗ Race condition detected!")
	}
}
