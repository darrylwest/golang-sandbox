package main

import (
	"encoding/base64"
	"fmt"
	"github.com/pborman/uuid"
	"strings"
)

func main() {
	session := uuid.New()
	id := strings.Replace(uuid.New(), "-", "", -1)

	fmt.Printf("session:%s, id: %s\n", session, id)

	bytes := []byte(fmt.Sprintf("%s|%s", session, id))

	fmt.Printf("bytes: %s\n", bytes)

	encoded := base64.StdEncoding.EncodeToString(bytes)

	fmt.Printf("encoded: %s\n", encoded)

	decoded, err := base64.StdEncoding.DecodeString(encoded)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("decoded: %s\n", decoded)
	}
}
