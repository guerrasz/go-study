package main

import (
	// _ is used here to indicate that the package is used only for its side effects
	"embed"
	"fmt"
	"io/fs"
)

//go:embed example.txt
var content string

//go:embed folder
var folder embed.FS

func main() {
	fmt.Printf("Content: %s\n", content)

	content, err := folder.ReadFile("folder/foo.txt")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Content: %s\n", content)

	fs.WalkDir(folder, "folder", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		fmt.Printf("Path: %s\n", path)
		fmt.Printf("Directory Entry: %v\n", d)
		return nil
	})
}
