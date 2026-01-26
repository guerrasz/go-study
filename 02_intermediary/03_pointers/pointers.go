package main

import "fmt"

func main() {
	// var ptr *int
	// zero value of a pointer is nil

	var ptr *int
	var a int = 10
	ptr = &a

	// address because of &
	fmt.Printf("Address of a: %p\n", &a)
	// also address because v of pointer is the address
	fmt.Printf("Value of ptr: %v\n", ptr)
	// dereferecing the pointer to get the value
	fmt.Printf("Value of *ptr: %d\n", *ptr)

	// value is changed because we change the address
	modifyValue(ptr)
	fmt.Printf("Value of *ptr: %d\n", *ptr)

}

// increments the value by 1
func modifyValue(ptr *int) {
	*ptr++
}
