package main

import (
	"errors"
	"fmt"
	"math"
)

type MyError struct {
	message string
}

// implementing the error interface
func (err MyError) Error() string {
	return fmt.Sprintf("Error: %s", err.message)
}

type CustomError struct {
	code    int
	message string
	er      error
}

func (err CustomError) Error() string {
	return fmt.Sprintf("Error: %s, Code: %d, Inner Error: %v", err.message, err.code, err.er)
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, &MyError{message: "negative number"}
	}
	return math.Sqrt(x), nil
}

func process(data []byte) error {
	if len(data) == 0 {
		return errors.New("Empty Data")
	}
	return nil
}

func doSomething() error {
	er := internal()
	return CustomError{
		code:    404,
		message: "Not Found",
		er:      er,
	}
}

func internal() error {
	return errors.New("Internal Error")
}

func main() {
	result, err := sqrt(3)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("Result: %.2f\n", result)

	result, err = sqrt(-3)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	if err := process([]byte{}); err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	if err := doSomething(); err != nil {
		fmt.Printf("%v\n", err)
	}

}
