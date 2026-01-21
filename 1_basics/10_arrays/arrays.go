package main

import "fmt"

func main() {
	numbers := [5]int{1, 6, 8, 4, 6}
	fmt.Println(numbers)

	for i, v := range numbers {
		fmt.Printf("Element at index %d is %d\n", i, v)
	}

	// Underscore is used to store unused values
	for _, v := range numbers {
		fmt.Println(v)
	}
}
