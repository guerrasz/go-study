package main

import (
	"fmt"
	"time"
)

// rate limiter using token bucket algorithm
type RateLimiter struct {
	tokens     chan struct{}
	refillTime time.Duration
}

// constructor for RateLimiter
func NewRateLimiter(rateLimit int, refillTime time.Duration) *RateLimiter {
	// create the struct and initialize the token channel
	rl := &RateLimiter{
		tokens:     make(chan struct{}, rateLimit),
		refillTime: refillTime,
	}
	// fill the token bucket
	for range rateLimit {
		rl.tokens <- struct{}{}
	}
	// start the refill goroutine that will add tokens at regular intervals (refillTime)
	go rl.StartRefill()
	// return the rate limiter
	return rl
}

// goroutine to refill tokens at regular intervals
func (rl *RateLimiter) StartRefill() {
	// initialize a ticker to trigger at the specified refill time
	ticker := time.NewTicker(rl.refillTime)
	// ensure the ticker is stopped when the function exits
	defer ticker.Stop()
	// continuously listen for ticker events to add tokens
	for range ticker.C {
		select {
		// add a token if there is space in the channel
		case rl.tokens <- struct{}{}:
		// if the channel is full, skip adding a token
		default:
		}
	}
}

// simple function to check if a request is allowed based on token availability
func (rl *RateLimiter) Allow() bool {
	select {
	// if a token is available, consume it and allow the request
	case <-rl.tokens:
		return true
	default:
		return false
	}
}

func token_bucket() {
	// create a rate limiter allowing 5 requests and adds tokens every second
	rateLimiter := NewRateLimiter(5, time.Second)

	// simulate 10 requests and check if they are allowed or denied
	for range 10 {
		if rateLimiter.Allow() {
			fmt.Println("Request allowed")
		} else {
			fmt.Println("Request denied")
		}
		// wait for 200 milliseconds before the next request
		time.Sleep(200 * time.Millisecond)
	}
}
