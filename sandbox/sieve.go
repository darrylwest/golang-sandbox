package main

import (
	"fmt"
	"math"
)

func sieve(n int) ([]bool, []int) {
	primes := []int{}
	s := make([]bool, n+1)

	// set them initially to true
	for i := 2; i <= n; i++ {
		s[i] = true
	}

	// 1, 2 are primes
	s[0], s[1] = false, false

	for i := 2; i*i <= n; i++ {
		if s[i] {
			primes = append(primes, i)
			for k := i * i; k <= n; k += i {
				s[k] = false
			}
		}
	}

	sqrt_n := int(math.Sqrt(float64(n)))
	for i := sqrt_n + 1; i <= n; i++ {
		if s[i] {
			primes = append(primes, i)
		}
	}

	return s, primes

}

func main() {
	n := 37
	result, primes := sieve(n)

	for i := 0; i < len(result); i++ {
		if result[i] {
			fmt.Println(i, "prime")
		} else {
			fmt.Println(i, "")
		}

	}

	fmt.Println(primes)
	fmt.Println(len(primes))

}
