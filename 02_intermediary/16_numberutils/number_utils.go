package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// int parsing
	numStr := "64"
	num, err := strconv.Atoi(numStr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", num)

	// parse float
	floatNum, err := strconv.ParseFloat("64.5", 64)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%.2f\n", floatNum)

	// parse from binary
	binNum, err := strconv.ParseInt("1000", 2, 64)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", binNum)

	// fixed seed will not produce a random number
	val := rand.New(rand.NewSource(5))

	// better seeding
	val2 := rand.New(rand.NewSource(time.Now().UnixNano()))

	// fixed seed
	fmt.Printf("Fixed seed: %v\n", val.Intn(10))

	// fixed better seed
	fmt.Printf("Fixed better seed: %v\n", val2.Intn(10))

	// auto seed
	fmt.Printf("Auto seed: %v\n", rand.Intn(10))

	for {
		fmt.Println()
		fmt.Printf("Welcome to the dice game!\n")
		fmt.Printf("1. Roll the dice\n")
		fmt.Printf("2. Exit\n")
		fmt.Printf("Enter your choice: ")

		reader := bufio.NewReader(os.Stdin)
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		fmt.Println()

		switch choice {
		case "1":
			fmt.Printf("Rolling the dice...\n")
			fmt.Printf("You rolled: %v\n", rand.Intn(15))
		case "2":
			fmt.Printf("Goodbye!\n")
			os.Exit(0)
		default:
			fmt.Printf("Invalid choice!\n")
			continue
		}
	}
}
