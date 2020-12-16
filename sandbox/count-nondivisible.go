/*
You are given a non-empty zero-indexed array A consisting of N integers.

For each number A[i] such that 0 ≤ i < N, we want to count the number of elements of the array that are not the divisors of A[i]. We say that these elements are non-divisors.

For example, consider integer N = 5 and array A such that:

    A[0] = 3
    A[1] = 1
    A[2] = 2
    A[3] = 3
    A[4] = 6
For the following elements:

A[0] = 3, the non-divisors are: 2, 6,
A[1] = 1, the non-divisors are: 3, 2, 3, 6,
A[2] = 2, the non-divisors are: 3, 3, 6,
A[3] = 3, the non-divisors are: 2, 6,
A[4] = 6, there aren't any non-divisors.
Write a function:

func Solution(A []int) []int
that, given a non-empty zero-indexed array A consisting of N integers, returns a sequence of integers representing the amount of non-divisors.

The sequence should be returned as:

a structure Results (in C), or
a vector of integers (in C++), or
a record Results (in Pascal), or
an array of integers (in any other programming language).
For example, given:

    A[0] = 3
    A[1] = 1
    A[2] = 2
    A[3] = 3
    A[4] = 6
the function should return [2, 4, 3, 2, 0], as explained above.

Assume that:

N is an integer within the range [1..50,000];
each element of array A is an integer within the range [1..2 * N].
Complexity:

expected worst-case time complexity is O(N*log(N));
expected worst-case space complexity is O(N), beyond input storage (not counting the storage required for input arguments).
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

func imax(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Solution(a []int) []int {

	// compute the frequency of each element in array a and get the max
	freq := make(map[int]int)
	maxa := 0
	for _, n := range a {
		freq[n] = freq[n] + 1
		maxa = imax(n, maxa)
	}
	fmt.Println(freq, maxa)

	// create a map of divisors for each unique element
	divisors := make(map[int][]int)
	for k, _ := range freq {
		list := make([]int, 1)
		// every number is divisible by 1
		list[0] = 1
		divisors[k] = list
	}

	// determine the divisors less than sqrt(maxn)
	sqrta := int(math.Sqrt(float64(maxa)))
	for divisor := 2; divisor <= sqrta; divisor++ {
		for multiple := divisor; multiple <= maxa; multiple += divisor {
			if v, ok := divisors[multiple]; ok {
				divisors[multiple] = append(v, divisor)
			}
		}
	}

	// complete all divisors greater than sqrta and filter out dups
	for k, v := range divisors {
		fmt.Println(k, v)
	}

	lena := len(a)
	results := make([]int, lena)
	for i, n := range a {

		r := 0
		results[i] = r
	}

	return results
}

const jsonData = `
    { "A":[ 3, 1, 2, 3, 6 ], "R":[ 2, 4, 3, 2, 0 ] }
`

type TestData struct {
	A []int
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

		r := Solution(data.A)
		if equal(r, data.R) {
			fmt.Printf("%v %v pass\n", data.R, r)
		} else {
			fmt.Printf("%v %v FAIL!\n", data.R, r)
		}
	}

}
