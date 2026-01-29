package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	testFile, err := os.Create("test.txt")

	if err != nil {
		fmt.Printf("Error creating file: %v", err)
		return
	}
	// defer file closing
	defer testFile.Close()

	// write data to file
	data := []byte("Hello , World!\n")
	_, err = testFile.Write(data)

	if err != nil {
		fmt.Printf("Error writing to file: %v", err)
		return
	}

	fmt.Printf("Data written to file successfully\n")

	test2File, err := os.Create("test2.txt")

	if err != nil {
		fmt.Printf("Error creating file: %v", err)
		return
	}
	defer test2File.Close()

	_, err = test2File.WriteString("Hello Go!\nI Am\nLucas Guerra\nLast line!\n")

	if err != nil {
		fmt.Printf("Error writing to file: %v", err)
		return
	}

	fmt.Printf("Data written to file successfully\n")

	openedFile, err := os.Open("test.txt")

	if err != nil {
		fmt.Printf("Error opening file: %v", err)
		return
	}
	defer openedFile.Close()

	// read file contents
	fileContent := make([]byte, 1024)

	_, err = openedFile.Read(fileContent)

	if err != nil {
		fmt.Printf("Error reading file: %v", err)
		return
	}

	fmt.Printf("File contents: %s\n", fileContent)

	// read second file line-by-line
	openedFile2, err := os.Open("test2.txt")

	if err != nil {
		fmt.Printf("Error opening file: %v", err)
		return
	}
	defer openedFile2.Close()

	scanner := bufio.NewScanner(openedFile2)

	// line by line
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("Line: %s\n", line)
	}

	err = scanner.Err()
	if err != nil {
		fmt.Printf("Error reading file: %v", err)
		return
	}

}
