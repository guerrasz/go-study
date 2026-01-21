package main

import (
	"fmt"
	"net/http"
)

// Import packages from STD Library
func main() {
	fmt.Println("Hello STD Library")

	// HTTP GET request
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")

	// This code is still a mystery and only serves as an example for the http package
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

}
