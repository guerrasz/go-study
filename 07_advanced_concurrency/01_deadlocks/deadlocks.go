package main

import (
	"fmt"
	"sync"
	"time"
)

// depending on the order of locking m1 and m2 a deadlock may occur because 
// both goroutines may be waiting each other to access blocked mutexes so:
// go1 waiting for m2 and go2 waiting for m1
// to avoid deadlocks always lock mutexes in the same order
func main() {

	var m1, m2 sync.Mutex

	go func() {
		m1.Lock()
		fmt.Printf("Goroutine 1 locked m1\n")
		time.Sleep(time.Second)
		m1.Unlock()
		m2.Lock()
		fmt.Printf("Goroutine 1 locked m2\n")
		m2.Unlock()
		fmt.Printf("Goroutine 1 finished\n")
	}()

	go func() {
		m1.Lock()
		fmt.Printf("Goroutine 2 locked m1\n")
		time.Sleep(time.Second)
		m1.Unlock()
		m2.Lock()
		fmt.Printf("Goroutine 2 locked m2\n")
		m2.Unlock()
		fmt.Printf("Goroutine 2 finished\n")
	}()

	time.Sleep(3 * time.Second)
	fmt.Printf("Main function exiting\n")
}
