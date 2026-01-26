package main

import "fmt"

func main() {

	// Basic print statement
	fmt.Printf("Hello, Go!\n")

	// Formatting functions
	s := fmt.Sprint("Hello, Go!", 123, 456, "\n")
	fmt.Println(s)

	// Scanning functions
	var name string
	var age int

	fmt.Printf("Enter your name and age:\n")
	fmt.Scan(&name, &age)

	fmt.Printf("Name: %s, Age: %d\n", name, age)

	// Error formatting
	err := checkAge(15)
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func checkAge(age int) error {
	if age < 18 {
		return fmt.Errorf("Age %d is too young to drive", age)
	}
	return nil
}
