package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// create struct
type AtomicCounter struct {
	count int64
}

// method to increment the counter atomically
func (ac *AtomicCounter) Increment() {
	atomic.AddInt64(&ac.count, 1)
}

// method to get the current value of the counter atomically
func (ac *AtomicCounter) GetValue() int64 {
	return atomic.LoadInt64(&ac.count)
}

func main() {
	var wg sync.WaitGroup

	numGoroutines := 10

	counter := &AtomicCounter{}

	// start multiple goroutines to increment the counter
	wg.Add(numGoroutines)
	for range numGoroutines {
		go func() {
			defer wg.Done()
			for range 1000 {
				counter.Increment()
			}
		}()
	}

	// wait for all goroutines to finish
	wg.Wait()
	fmt.Printf("Final counter value %d\n", counter.GetValue())
}
