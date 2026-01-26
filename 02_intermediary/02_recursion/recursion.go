package main

import "fmt"

func main() {
	fmt.Printf("Factorial of 5 is: %d\n", factorial(5))
	fmt.Printf("Sum of digits of 12345 is: %d\n", sumOfDigits(12345))
}

func factorial(n int) int {
	// base case 0! is 1
	if n == 0 {
		return 1
	}
	// recursive case n! is n * (n-1)!
	return n * factorial(n-1)
}

func sumOfDigits(n int) int {
	// base case n < 10
	if n < 10 {
		return n
	}
	// recursive case n > 10 so we sum the last digit with the sum of the other digits
	return n%10 + sumOfDigits(n/10)
}
