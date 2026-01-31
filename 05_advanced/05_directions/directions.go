package main

import "fmt"

func main() {
	ch := make(chan int)

	go producer(ch)

	consumer(ch)
}

// producer only channel
func producer(ch chan<- int) {
	for i := range 5 {
		ch <- i
	}
	close(ch)
}

// receive only channel
func consumer(ch <-chan int) {
	for i := range ch {
		fmt.Println(i)
	}
}
