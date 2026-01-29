package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// creating a buffer on top of a reader
	reader := bufio.NewReader(strings.NewReader("Hello bufio package!\n"))

	// create blank byte slice
	data := make([]byte, 20)

	// read data into data
	n, err := reader.Read(data)

	if err != nil {
		fmt.Printf("Error reading %v\n", err)
		return
	}

	// here %s prints the slice of bytes as a string
	fmt.Printf("Bytes read %d, data: %s\n", n, data[:n]) // :n is to print only the bytes read

	// read until \n
	line, err := reader.ReadString('\n')

	if err != nil {
		fmt.Printf("Error reading %v\n", err)
		return
	}

	fmt.Printf("Line read: %s\n", line)

	writer := bufio.NewWriter(os.Stdout)

	// write byte slice
	data = []byte("Hello from buffer!\n")

	// writes to buffer
	n, err = writer.Write(data)

	if err != nil {
		fmt.Printf("Error writing %v\n", err)
		return
	}

	fmt.Printf("Bytes written %d\n", n)

	// flush the buffer to the target
	err = writer.Flush()

	if err != nil {
		fmt.Printf("Error flushing %v\n", err)
		return
	}

	// writing string
	str := "Ablueblueblue\n"
	n, err = writer.WriteString(str)

	if err != nil {
		fmt.Printf("Error writing %v\n", err)
		return
	}

	fmt.Printf("Bytes written %d\n", n)

	// flush the buffer to the target
	err = writer.Flush()

	if err != nil {
		fmt.Printf("Error flushing %v\n", err)
		return
	}
}
