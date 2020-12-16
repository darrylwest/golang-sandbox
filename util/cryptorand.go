package main

// see : https://elithrar.github.io/article/generating-secure-random-numbers-crypto-rand/

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)

	if err != nil {
		return nil, err
	}

	return b, nil
}

func GenerateRandomString(n int) (string, error) {
	b, err := GenerateRandomBytes(n)

	return base64.URLEncoding.EncodeToString(b), err
}

func GenerateRandomHex(n int) (string, error) {
	b, err := GenerateRandomBytes(n)

	return hex.EncodeToString(b), err
}

func main() {
	var size = 32 // 256
	token, err := GenerateRandomString(size)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("random token based on size: %d = %s (base64)\n", size, token)
	fmt.Printf("small token: %s\n", token[1:19])

	token, err = GenerateRandomHex(size)
	fmt.Printf("random hex: %s (hex), size: %d\n", token, len(token))

	decoded, err := hex.DecodeString(token)
	fmt.Printf("decoded: %v, size: %d\n", decoded, len(decoded))
}
