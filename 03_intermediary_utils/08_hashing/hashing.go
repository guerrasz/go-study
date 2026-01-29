package main

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"io"
)

func main() {
	password := "password123"

	hash := sha256.Sum256([]byte(password))
	hash512 := sha512.Sum512([]byte(password))

	fmt.Printf("Password: %s\nHash in hex: %x\nHash in byte slice: %v\n", password, hash, hash)
	fmt.Println()
	fmt.Printf("Password: %s\nHash in hex: %x\nHash in byte slice: %v\n", password, hash512, hash512)

	salt, err := generateSalt()
	if err != nil {
		fmt.Printf("Error generating salt: %v\n", err)
		return
	}

	fmt.Printf("Salt (hex): %x\n", salt)
	fmt.Printf("Salt (string): %s\n", base64.StdEncoding.EncodeToString(salt))

	hashedPassword := hashPassword(password, salt)
	fmt.Printf("Hashed Password: %s\n", hashedPassword)

	// verify pass
	decodedSalt, err := base64.StdEncoding.DecodeString(base64.StdEncoding.EncodeToString(salt))
	if err != nil {
		fmt.Printf("Error decoding salt: %v\n", err)
		return
	}
	fmt.Printf("Decoded Salt: %x\n", decodedSalt)

	loginHash := hashPassword(password, decodedSalt)
	fmt.Printf("Login Hash: %s\n", loginHash)

	fmt.Printf("Password verified: %v\n", loginHash == hashedPassword)
}

func generateSalt() ([]byte, error) {
	// create 16 bytes salt
	salt := make([]byte, 16)

	// read random bytes into salt
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

func hashPassword(password string, salt []byte) string {
	saltedPassword := append(salt, []byte(password)...)
	hash := sha256.Sum256(saltedPassword)
	return base64.StdEncoding.EncodeToString(hash[:])
}
