package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 2)

	go func() {
		time.Sleep(1 * time.Second)
		ch <- 1
	}()

	// will block until a value is received
	fmt.Printf("Received: %d\n", <-ch)
}
