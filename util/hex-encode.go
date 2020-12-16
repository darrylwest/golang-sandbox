package main

import (
	"code.google.com/p/go-uuid/uuid"
	"encoding/hex"
	"fmt"
	"strings"
)

func main() {
	session := uuid.New()
	id := strings.Replace(uuid.New(), "-", "", -1)

	fmt.Printf("session:%s, id: %s\n", session, id)

	bytes := []byte(fmt.Sprintf("%s|%s", session, id))

	fmt.Printf("bytes: %s\n", bytes)

	encoded := hex.EncodeToString(bytes)

	fmt.Printf("encoded: %s\n", encoded)

	decoded, err := hex.DecodeString(encoded)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("decoded: %s\n", decoded)
	}
}
