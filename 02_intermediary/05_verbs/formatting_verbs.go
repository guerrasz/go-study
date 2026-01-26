package main

import "fmt"

func main() {
	// General formatting verbs
	// %v default format
	// %#v Go-syntax representation
	// %T type of the value
	// %% literal percent sign

	num := 560_780
	message := "Hello World!"

	fmt.Printf("%v\n", num)
	fmt.Printf("%#v\n", num)
	fmt.Printf("%T\n", num)
	fmt.Printf("%v\n", message)
	fmt.Printf("%#v\n", message)
	fmt.Printf("%T\n", message)

	fmt.Println()

	// Integer formatting verbs
	// %b base 2
	// %d base 10
	// %+d base 10 and show sign
	// %o base 8
	// %x base 16
	// %X base 16 with uppercase

	num = 18

	fmt.Printf("%b\n", num)
	fmt.Printf("%d\n", num)
	fmt.Printf("%+d\n", num)
	fmt.Printf("%o\n", num)
	fmt.Printf("%x\n", num)
	fmt.Printf("%X\n", num)

	fmt.Println()

	// String formatting verbs
	// %s string
	// %q double-quoted string with Go syntax
	// %8s string with padding of 8 characters
	// %-8s string with padding of 8 characters left aligned
	// %x hex dump of raw bytes
	// %X hex dump of raw bytes with uppercase
	// % x hex dump of raw bytes with space padding

	fmt.Printf("%s\n", message)
	fmt.Printf("%q\n", message)
	fmt.Printf("%16s\n", message)
	fmt.Printf("%-16s\n", message)
	fmt.Printf("%x\n", message)
	fmt.Printf("%X\n", message)
	fmt.Printf("% x\n", message)

	fmt.Println()

	// Boolean formatting verbs
	// %t boolean

	t := true
	f := false

	fmt.Printf("%t\n", t)
	fmt.Printf("%t\n", f)

	// Float formatting verbs
	// %e scientific notation
	// %f decimal point
	// %.2f decimal point with 2 decimal places
	// %6.2f 6 characters wide with 2 decimal places
	// %g exponent as needed, only necessary digits

	floatNum := 918_000.78

	fmt.Printf("%e\n", floatNum)
	fmt.Printf("%f\n", floatNum)
	fmt.Printf("%.2f\n", floatNum)
	fmt.Printf("%6.2f\n", floatNum)
	fmt.Printf("%g\n", floatNum)

}
