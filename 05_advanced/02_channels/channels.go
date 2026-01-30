package main

import "fmt"

func main() {
	// create channel
	greeting := make(chan string)

	// create go routine that sends data to channel  5 times
	go func() {

		for _, c := range "abcde" {
			greeting <- string(c)
		}
	}()

	// create for loop that makes receiver block until data is available
	for range 5 {
		receiver := <-greeting
		fmt.Printf("Receiver: %v\n", receiver)
	}

}
