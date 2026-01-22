package main

import "fmt"

// Package level variables (only with var keyword)
var nickname = "guerra"

func main() {
	// More standard variable declarations
	var age int
	var name string = "Lucas"
	var middle_name = "Guerra"

	// Go way of declaring variables
	count := 10
	lastName := "Sardo"

	//& Variables have default values
	// int -> 0
	// string -> ""
	// bool -> false
	// float -> 0.0
	// pointer -> nil

	fmt.Println(age)
	fmt.Println(name)
	fmt.Println(middle_name)
	fmt.Println(count)
	fmt.Println(lastName)
	fmt.Println(nickname)

	// Overwrite package level variable
	nickname := "foo"
	fmt.Println(nickname)
}
