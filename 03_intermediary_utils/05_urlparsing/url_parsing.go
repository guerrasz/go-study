package main

import (
	"fmt"
	"net/url"
)

func main() {

	rawURL := "https://example.com:8080/path?query=param#fragment"

	parsedURL, err := url.Parse(rawURL)

	if err != nil {
		fmt.Printf("Error parsing URL: %v\n", err)
		return
	}

	// prints url info
	fmt.Printf("Scheme: %v\n", parsedURL.Scheme)
	fmt.Printf("Host: %v\n", parsedURL.Host)
	fmt.Printf("Port: %v\n", parsedURL.Port())
	fmt.Printf("Path: %v\n", parsedURL.Path)
	fmt.Printf("Fragment: %v\n", parsedURL.Fragment)
	fmt.Printf("RawQuery: %v\n", parsedURL.RawQuery)

	fmt.Println()

	rawURL = "https://example.com:8080/path?name=Lucas&age=22"

	parsedURL, err = url.Parse(rawURL)

	if err != nil {
		fmt.Printf("Error parsing URL: %v\n", err)
		return
	}

	// prints info
	fmt.Printf("Scheme: %v\n", parsedURL.Scheme)
	fmt.Printf("Host: %v\n", parsedURL.Host)
	fmt.Printf("Port: %v\n", parsedURL.Port())
	fmt.Printf("Path: %v\n", parsedURL.Path)
	fmt.Printf("RawQuery: %v\n", parsedURL.RawQuery)

	// returns a map with the params
	queryParams := parsedURL.Query()
	fmt.Printf("Query: %v\n", queryParams)
	fmt.Printf("Name: %v\n", queryParams.Get("name"))
	fmt.Printf("Age: %v\n", queryParams.Get("age"))

	// building url
	baseURL := &url.URL{
		Scheme: "https",
		Host:   "example.com",
		Path:   "/path",
	}

	query := baseURL.Query()
	query.Set("name", "Lucas")
	query.Set("age", "22")
	baseURL.RawQuery = query.Encode()

	fmt.Printf("URL: %v\n", baseURL.String())
}
