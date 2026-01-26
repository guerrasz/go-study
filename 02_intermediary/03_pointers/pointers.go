package main

import "fmt"

func main() {
	// var ptr *int

	var ptr *int
	var a int = 10
	ptr = &a

	// address because of &
	fmt.Printf("Address of a: %p\n", &a)
	// also address because v of pointer is the address
	fmt.Printf("Value of ptr: %v\n", ptr)
	// dereferecing the pointer to get the value 
	fmt.Printf("Value of *ptr: %d\n", *ptr)

}
