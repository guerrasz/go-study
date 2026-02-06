package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	var wg sync.WaitGroup
	// print the current process ID
	fmt.Printf("PID: %v\n", os.Getpid())

	// create a channel to receive OS signals
	sigs := make(chan os.Signal, 1)

	// channel to notify when to exit to other goroutine
	done := make(chan bool, 1)

	// notify the channel on SIGINT and SIGTERM signals
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// start a goroutine to handle the received signal
	wg.Go(func() {
		sig := <-sigs
		fmt.Printf("Signal: %v Received\n", sig)
		done <- true
	})

	// start a goroutine to simulate program work and exit when done is has a value
	wg.Go(func() {
		for {
			select {
			case <-done:
				fmt.Printf("Exiting program.\n")
				return
			default:
				fmt.Printf("Program running... Press Ctrl+C to exit.\n")
				time.Sleep(2 * time.Second)
			}
		}
	})

	// wait for all goroutines to finish
	wg.Wait()
}
