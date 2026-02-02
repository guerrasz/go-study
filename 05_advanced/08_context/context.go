package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// base contexts
	todoContext := context.TODO()
	contextBkg := context.Background()

	// store key-value pairs in context
	ctx := context.WithValue(todoContext, "key", "value")
	fmt.Printf("Context: %v\n", ctx)

	ctx = context.WithValue(contextBkg, "key", "value")
	fmt.Printf("Context: %v\n", ctx)

	// set timeout for context
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)

	defer cancel()

	result := checkNumber(ctx, 10)

	time.Sleep(2 * time.Second)
	fmt.Printf("Result: %v\n", result)

	result = checkNumber(ctx, 15)
	fmt.Printf("Result: %v\n", result)

	// NOTE: define parent context
	parent := context.TODO()

	// NOTE: create a child context with value and timeout
	workCtx := context.WithValue(parent, "key", "value")
	workCtx, cancel = context.WithTimeout(workCtx, 2*time.Second)

	defer cancel()

	go doWork(workCtx)

	time.Sleep(3 * time.Second)

	//NOTE: okay way to define value -> value := workCtx.Value("key")

	if value := workCtx.Value("key"); value != nil {
		fmt.Printf("Work context key: %v\n", value)
	}

}

func checkNumber(ctx context.Context, num int) string {
	select {
	case <-ctx.Done():
		return "Context cancelled"
	default:
		if num%2 == 0 {
			return "Even"
		} else {
			return "Odd"
		}
	}
}

func doWork(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Work cancelled: %s\n", ctx.Err())
			return
		default:
			fmt.Printf("Working...\n")
		}
		time.Sleep(500 * time.Millisecond)
	}
}
