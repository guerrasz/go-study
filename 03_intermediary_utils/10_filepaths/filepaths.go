package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	// join path
	fmt.Printf("Joined path: %v\n", filepath.Join("Downloads", "foo.zip"))

	// print normalized path
	fmt.Printf("Normalized path: %v\n", filepath.Clean("Downloads/./foo.zip"))

	// split path and file from filepath
	dir, file := filepath.Split("Downloads/foo.zip")
	fmt.Printf("Split path: %v\n", dir)
	fmt.Printf("Split file: %v\n", file)

	// get file extension
	fmt.Printf("File extension: %v\n", filepath.Ext("Downloads/foo.zip"))

	// get base name
	fmt.Printf("Base name: %v\n", filepath.Base("Downloads/foo.zip"))

	// is absolute path?
	fmt.Printf("Is absolute path: %v\n", filepath.IsAbs("Downloads/foo.zip"))

	// relative path
	relativePath, err := filepath.Rel("Downloads/foo.zip", "Downloads")
	if err != nil {
		fmt.Printf("Error getting relative path: %v\n", err)
		return
	}
	fmt.Printf("Relative path: %v\n", relativePath)

	// absolute path
	absolutePath, err := filepath.Abs("Downloads/foo.zip")
	if err != nil {
		fmt.Printf("Error getting absolute path: %v\n", err)
		return
	}
	fmt.Printf("Absolute path: %v\n", absolutePath)
}
