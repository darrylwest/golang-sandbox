package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 100; i++ {
		s := fmt.Sprintf("rand int: %x%x", r.Intn(9e7)+1e8, r.Intn(9e6)+1e7)
		fmt.Printf("%s len: %d\n", s, len(s))
	}
}
