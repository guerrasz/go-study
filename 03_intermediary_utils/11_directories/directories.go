package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	checkError(os.Mkdir("foo", 0755))
	defer deleteFiles()
	// create bar.txt in foo directory. with hello world string and 0644 permissions
	checkError(os.WriteFile("foo/bar.txt", []byte("Hello, World!"), 0644))
	// does not error if parent already exists
	checkError(os.MkdirAll("foo/buzz", 0755))
	checkError(os.MkdirAll("foo/buzz/bar", 0755))

	entries, err := os.ReadDir("foo")
	checkError(err)

	for i, entry := range entries {
		fmt.Printf("Entry %d: %s\n", (i + 1), entry)
	}

	// change dir to foo
	checkError(os.Chdir("./foo/"))
	currentDir, err := os.Getwd()
	checkError(err)
	fmt.Printf("Current directory: %s\n", currentDir)

	// iterate again but now in foo directory
	entries, err = os.ReadDir(".")
	checkError(err)

	for i, entry := range entries {
		fmt.Printf("Entry %d: %s\n", (i + 1), entry)
	}

	// change dir to parent directory
	checkError(os.Chdir(".."))
	currentDir, err = os.Getwd()
	checkError(err)
	fmt.Printf("Current directory: %s\n", currentDir)

	// filepath.WalkDir
	filepath.WalkDir("foo", func(path string, d os.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("Error: %v", err)
			return err
		}
		fmt.Printf("Path: %s\n", path)
		fmt.Printf("Directory Entry: %v\n", d)
		return nil
	})
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func deleteFiles() {
	//time.Sleep(2 * time.Second)
	checkError(os.RemoveAll("foo"))
}
