package main

import "fmt"

func main() {
	// Basic arithmetics
	var a, b = 10, 6

	fmt.Println("Addition:", a+b)

	fmt.Println("Subtraction:", a-b)

	fmt.Println("Multiplication:", a*b)

	// When both operands are integers, the result is an integer
	fmt.Println("Division:", a/b)

	// When at least one operand is a float, the result is a float
	fmt.Println("Division:", float32(a)/float32(b))

	fmt.Println("Modulo:", a%b)

	// Overflow
	var maxInt8 int8 = 127
	fmt.Println("Overflow (127 + 1):", maxInt8+1)

	// Underflow
	var minInt8 int8 = -128
	fmt.Println("Underflow (-128 - 1):", minInt8-1)
}
