package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	// base data
	data := []byte("He~lo base64")

	// encoded data
	encoded := base64.StdEncoding.EncodeToString(data)

	fmt.Printf("Encoded: %v\n", encoded)

	// decode data
	decoded, err := base64.StdEncoding.DecodeString(encoded)

	if err != nil {
		fmt.Printf("Error decoding: %v\n", err)
		return
	}

	fmt.Printf("Decoded: %v\n", string(decoded))

	urlSafeEncoded := base64.URLEncoding.EncodeToString(data)

	fmt.Printf("URL Safe Encoded: %v\n", urlSafeEncoded)

	urlSafeDecoded, err := base64.URLEncoding.DecodeString(urlSafeEncoded)

	if err != nil {
		fmt.Printf("Error decoding: %v\n", err)
		return
	}

	fmt.Printf("URL Safe Decoded: %v\n", string(urlSafeDecoded))
}
