package main

import (
	"errors"
	"fmt"
)

func init() {
	fmt.Printf("Initializing package...\n")
}

func main() {
	fmt.Println(add(1, 2))

	// Atribute anonymous function to a variable
	greet := func() {
		fmt.Println("Hello Stanger")
	}

	greet()

	fmt.Println(applyAdd(1, 2, add))

	// Attribute return value to operation
	operation := createAdd()

	fmt.Println(operation(5, 8))

	// Multiple return values
	quotient, remainder := divide(10, 3)
	fmt.Println(quotient, remainder)

	// Error handling with mulitple return values
	result, err := compare(10, 10)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("%v\n", result)
	}

	// executed at the end. of main but before if proces2() panics
	defer fmt.Println(sum(1, 2, 3, 4, 5))

	process()

	process2(10)

	process2(2)

	process3()

	//& If this executes all defer statement that are queued (line 39) won't get executed
	// os.Exit(1)

	fmt.Printf("After panic\n")

}

// Add functions that simply returns the sum from a and b
func add(a int, b int) int {
	return a + b
}

// Function that receives a function
func applyAdd(a int, b int, add func(int, int) int) int {
	return add(a, b)
}

// Function that returns a function
func createAdd() func(int, int) int {
	return func(a int, b int) int {
		return a + b
	}
}

func divide(a int, b int) (int, int) {
	return a / b, a % b
}

func compare(a int, b int) (string, error) {
	if a > b {
		return "Greater", nil
	} else if a < b {
		return "Less", nil
	} else {
		//& Here string has to be "" because zero string value is "" instead of nil
		return "", errors.New("Unable to obtain greater")
	}
}

// Variadic function accepts dynamic number of arguments and functions like a list
// ^ Must be the last parameter
func sum(nums ...int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}

// Using defer keyworkd
func process() {
	defer fmt.Printf("First deferred statement\n")
	defer fmt.Printf("Second deferred statement\n")
	fmt.Printf("Normal statement\n")
}

// Using panick keyword
func process2(input int) {
	if input < 0 {
		panic("Input must be a non-negative number")
	}
	fmt.Printf("Input value is: %d\n", input)
}

// Using recover
func process3() {
	// This executed only when exiting the function to check if it panicked to not terminate
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	fmt.Println("Start process3")
	panic("Something went wrong")
}
