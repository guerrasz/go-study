package main

import "fmt"

func main() {
	// Comprate againts a score and print grade
	var score int = 85

	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else if score >= 70 {
		fmt.Println("Grade: C")
	} else if score >= 60 {
		fmt.Println("Grade: D")
	} else {
		fmt.Println("Grade: F")
	}

	num := 18

	// check if number is divisible by 2 and 3 in nested if
	if num%2 == 0 {
		fmt.Println("Number is divisible by 2")
		if num%3 == 0 {
			fmt.Println("Number is divisible by 3")
		} else {
			fmt.Println("Number is not divisible by 3")
		}
	} else {
		fmt.Println("Number is not divisible by 2")
	}

	// Simple switch statement
	fruit := "Apple"

	switch fruit {
	case "Apple":
		fmt.Println("Apple")
	case "Banana":
		fmt.Println("Banana")
	default:
		fmt.Println("Unknown fruit")
	}

	day := "Monday"

	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("Weekday")
	case "Saturday", "Sunday":
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid day")
	}

	// Switch case with fallthrough

	number := 2

	switch {
	case number > 1:
		fmt.Println("Number is greater than or equal to 1")
		fallthrough
	case number == 2:
		fmt.Println("Number is equal to 2")
	default:
		fmt.Println("Number is not greater than or equal to 1")
	}

	checkType(number)
	checkType(fruit)
}

func checkType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("x is an int")
	case string:
		fmt.Println("x is a string")
	default:
		fmt.Println("x is of unknown type")
	}
}
