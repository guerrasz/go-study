package main

import (
	"fmt"
	"time"
)

// NOTE in this example if timeout is reached before task is done, task is terminated
func main() {
	timeout := time.After(3 * time.Second)
	done := make(chan bool)

	go func() {
		task()
		done <- true
	}()

	// NOTE holds execution until one of the channels receive a value
	select {
	case <-timeout:
		fmt.Println("Timeout reached, terminating task.")
	case <-done:
		fmt.Println("Task completed successfully.")
	}
}

func task() {
	for i := range 20 {
		fmt.Printf("%d\n", i)
		time.Sleep(time.Second)
	}
}

// NOTE: Basic timer example
// func main() {
// 	timer := time.NewTimer(2 * time.Second) // starts ticking immediately but is not blocking

// 	success := timer.Stop()
// 	if success {
// 		fmt.Printf("Timer stopped: %t\n", success)
// 	}

// 	success = timer.Reset(time.Second)

// 	if !success {
// 		fmt.Printf("Timer reset: %t\n", success)
// 	}
// 	<-timer.C //NOTE: blocking operation until channel receives value

// 	fmt.Printf("Timer expired\n")
// }
