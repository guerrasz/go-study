package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "message"
		close(ch)
	}()

	// if first channel is empty after 3 seconds, timeout will be executed
	for {
		select {
		case msg, ok := <-ch:
			if !ok {
				fmt.Println("channel closed")
				return
			}
			fmt.Println(msg)
		case <-time.After(1 * time.Second):
			fmt.Println("timeout")
		}
	}

}
