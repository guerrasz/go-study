package main

import "fmt"

func main() {
	var numbers = [5]int{1, 2, 3, 4, 5}

	slice := make([]int, 5)
	copy(slice, numbers[:])

	// Create a slice from the array
	numbersSlice := numbers[1:3]
	numbersSlice = append(numbersSlice, 6, 7, 8, 9)

	fmt.Println(numbersSlice)

	fmt.Println(cap(numbersSlice))
	fmt.Println(numbers)

	// 1. Create an array
	arr := [5]int{10, 20, 30, 40, 50}
	fmt.Printf("Original Array: %v, Address: %p\n", arr, &arr)

	// 2. Create a slice from the array
	// Pointer: &arr[1], Len: 2, Cap: 4 (from index 1 to the end)
	s := arr[1:3]
	fmt.Printf("\nSlice s (arr[1:3]): %v, Len: %d, Cap: %d, Data Pointer: %p\n", s, len(s), cap(s), s)

	// 3. Append within capacity
	// This WILL modify the original array!
	s = append(s, 99)
	fmt.Println("\n--- After append(s, 99) (Within Capacity) ---")
	fmt.Printf("Slice s: %v, Data Pointer: %p\n", s, s)
	fmt.Printf("Original Array: %v (Note index 3 changed!)\n", arr)

	// 4. Append beyond capacity
	// This will trigger a reallocation
	fmt.Println("\n--- Appending more to trigger reallocation ---")
	s = append(s, 100, 101)
	fmt.Printf("Slice s: %v, Len: %d, Cap: %d, New Data Pointer: %p\n", s, len(s), cap(s), s)
	fmt.Printf("Original Array: %v (Unchanged now, slice is detached)\n", arr)

	// Modifying s now won't affect arr
	s[0] = 777
	fmt.Printf("\nSlice s modified: %v\n", s)
	fmt.Printf("Original Array remains: %v\n", arr)
}
