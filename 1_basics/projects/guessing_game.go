// Package main implements a simple command-line number guessing game.
// It demonstrates basic Go concepts like random number generation,
// loops, conditional logic, and user input using fmt.Scanln.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Initialize the random number generator with a seed based on the current time.
	// This ensures we get a different sequence of numbers every time the game runs.
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	// Generate a secret random number between 1 and 100.
	secretNumber := random.Intn(100) + 1

	// Display the initial game instructions to the user.
	fmt.Println("Welcome to the guessing game!")
	fmt.Println("I have chosen a number between 1 and 100.")
	fmt.Println("Can you guess what it is?")

	var guess int
	// Start an infinite loop that continues until the user guesses the correct number.
	for {
		fmt.Print("Enter your guess: ")
		// fmt.Scanln requires a pointer (&guess) so it can modify the variable's value
		// directly in memory. Without the &, it would receive a copy and the original
		// variable wouldn't be updated.
		fmt.Scanln(&guess)

		// Compare the user's guess with the secret number.
		if guess < secretNumber {
			fmt.Println("Too low! Try again.")
		} else if guess > secretNumber {
			fmt.Println("Too high! Try again.")
		} else {
			// Correct guess! congratulate the user and exit the loop.
			fmt.Println("🎉 Congratulations! You guessed it!")
			break
		}
	}
}
