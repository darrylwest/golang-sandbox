package main

/*
A binary gap within a positive integer N is any maximal sequence of consecutive zeros that is surrounded by ones at both ends in the binary representation of N.

For example, number 9 has binary representation 1001 and contains a binary gap of length 2. The number 529 has binary representation 1000010001 and contains two
binary gaps: one of length 4 and one of length 3. The number 20 has binary representation 10100 and contains one binary gap of length 1. The number 15 has binary
representation 1111 and has no binary gaps.

Write a function:

func Solution(N int) int
that, given a positive integer N, returns the length of its longest binary gap. The function should return 0 if N doesn't contain a binary gap.

For example, given N = 1041 the function should return 5, because N has binary representation 10000010001 and so its longest binary gap is of length 5.

Assume that:

N is an integer within the range [1..2,147,483,647].
Complexity:

expected worst-case time complexity is O(log(N));
expected worst-case space complexity is O(1).
*/

import "fmt"

func imax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Solution(n int) int {
	s := []byte(fmt.Sprintf("%b", n))

	if len(s) < 3 {
		return 0
	}

	var zero byte = 48 // ascii value
	// trim the trailing zeros
	for i := len(s) - 1; i > 0; i-- {
		if s[i] == zero {
			s = s[:i]
		} else {
			break
		}
	}

	maxcount := 0
	count := 0

	for _, v := range s {
		if v == zero {
			count++
		} else {
			count = 0
		}

		maxcount = imax(maxcount, count)
	}

	return maxcount
}

func main() {
	n := 1041
	r := Solution(n)

	fmt.Println(n, fmt.Sprintf("%b", n), r)

	n = 2147483647
	r = Solution(n)

	fmt.Println(n, fmt.Sprintf("%b", n), r)

	n = 2015
	fmt.Println(n, fmt.Sprintf("%b", n), Solution(n))

	n = 6
	fmt.Println(n, fmt.Sprintf("%b", n), Solution(n))

	n = 328
	fmt.Println(n, fmt.Sprintf("%b", n), Solution(n))

	n = 5
	fmt.Println(n, fmt.Sprintf("%b", n), Solution(n))

	n = 16
	fmt.Println(n, fmt.Sprintf("%b", n), Solution(n))

	n = 1024
	fmt.Println(n, fmt.Sprintf("%b", n), Solution(n))
}
