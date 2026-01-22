package main

import "fmt"

func main() {
	sequence := adder()
	// i should be out of scope here but it is still accessible because of the closure
	sequence()
	sequence()
	sequence()

	// closure that is executed returning another function that increases sum by x
	sum := 0
	fmt.Printf("Previous value of sum: %d\n", sum)

	closure := func(x int) int {
		sum += x
		fmt.Printf("Current value of sum: %d\n", sum)
		return sum

	}

	closure(2)
	closure(2)
	closure(2)
}

func adder() func() int {
	// i is defined here
	i := 0
	// initial value is 0
	fmt.Printf("Previous value of i: %d\n", i)
	// return a function that increments i
	return func() int {
		i++
		fmt.Printf("Current value of i: %d\n", i)
		return i
	}
}
