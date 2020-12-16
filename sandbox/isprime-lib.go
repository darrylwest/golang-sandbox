package main

import (
	"fmt"
	"math/big"
)

func main() {
	count := 0

	step := big.NewInt(1)
	limit := big.NewInt(110)

	for n := big.NewInt(1); n.Cmp(limit) < 0; n = n.Add(n, step) {
		p := n.ProbablyPrime(20)
		if p {
			fmt.Printf("%d prime: %t\n", n, p)
			count++
		} else {
			fmt.Printf("%d\n", n)
		}
	}
	fmt.Println(count, "primes")
}
