package main

import "fmt"

func main() {
	// For loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	numbers := []int{1, 2, 3, 4, 5}

	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// For loop that skips even numbers and breaks at 9
	// Can also be for i := range 10
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		if i == 9 {
			break
		}
		fmt.Println(i)
	}

	rows := 5

	// Draw a triangle with * using for loop
	for i := 1; i <= rows; i++ {
		// Print spaces
		for j := 1; j <= rows-i; j++ {
			fmt.Print(" ")
		}
		// Print stars
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		// Print new line
		fmt.Println()
	}

	// Using for as a while loop
	count := 0
	for count < 5 {
		fmt.Println(count)
		count++
	}

	// or without a condition
	sum := 0
	for {
		sum += 10
		fmt.Println(sum)
		if sum == 50 {
			break
		}
	}
}
