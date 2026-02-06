package main

import (
	"fmt"
	"sync"
)

var (
	once sync.Once
)

func initialize() {
	fmt.Printf("This won't repeat\n")
}

func main() {

	var wg sync.WaitGroup

	for i := range 5 {
		wg.Go(func() {
			fmt.Printf("Goroutine %d is calling once.Do\n", i)
			once.Do(initialize)
		})
	}

	wg.Wait()
}
