package main

import (
	"fmt"
	"sync"
	"time"
)

// define LeakyBucket struct
type LeakyBucket struct {
	// maximum token capacity of the bucket
	capacity int
	// how often tokens are added to the bucket
	leakRate time.Duration
	// current number of tokens in the bucket
	tokens int
	// last time tokens were added
	lastLeak time.Time
	mu       sync.Mutex
}

func NewLeakyBucket(capacity int, leakRate time.Duration) *LeakyBucket {
	// initialize and return a new full LeakyBucket
	return &LeakyBucket{
		capacity: capacity,
		leakRate: leakRate,
		tokens:   capacity,
		lastLeak: time.Now(),
	}
}

func (lb *LeakyBucket) Allow() bool {
	// lock
	lb.mu.Lock()
	defer lb.mu.Unlock()

	// calculate how many tokens to add based on elapsed time since last leak
	now := time.Now()
	// elapsed time since last leak
	elapsed := now.Sub(lb.lastLeak)

	// determine how many tokens to add if elapsed > leakRate then add tokens
	tokensToAdd := int(elapsed / lb.leakRate)

	// add tokens
	lb.tokens += tokensToAdd

	// cap tokens to capacity
	if lb.tokens > lb.capacity {
		lb.tokens = lb.capacity
	}

	// update last leak time to
	lb.lastLeak = lb.lastLeak.Add(time.Duration(tokensToAdd) * lb.leakRate)

	// if there are tokens available, consume one and allow the request
	if lb.tokens > 0 {
		lb.tokens--
		return true
	}

	// no tokens available, deny the request
	return false
}

func main() {
	leaky_bucket := NewLeakyBucket(5, 500*time.Millisecond)

	for range 10 {
		if leaky_bucket.Allow() {
			fmt.Printf("Request accepted\n")
		} else {
			fmt.Printf("Request denied\n")
		}
		time.Sleep(200 * time.Millisecond)
	}
}
