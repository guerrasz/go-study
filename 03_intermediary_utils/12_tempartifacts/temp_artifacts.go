package main

import (
	"fmt"
	"os"
)

func main() {
	// create tmp directory
	tmpDir, err := os.MkdirTemp(".", "")
	checkError(err)

	// defer dir removal
	defer os.RemoveAll(tmpDir)

	fmt.Printf("Temporary directory created: %v\n", tmpDir)

	// create temp file inside tmp dir
	tmpFile, err := os.CreateTemp(tmpDir, "")
	checkError(err)

	// defer must be called in reverse order of creation
	defer tmpFile.Close()

	fmt.Printf("Temporary file created: %v\n", tmpFile.Name())
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
