package main

// very simple chacha20 symetric key example.  
// to make this usable the nonce(s) and keys would need to be safely stored
// nonce could be stored by prepending the ciphertext

import (
  cha "golang.org/x/crypto/chacha20poly1305"
  cryptorand "crypto/rand"
  "encoding/hex"
  "fmt"
  "log"
)


func main() {
  key := make([]byte, cha.KeySize)
  aead, err := cha.NewX(key)
  if err != nil {
    log.Fatalln("Failed to instantiate XChaCha20-Poly1305:", err)
  }
  fmt.Println("key...")
  fmt.Printf("%s", hex.Dump(aead))

  for _, msg := range []string{
    "Attack at dawn.",
    "The eagle has landed.",
    "Gophers erverywhere!",
  } {
    // encryption
    nonce := make([]byte, cha.NonceSizeX)
    if _, err := cryptorand.Read(nonce); err != nil {
        panic(err)
    }

    fmt.Println("nonce...")
    fmt.Printf("%s", hex.Dump(nonce))

    ciphertext := aead.Seal(nil, nonce, []byte(msg), nil)
    fmt.Println("cipher...")
    fmt.Printf("%s", hex.Dump(ciphertext))

    // decrypt
    plaintext, err := aead.Open(nil, nonce, ciphertext, nil)
    if err != nil {
        log.Fatalln("Failed to decrypt message:", err)
    }

    fmt.Println("plain text...")
    fmt.Printf("%s\n", plaintext)
    fmt.Println("...")
  }
}

