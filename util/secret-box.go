package main

import (
	"crypto/rand"
	"errors"
	"fmt"
	"golang.org/x/crypto/nacl/secretbox"
	"io"
	"io/ioutil"
)

const (
	KeySize   = 32
	NonceSize = 24
)

var (
	ErrEncrypt = errors.New("secret encryption failed...")
	ErrDecrypt = errors.New("secret decryption failed...")
)

func GenerateKey() (*[KeySize]byte, error) {
	key := new([KeySize]byte)
	_, err := io.ReadFull(rand.Reader, key[:])

	if err != nil {
		return nil, err
	}

	return key, nil
}

func GenerateNonce() (*[NonceSize]byte, error) {
	nonce := new([NonceSize]byte)
	_, err := io.ReadFull(rand.Reader, nonce[:])

	if err != nil {
		return nil, err
	}

	return nonce, nil
}

func Encrypt(key *[KeySize]byte, message []byte) ([]byte, error) {
	nonce, err := GenerateNonce()
	if err != nil {
		return nil, ErrEncrypt
	}

	out := make([]byte, len(nonce))
	copy(out, nonce[:])
	out = secretbox.Seal(out, message, nonce, key)

	return out, nil
}

func Decrypt(key *[KeySize]byte, message []byte) ([]byte, error) {
	if len(message) < (NonceSize + secretbox.Overhead) {
		return nil, ErrDecrypt
	}

	var nonce [NonceSize]byte
	copy(nonce[:], message[:NonceSize])
	out, ok := secretbox.Open(nil, message[NonceSize:], &nonce, key)
	if !ok {
		return nil, ErrDecrypt
	}

	return out, nil
}

func main() {
	key, _ := GenerateKey()
	nonce, _ := GenerateNonce()

	fmt.Printf("key: %v, nonce: %v\n", key, nonce)

	message, err := ioutil.ReadFile("secret.tgz")
	if err != nil {
		panic(err)
	}

	enc, _ := Encrypt(key, message)

	fmt.Printf("enc: %v\n", enc)
	fmt.Printf("encrypted size: %d\n", len(enc))

	dec, _ := Decrypt(key, enc)
	fmt.Printf("dec size: %d\n", len(dec))

	err = ioutil.WriteFile("secret-dec.tgz", dec, 0777)
	if err != nil {
		fmt.Printf("file write error: ", err)
	}
}
