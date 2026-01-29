package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	sub1 := flag.NewFlagSet("sub1", flag.ExitOnError)
	sub2 := flag.NewFlagSet("sub2", flag.ExitOnError)

	sub1First := sub1.Bool("processing", false, "Processing status")
	sub1Second := sub1.Int("bytes", 1024, "Byte size")

	sub2First := sub2.String("language", "go", "Language")

	if len(os.Args) < 2 {
		fmt.Println("Please provide a subcommand")
		return
	}

	switch os.Args[1] {
	case "sub1":
		sub1.Parse(os.Args[2:])
		fmt.Printf("Sub1: %v\n", *sub1First)
		fmt.Printf("Sub1: %v\n", *sub1Second)
	case "sub2":
		sub2.Parse(os.Args[2:])
		fmt.Printf("Sub2: %v\n", *sub2First)
	default:
		fmt.Println("Unknown subcommand")
	}
}
