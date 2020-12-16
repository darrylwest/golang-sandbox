package main

// TODO: modify to use a password with bcrypt or scrypt

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/hex"
    "io"
    "fmt"
)

func encrypt(data []byte, key []byte) []byte {
    block, _ := aes.NewCipher(key)
    gcm, _ := cipher.NewGCM(block)
    nonce := make([]byte, gcm.NonceSize())
    io.ReadFull(rand.Reader, nonce)
    ciphertext := gcm.Seal(nonce, nonce, data, nil)

    return ciphertext
}

func decrypt(data []byte, key []byte) []byte {
    block, _ := aes.NewCipher(key)
    gcm, _ := cipher.NewGCM(block)

    nonceSize := gcm.NonceSize()
    nonce, ciphertext := data[:nonceSize], data[nonceSize:]
    plaintext, _ := gcm.Open(nil, nonce, ciphertext, nil)

    return  plaintext
}

func generateKey(keysize int) []byte {
    key := make([]byte, keysize)
    io.ReadFull(rand.Reader, key)

    return key
}

func main() {
    key := generateKey(32)
    fmt.Printf("key: %s\n", hex.EncodeToString(key))

    ciphertext := encrypt([]byte("this is my plain text"), key)
    fmt.Println(ciphertext)

    plaintext := decrypt(ciphertext, key)
    fmt.Printf("%s\n", plaintext)
}
