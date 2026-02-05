package main

import (
	"fmt"
	"time"
)

type StatefulWorker struct {
	count int
	ch    chan int
}

func (w *StatefulWorker) start() {
	go func() {
		for val := range w.ch {
			w.count += val
			fmt.Printf("Current count %d\n", w.count)
		}
	}()
}

func (w *StatefulWorker) send(value int) {
	w.ch <- value
}

func main() {
	// struct is created in main func
	stWorker := &StatefulWorker{
		ch: make(chan int),
	}

	// start the worker goroutine
	stWorker.start()

	// NOTE: goroutine is still running and maintaining state even out of main function scope

	// send values to the worker
	for i := range 5 {
		stWorker.send(i)
		time.Sleep(500 * time.Millisecond)
	}

}
