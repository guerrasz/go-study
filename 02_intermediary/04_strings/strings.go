package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// normal string, \n is escaped
	message := "Hello\nWorld!"
	fmt.Printf("%s\n", message)

	// \t is a tab
	message = "Hello\tWorld!"
	fmt.Printf("%s\n", message)

	// \r is a carriage return
	message = "Hello\rGo!"
	fmt.Printf("%s\n", message)

	// raw string using backticks
	rawMessage := `Hello\nGo`

	fmt.Printf("%s\n", rawMessage)

	// string in ascii
	message = "Hello, World!"
	for i, char := range message {
		fmt.Printf("Index: %d, Char: %c, ASCII: %v, Hexadecimal: %x\n", i, char, char, char)
	}
	fmt.Println()

	// Runecount in string
	message = "Hello, World!"
	fmt.Printf("Number of runes in %s is %d\n", message, utf8.RuneCountInString(message))

	// simple rune
	var ch rune = 'a'

	// japanese character
	jch := '中'

	fmt.Printf("The rune %c has a value of: %v\n", ch, ch)
	fmt.Printf("The rune %c has a value of: %v\n", jch, jch)

	cstr := string(ch)
	cstr2 := string(jch)

	// placeholder for type
	fmt.Printf("The type of %s is %T\n", cstr, cstr)
	fmt.Printf("The type of %s is %T\n", cstr2, cstr2)

}
