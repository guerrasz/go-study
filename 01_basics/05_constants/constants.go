package main

import "fmt"

const pi float32 = 3.14

func main() {
	const days int8 = 7

	const (
		monday    = 1
		tuesday   = 2
		wednesday = 3
		thursday  = 4
		friday    = 5
		saturday  = 6
		sunday    = 7
	)

	fmt.Println(pi)
	fmt.Println(days)

	fmt.Println(monday)
	fmt.Println(tuesday)
	fmt.Println(wednesday)
	fmt.Println(thursday)
	fmt.Println(friday)
	fmt.Println(saturday)
	fmt.Println(sunday)
}
