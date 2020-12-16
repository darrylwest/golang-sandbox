package main

import (
	"encoding/base32"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Println(rnd.Int(), rnd.Int())
	fmt.Println(rnd.Int63(), rnd.Int63())
	fmt.Println(rnd.Uint32(), rnd.Uint32())
	fmt.Println(rnd.Float64(), rnd.Float64())

	n := rnd.Int63()
	fmt.Printf("%d %X\n", n, n)

	data := []byte("this is a test")
	str := base32.StdEncoding.EncodeToString(data)
	fmt.Println(str)
}
