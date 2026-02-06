package main

import (
	"fmt"
	"sync"
	"time"
)

const BUFFERSIZE = 5

type Buffer struct {
	items []int
	mu    sync.Mutex
	cond  *sync.Cond
}

func NewBuffer() *Buffer {
	// create the buffer struct and initialize the items slice with the specified max size
	b := &Buffer{
		items: make([]int, 0, BUFFERSIZE),
	}

	// initialize the condition variable with the buffer's mutex (mutex imlements Locker interface required by NewCond)
	b.cond = sync.NewCond(&b.mu)
	return b
}

func (b *Buffer) Produce(item int) {
	b.mu.Lock()
	defer b.mu.Unlock()

	for len(b.items) == BUFFERSIZE {
		b.cond.Wait() // wait until there is space in the buffer
	}

	fmt.Printf("%v\n", cap(b.items))

	b.items = append(b.items, item) // add item to buffer
	fmt.Printf("Produced: %d\n", item)
	b.cond.Signal() // signal that an item has been produced
}

func (b *Buffer) Consume() int {
	b.mu.Lock()
	defer b.mu.Unlock()

	for len(b.items) == 0 {
		b.cond.Wait() // wait until there is at least one item to consume
	}

	item := b.items[0]    // get the first item from the buffer
	b.items = b.items[1:] // remove the consumed item from the buffer
	fmt.Printf("Consumed: %d\n", item)
	b.cond.Signal()
	return item
}

func producer(buffer *Buffer, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := range 10 {
		buffer.Produce(i + 100)
		time.Sleep(100 * time.Millisecond)
	}
}

func consumer(buffer *Buffer, wg *sync.WaitGroup) {
	defer wg.Done()

	for range 10 {
		buffer.Consume()
		time.Sleep(200 * time.Millisecond)
	}
}

func main() {
	buffer := NewBuffer()

	var wg sync.WaitGroup

	wg.Add(2) // we have 2 goroutines to wait for

	go producer(buffer, &wg)

	go consumer(buffer, &wg)

	wg.Wait()
}
