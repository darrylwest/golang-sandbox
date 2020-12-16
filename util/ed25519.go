package main

import (
	"crypto/rand"
	"fmt"
	"github.com/agl/ed25519"
)

func main() {
	pub, priv, _ := ed25519.GenerateKey(rand.Reader)

	fmt.Println(pub)
	fmt.Println(priv)

	message := []byte("this is a plain text message the requires a signature")

	sig := ed25519.Sign(priv, message)

	if ed25519.Verify(pub, message, sig) {
		fmt.Println(sig)
	}
}
