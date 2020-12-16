package main

import (
	"crypto/rand"
	"errors"

	"fmt"
	"golang.org/x/crypto/nacl/box"
	"io"
)

const (
	KeySize   = 32
	NonceSize = 24
)

var (
	ErrEncrypt = errors.New("secret encryption failed...")
	ErrDecrypt = errors.New("secret decryption failed...")
)

type Message struct {
	number  int
	message []byte
}

// encrypt a message
func GenerateNonce() (*[NonceSize]byte, error) {
	nonce := new([NonceSize]byte)
	_, err := io.ReadFull(rand.Reader, nonce[:])

	if err != nil {
		return nil, err
	}

	return nonce, nil
}

func Encrypt(pub, priv *[KeySize]byte, message []byte) ([]byte, error) {
	nonce, err := GenerateNonce()
	if err != nil {
		panic(err)
	}

	out := make([]byte, NonceSize)
	copy(out, nonce[:])

	fmt.Println("nonce: ", nonce)

	out = box.Seal(out, message, nonce, pub, priv)

	return out, nil
}

func Decrypt(pub, priv *[KeySize]byte, message []byte) ([]byte, error) {
	if len(message) < (box.Overhead + NonceSize) {
		return nil, ErrDecrypt
	}

	var nonce [NonceSize]byte
	copy(nonce[:], message[:NonceSize])

	fmt.Println("nonce: ", nonce)

	out, ok := box.Open(nil, message[NonceSize:], &nonce, pub, priv)
	if !ok {
		return nil, ErrDecrypt
	}

	return out, nil
}

func main() {
	pub, priv, _ := box.GenerateKey(rand.Reader)

	fmt.Printf("pub: %v, size: %d\n", pub, len(pub))
	fmt.Printf("priv: %v, size: %d\n", priv, len(priv))

	message := []byte("this is a message to be encrypted using box seal pub/priv keys")

	enc, err := Encrypt(pub, priv, message)

	if err != nil {
		panic(err)
	}

	fmt.Printf("enc: %v\n", enc)

	dec, err := Decrypt(pub, priv, enc)

	if err != nil {
		panic(err)
	}

	fmt.Printf("dec: %s\n", dec)
}
