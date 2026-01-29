package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Command: %v\n", os.Args[0])
	fmt.Printf("Arguments: %v\n", os.Args[1:])

	// define flags
	var name string
	var age int
	var male bool

	flag.StringVar(&name, "name", "Lucas", "Name of the user")
	flag.IntVar(&age, "age", 20, "Age of the user")
	flag.BoolVar(&male, "male", true, "Gender of the user")

	flag.Parse()

	fmt.Printf("Name: %v\n", name)
	fmt.Printf("Age: %v\n", age)
	fmt.Printf("Male: %v\n", male)
}
