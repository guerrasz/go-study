package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go produce(ch1)
	go filter(ch1, ch2)

	for filtered := range ch2 {
		fmt.Printf("Value: %d\n", filtered)
	}
}

// produces numbers from 0 to 5 and sends them to channel
func produce(ch chan<- int) {
	for i := range 5 {
		ch <- i
	}
	// close the channel after everything has been sent 
	close(ch)
}

// Loops over channel which has the 5 nums called in produce()
// and filters them sending them to the other channel
func filter(source <-chan int, receiver chan<- int) {
	for value := range source {
		if value%2 == 0 {
			receiver <- value
		}
	}
	// we call this channel receiver but it receives the values and send them right away
	close(receiver)
}
