package main

import "fmt"

func example1() {
	ch := make(chan int, 2)

	// channel is full
	ch <- 1
	ch <- 2

	// goroutine starts
	go func() {

		fmt.Printf("Received: %d\n", <-ch)
	}()

	// buffer is full so execution is blocked until a value is received (happens in the goroutine)
	ch <- 3

	// buffered is emptied
	fmt.Printf("Received: %d\n", <-ch)
	fmt.Printf("Received: %d\n", <-ch)

}
