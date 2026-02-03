package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	// create base tickers to tick every 2 and 5 seconds
	printer := time.NewTicker(2 * time.Second)
	logger := time.NewTicker(5 * time.Second)

	// create context to later propagate cancellation to main and goroutine
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// also stop tickers when end of main is reached
	defer printer.Stop()
	defer logger.Stop()

	// start goroutine that pools for "updates every second and cancels with context"
	go poolUpdates(ctx)

	// main loop to handle ticks and context cancellation
	for {
		select {
		case tick := <-printer.C:
			fmt.Printf("Hello tick: %v\n", tick)
		case tick := <-logger.C:
			log.Printf("Logger tick: %v\n", tick)
		case <-ctx.Done():
			fmt.Printf("Timer expired, stopping tickers.\n")
			return
		}
	}
}

// poolUpdates simulates polling for updates every second until the context is done
func poolUpdates(ctx context.Context) {
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Context done\n")
			return
		case tick := <-ticker.C:
			fmt.Printf("Retrieving info at: %v\n", tick)
		}
	}
}
