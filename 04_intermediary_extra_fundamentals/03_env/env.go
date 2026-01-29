package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// get env var
	fmt.Printf("Env var: %v\n", os.Getenv("USER"))

	// set env var
	err := os.Setenv("USER2", "Lucas")

	if err != nil {
		fmt.Printf("Error setting env var: %v\n", err)
		return
	}

	fmt.Printf("Env var: %v\n", os.Getenv("USER2"))

	err = os.Unsetenv("USER2")

	if err != nil {
		fmt.Printf("Error unsetting env var: %v\n", err)
		return
	}

	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Printf("%v: %v\n", pair[0], pair[1])
	}
}
