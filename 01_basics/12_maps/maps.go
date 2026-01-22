package main

import "fmt"

func main() {
	// 1. Declare and initialize with make (READY FOR WRITING)
	myMap := make(map[string]int)

	// 2. Declare without initialization (NIL MAP)
	// var nilMap map[string]int
	// Writing to nilMap would cause a panic!
	// But it's useful for struct fields or return types.
	// 1. Declare a nil map using 'var'
	var nilMap map[string]int

	fmt.Println("--- Nil Map Behavior ---")
	fmt.Printf("Value: %v, is nil: %v, length: %d\n", nilMap, nilMap == nil, len(nilMap))

	// 2. Reading from a nil map is SAFE
	// It returns the zero value of the map's value type (int -> 0)
	fmt.Println("Reading 'key':", nilMap["key"])

	// 3. Deleting from a nil map is SAFE
	delete(nilMap, "key")
	fmt.Println("Delete performed (noop)")

	// 4. Writing to a nil map triggers a PANIC
	fmt.Println("\nAttempting to write will panic. Uncomment the line below to see it:")
	// nilMap["key"] = 1 // <--- PANIC: assignment to entry in nil map

	// 5. Why use it?
	fmt.Println("\nUse cases:")
	fmt.Println("- Struct fields (lazy allocation)")
	fmt.Println("- Returning nil on error")
	fmt.Println("- Comparing 'if m == nil'")

	// Populate map
	myMap["one"] = 1
	myMap["two"] = 2
	myMap["three"] = 3

	// Display map
	fmt.Println(myMap)
	fmt.Println(myMap["one"])

	// Delete element
	delete(myMap, "one")
	fmt.Println(myMap)

	// Clear map
	// clear(myMap)
	// fmt.Println(myMap)

	value, ok := myMap["two"]
	fmt.Println(value, ok)

	// Declare another map but literal init
	otherMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	fmt.Println(otherMap)

	otherMap["four"] = 4
	fmt.Println(otherMap)

	for k, v := range otherMap {
		fmt.Println(k, v)
	}
}
