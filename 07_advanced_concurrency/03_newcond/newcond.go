package main

import (
	"fmt"
	"sync"
	"time"
)

type Buffer struct {
	items    []int
	mu       sync.Mutex
	cond     *sync.Cond
	capacity int
	head     int // index of the first item
	tail     int // index where the next item will be added
}

func NewBuffer(size int) *Buffer {
	// create the buffer struct with a circular buffer approach
	b := &Buffer{
		items:    make([]int, size),
		head:     0,
		tail:     0,
		capacity: size,
	}

	// initialize the condition variable with the buffer's mutex (mutex implements Locker interface required by NewCond)
	b.cond = sync.NewCond(&b.mu)
	return b
}

func (b *Buffer) isFull() bool {
	return (b.tail+1)%b.capacity == b.head
}

func (b *Buffer) isEmpty() bool {
	return b.head == b.tail
}

func (b *Buffer) Produce(item int) {
	b.mu.Lock()
	defer b.mu.Unlock()

	for b.isFull() {
		fmt.Printf("buffer: %v, head:  %d, tail: %d\n", b.items, b.head, b.tail)
		b.cond.Wait() // wait until there is space in the buffer
	}

	fmt.Printf("cap: %v\n", cap(b.items))

	b.items[b.tail] = item // add item to buffer
	b.tail = (b.tail + 1) % b.capacity
	fmt.Printf("Produced: %d\n", item)
	b.cond.Signal() // signal that an item has been produced
}

func (b *Buffer) Consume() int {
	b.mu.Lock()
	defer b.mu.Unlock()

	for b.isEmpty() {
		b.cond.Wait() // wait until there is at least one item to consume
	}

	item := b.items[b.head]            // get the item at head
	b.head = (b.head + 1) % b.capacity // move head forward
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
	buffer := NewBuffer(3)

	var wg sync.WaitGroup

	wg.Add(2) // we have 2 goroutines to wait for

	go producer(buffer, &wg)

	go consumer(buffer, &wg)

	wg.Wait()
}
