package main

import (
	"fmt"
	"io"
	"os/exec"
)

func main() {
	pr, pw := io.Pipe()

	cmd := exec.Command("grep", "foo")

	cmd.Stdin = pr

	go func() {
		defer pw.Close()
		pw.Write([]byte("food is great\nI love bsd\nI love food"))
	}()

	output, err := cmd.Output()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Output: %s", string(output))
	
	

}
