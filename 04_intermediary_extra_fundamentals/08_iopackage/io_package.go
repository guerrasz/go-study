package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

// readFromReader reads from a reader and prints the data
func readFromReader(r io.Reader) {
	buf := make([]byte, 1024)
	n, err := r.Read(buf)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", string(buf[:n]))
}

// writeToWriter writes data to a writer
func writeToWriter(w io.Writer, data string) {
	_, err := w.Write([]byte(data))
	if err != nil {
		panic(err)
	}
}

// closeResource closes a closer
func closeResource(c io.Closer) {
	err := c.Close()
	if err != nil {
		panic(err)
	}
}

// bufferExamples shows examples of using bytes.Buffer
func bufferExamples() {
	var buf bytes.Buffer
	buf.WriteString("Hello World")
	fmt.Printf("%v\n", buf.String())
}

func multiReaderExample() {
	r1 := strings.NewReader("Hello world")
	r2 := strings.NewReader("Hello world")

	multiReader := io.MultiReader(r1, r2)

	buf := new(bytes.Buffer)

	_, err := buf.ReadFrom(multiReader)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", buf.String())
}

func pipeExamples() {
	pr, pw := io.Pipe()

	go func() {
		pw.Write([]byte("Hello pipe"))
		pw.Close()
	}()

	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(pr)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", buf.String())
}

func writeToFile(filePath string, data string) {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer closeResource(file)

	//& one possible way to write to a file
	// _, err = file.WriteString(data)
	// if err != nil {
	// 	panic(err)
	// }

	//& other possible way
	// writer := io.Writer(file)
	// _, err = writer.Write([]byte(data))
	// if err != nil {
	// 	panic(err)
	// }

	writer := bufio.NewWriter(file)
	_, err = writer.Write([]byte(data))
	if err != nil {
		panic(err)
	}

	err = writer.Flush()
	if err != nil {
		panic(err)
	}

}

func main() {
	fmt.Printf("=== Buffer Examples ===\n")
	bufferExamples()
	fmt.Println()

	fmt.Printf("=== Multi Reader Example ===\n")
	multiReaderExample()
	fmt.Println()

	fmt.Printf("=== Read From Reader Example ===\n")
	readFromReader(strings.NewReader("Hello Reader"))
	fmt.Println()

	fmt.Printf("=== Write To Writer Example ===\n")
	var writer bytes.Buffer

	writeToWriter(&writer, "Hello Writer")

	fmt.Printf("%v\n", writer.String())
	fmt.Println()

	fmt.Printf("=== Pipe Example ===\n")
	pipeExamples()
	fmt.Println()

	fmt.Printf("=== Write To File Example ===\n")
	writeToFile("test.txt", "Hello File\n")
	fmt.Println()

	fmt.Printf("=== MyResource Example ===\n")
	resource := MyResource{name: "MyResource"}
	defer closeResource(resource)
}

type MyResource struct {
	name string
}

func (m MyResource) Close() error {
	fmt.Printf("Closing %s\n", m.name)
	return nil
}
