package main

import (
	"fmt"
	"sync"
	"time"
)

// FixedWindowRateLimiter implements a fixed window rate limiting algorithm
type FixedWindowRateLimiter struct {
	// mu protects access to the rate limiter's state
	mu sync.Mutex
	// count is the number of requests in the current window
	count int
	// limit is the maximum number of allowed requests per window
	limit int
	// window is the duration of the time window
	window time.Duration
	// resetTime is the time when the current window resets
	resetTime time.Time
}

// NewFixedWindowRateLimiter creates a new FixedWindowRateLimiter
func NewFixedWindowRateLimiter(limit int, window time.Duration) *FixedWindowRateLimiter {
	return &FixedWindowRateLimiter{
		limit:  limit,
		window: window,
	}
}

// Allow checks if a request is allowed under the rate limit
func (rl *FixedWindowRateLimiter) Allow() bool {
	// lock the mutex to protect shared state
	rl.mu.Lock()
	// ensure the mutex is unlocked when the function exits
	defer rl.mu.Unlock()
	// get the current time
	now := time.Now()
	// if the current time is past the reset time, reset the count and set a new reset time
	if now.After(rl.resetTime) {
		rl.count = 0
		rl.resetTime = now.Add(rl.window)
	}
	// if the count is below the limit, increment the count and allow the request
	if rl.count < rl.limit {
		rl.count++
		return true
	}
	// otherwise, deny the request
	return false
}

func fixed_window() {
	var wg sync.WaitGroup
	// create a rate limiter allowing 5 requests every 2 seconds
	rateLimiter := NewFixedWindowRateLimiter(5, 1*time.Second)

	// simulate 10 requests and check if they are allowed or denied with a 200ms interval between requests
	for range 10 {
		// increment the wait group counter for each request
		wg.Go(func() {
			// check if the request is allowed and print the result
			if rateLimiter.Allow() {
				fmt.Printf("Request Allowed\n")
			} else {
				fmt.Printf("Request Denied\n")
			}
		})
		// wait 150ms before the next request
		time.Sleep(150 * time.Millisecond)
	}

	// wait for all requests to complete
	wg.Wait()
}
