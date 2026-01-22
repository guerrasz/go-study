package main

import "fmt"

func main() {
	message := "Hello, World!"

	// Shows unicode values associated with characters
	for i, v := range message {
		fmt.Printf("%c is: %d in index %d\n", v, v, i)
	}

	// Range over a slice
	nums := []int{10, 20, 30}
	for i, num := range nums {
		fmt.Printf("Index %d: %d\n", i, num)
	}

	// Range over a map
	fruits := map[string]int{"apple": 5, "orange": 10}
	for name, count := range fruits {
		fmt.Printf("%s: %d\n", name, count)
	}

}
