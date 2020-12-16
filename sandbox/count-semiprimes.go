/*
A prime is a positive integer X that has exactly two distinct divisors: 1 and X. The first few prime integers are 2, 3, 5, 7, 11 and 13.

A semiprime is a natural number that is the product of two (not necessarily distinct) prime numbers. The first few semiprimes are 4, 6, 9, 10, 14, 15, 21, 22, 25, 26.

You are given two non-empty zero-indexed arrays P and Q, each consisting of M integers. These arrays represent queries about the number of semiprimes within specified ranges.

Query K requires you to find the number of semiprimes within the range (P[K], Q[K]), where 1 ≤ P[K] ≤ Q[K] ≤ N.

For example, consider an integer N = 26 and arrays P, Q such that:

    P[0] = 1    Q[0] = 26
    P[1] = 4    Q[1] = 10
    P[2] = 16   Q[2] = 20
The number of semiprimes within each of these ranges is as follows:

(1, 26) is 10,
(4, 10) is 4,
(16, 20) is 0.
Write a function:

func Solution(N int, P []int, Q []int) []int
that, given an integer N and two non-empty zero-indexed arrays P and Q consisting of M integers, returns an array consisting of M elements specifying
the consecutive answers to all the queries.

For example, given an integer N = 26 and arrays P, Q such that:

    P[0] = 1    Q[0] = 26
    P[1] = 4    Q[1] = 10
    P[2] = 16   Q[2] = 20
the function should return the values [10, 4, 0], as explained above.

Assume that:

N is an integer within the range [1..50,000];
M is an integer within the range [1..30,000];
each element of arrays P, Q is an integer within the range [1..N];
P[i] ≤ Q[i].
Complexity:

expected worst-case time complexity is O(N*log(log(N))+M);
expected worst-case space complexity is O(N+M), beyond input storage (not counting the storage required for input arguments).
Elements of input arrays can be modified.
*/

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"strings"
)

// return a list of all the primes up to n
func sieve(n int) []int {
	sqrt_n := int(math.Sqrt(float64(n)))
	primes := []int{}
	s := make([]bool, n+1)

	// set them initially to true
	for i := 2; i <= n; i++ {
		s[i] = true
	}

	// 1, 2 are primes
	s[0], s[1] = false, false

	i := 2
	for i*i <= n {
		if s[i] {
			primes = append(primes, i)
			k := i * i
			for k <= n {
				s[k] = false
				k += i
			}
		}
		i++
	}

	for i := sqrt_n + 1; i <= n; i++ {
		if s[i] {
			primes = append(primes, i)
		}
	}

	return primes
}

func Solution(n int, p, q []int) []int {
	primes := sieve(n)
	pcount := len(primes)

	semiprimes := make([]int, n+1)

	// semiprimes[i] == 1 -> when i i semiprime, else 0

	for i := 0; i <= pcount-1; i++ {
		for j := i; j < pcount; j++ {
			k := primes[i] * primes[j]
			if k > n {
				break
			}

			semiprimes[k] = 1
		}
	}

	// count the number of semiprimus for each postion
	for i := 1; i <= n; i++ {
		semiprimes[i] += semiprimes[i-1]
	}

	// now get the results for each query range
	results := make([]int, len(p))
	for i := 0; i < len(p); i++ {
		results[i] = semiprimes[q[i]] - semiprimes[p[i]-1]
	}

	return results
}

const jsonData = `
    { "N":26, "P":[ 1, 4, 16 ], "Q":[ 26, 10, 20 ], "R":[ 10, 4, 0 ] }
`

type TestData struct {
	N int
	P []int
	Q []int
	R []int
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {
	decoder := json.NewDecoder(strings.NewReader(jsonData))
	for {
		var data TestData
		if err := decoder.Decode(&data); err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("ERROR! %v\n", err)
		}

		r := Solution(data.N, data.P, data.Q)
		if equal(r, data.R) {
			fmt.Printf("%v %v pass\n", data.R, r)
		} else {
			fmt.Printf("%v %v FAIL!\n", data.R, r)
		}
	}
}
