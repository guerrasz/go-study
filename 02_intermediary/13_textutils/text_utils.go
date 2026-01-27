package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	str := "Hello, Go"

	fmt.Printf("%v\n", str[1:7])

	// utils
	num := 18
	fromNum := strconv.Itoa(num)
	fmt.Printf("%v\n", len(fromNum))

	// splitting
	fruits := "apple,banana,orange,grape"
	parts := strings.Split(fruits, ",")
	fmt.Printf("%v\n", parts)

	// joining
	newFruits := strings.Join(parts, " | ")
	fmt.Printf("%v\n", newFruits)

	// contains
	fmt.Printf("%v\n", strings.Contains(newFruits, "banana"))

	// replace
	fmt.Printf("%v\n", strings.Replace(newFruits, "banana", "kiwi", 1))

	// uppercase
	fmt.Printf("%v\n", strings.ToUpper(newFruits))

	// lowercase
	fmt.Printf("%v\n", strings.ToLower(newFruits))

	r := " " + newFruits
	fmt.Printf("%v\n", r)

	// trim
	fmt.Printf("%v\n", strings.TrimSpace(r))

	// repeat and count
	fmt.Printf("%v\n", strings.Repeat(newFruits, 3))
	fmt.Printf("%v\n", strings.Count(newFruits, "banana"))

	// regex
	find := "Hello 123 Go! 11"
	re := regexp.MustCompile(`\d+`)

	// -1 means find all
	fmt.Printf("%v\n", re.FindAllString(find, -1))

	// string builder
	var builder strings.Builder
	for range 10 {
		builder.WriteString("Hello ")
	}
	fmt.Printf("%v\n", builder.String())

}
