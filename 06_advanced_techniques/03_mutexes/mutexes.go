package main

import (
	"fmt"
	"sync"
)

// counter struct implementing a mutex to ensure safe concurrent access
type Counter struct {
	mu    sync.Mutex
	count int
}

// Increment method to safely increment the counter
func (c *Counter) Increment() {
	// Lock the mutex before modifying the count
	c.mu.Lock()
	// Ensure the mutex is unlocked after the increment operation
	defer c.mu.Unlock()
	c.count++
}

// Value method to safely retrieve the current value of the counter
func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func main() {
	// create wait group
	var wg sync.WaitGroup

	// create counter as a pointer
	counter := &Counter{}

	numGoroutines := 10

	// set the number of goroutines to wait for
	wg.Add(numGoroutines)

	// create 10 goroutines to increment the counter concurrently
	for range numGoroutines {
		// each goroutine increments the counter 1000 times
		go func() {
			// signal that this goroutine is done when it completes
			defer wg.Done()
			for range 1000 {
				counter.Increment()
			}
		}()
	}

	// wait for all goroutines to finish
	wg.Wait()

	fmt.Printf("Final counter value %d\n", counter.Value())
}
