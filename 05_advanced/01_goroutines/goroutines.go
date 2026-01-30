package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("Before say\n")
	go say("hello")
	fmt.Printf("After say\n")

	go printNumbers()
	go printLetter()

	time.Sleep(1 * time.Second)
}

func say(word string) {
	time.Sleep(1 * time.Second)
	fmt.Println(word)
}

func printNumbers() {
	for i := range 5 {
		fmt.Printf("Humber: %d %v\n", i, time.Now())
		time.Sleep(100 * time.Millisecond)
	}
}

func printLetter() {
	for _, char := range "abcde" {
		fmt.Printf("Letter: %c %v\n", char, time.Now())
		time.Sleep(100 * time.Millisecond)
	}
}
