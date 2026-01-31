package main

import (
	"fmt"
	"time"
)

func main() {
	// create data stream channel
	data := make(chan string)

	// goroutines will loop 5 times sending data
	go func() {
		for i := range 5 {
			data <- fmt.Sprintf("hello %d", i)
			time.Sleep(100 * time.Millisecond)
		}
		// close channel after loop ends so consuming the channel doesn't errors
		close(data)
	}()

	// receive data from active channel
	for value := range data {
		fmt.Printf("Received: %s\n", value)
	}
}

//& synchronizing the number of goroutines with num of receivers
// func main() {
// 	numGoroutines := 3
// 	done := make(chan int, 3)

// create 3 go routines that work for 1 second and then send the id to the channel
// 	for i := range numGoroutines {
// 		go func(id int) {
// 			fmt.Printf("Goroutine %d working...\n", id)
// 			time.Sleep(1 * time.Second)
// 			done <- id
// 		}(i)
// 	}

// wait for all go routines to finish
// 	for range numGoroutines {
// 		<-done
// 	}

// 	fmt.Printf("All go routines are finished\n")
// }

//& basic blocking concept
// func main() {
// 	done := make(chan bool)

// 	go func() {
// 		fmt.Printf("Working...\n")
// 		time.Sleep(2 * time.Second)
// 		done <- true
// 	}()

// 	<-done

// 	fmt.Printf("Done\n")

// }
