package main

import (
	"fmt"
	"sync"
)

var (
	rwmu    sync.RWMutex
	counter int
)

func readCounter() {
	rwmu.RLock()
	fmt.Printf("Read Counter: %d\n", counter)
	rwmu.RUnlock()
}

func writeCounter(value int) {
	rwmu.Lock()
	counter = value
	rwmu.Unlock()
}

func main() {
	var wg sync.WaitGroup

	for range 5 {
		wg.Go(readCounter)
	}

	wg.Go(func() {
		writeCounter(42)
	})

	wg.Wait()
}
