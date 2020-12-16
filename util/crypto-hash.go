package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	username, pw := "darryl@bluelasso.com", "secret!"
	unpw := username + "!" + pw

	fmt.Printf("username: %s, pw: %s, unpw: %s\n", username, pw, unpw)

	hasher := sha256.New()
	hasher.Write([]byte(unpw))
	key := hex.EncodeToString(hasher.Sum(nil))

	fmt.Println(key)
	fmt.Printf("key length: %d\n", len(key))

	if key != "862948cf6aacc16c11e652c742cd78b31cce7dd2290be9d538e96b9ed75d1147" {
		fmt.Println("ERROR: key should match '862948cf6aacc16c11e652c742cd78b31cce7dd2290be9d538e96b9ed75d1147'")
	}
}
