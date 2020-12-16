/**
 * isprime
 *
 * @author darryl.west <darwest@ebay.com>
 * @created 2017-03-02 13:55:09
 */

package main

import (
	"fmt"
	"math"
)

func isprime(n int) bool {
	switch {
	case n == 2:
		return true
	case n < 2 || n%2 == 0:
		return false
	default:
		top := int(math.Sqrt(float64(n)))
		for i := 3; i <= top; i += 2 {
			if n%i == 0 {
				return false
			}
		}
	}

	return true
}

func main() {
	count := 0
	for n := 1; n < 110; n++ {
		p := isprime(n)
		if p {
			fmt.Printf("%d prime: %t\n", n, p)
			count++
		} else {
			fmt.Printf("%d\n", n)
		}
	}
	fmt.Println(count, "primes")
}
