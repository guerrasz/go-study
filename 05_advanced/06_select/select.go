package main

import "fmt"

func selectExample() {

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		ch1 <- 1
	}()

	go func() {
		ch2 <- 2
	}()

	// evaluates if statements are true
	select {
	case msg := <-ch1:
		fmt.Printf("Message from ch1: %d", msg)
	case msg := <-ch2:
		fmt.Printf("Message from ch2: %d", msg)
	default:
		fmt.Println("No message")
	}

}
