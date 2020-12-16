/*
Two positive integers N and M are given. Integer N represents the number of chocolates arranged in a circle, numbered from 0 to N − 1.

You start to eat the chocolates. After eating a chocolate you leave only a wrapper.

You begin with eating chocolate number 0. Then you omit the next M − 1 chocolates or wrappers on the circle, and eat the following one.

More precisely, if you ate chocolate number X, then you will next eat the chocolate with number (X + M) modulo N (remainder of division).

You stop eating when you encounter an empty wrapper.

For example, given integers N = 10 and M = 4. You will eat the following chocolates: 0, 4, 8, 2, 6.

The goal is to count the number of chocolates that you will eat, following the above rules.

Write a function:

func Solution(N int, M int) int
that, given two positive integers N and M, returns the number of chocolates that you will eat.

For example, given integers N = 10 and M = 4. the function should return 5, as explained above.

Assume that:

N and M are integers within the range [1..1,000,000,000].
Complexity:

expected worst-case time complexity is O(log(N+M));
expected worst-case space complexity is O(log(N+M)).
*/

/*
NOTES: eaten := a / gcd(a,b)
*/

package main

import "fmt"

func gcd(a, b int) int {
	// binary euclidian algorithm
	var bgcd func(a, b, res int) int

	bgcd = func(a, b, res int) int {
		switch {
		case a == b:
			return res * a
		case a%2 == 0 && b%2 == 0:
			return bgcd(a/2, b/2, 2*res)
		case a%2 == 0:
			return bgcd(a/2, b, res)
		case b%2 == 0:
			return bgcd(a, b/2, res)
		case a > b:
			return bgcd(a-b, b, res)
		default:
			return bgcd(a, b-a, res)
		}
	}

	return bgcd(a, b, 1)
}

func Solution(a, b int) int {
	return a / gcd(a, b)
}

type TestData struct {
	a int
	b int
	R int
}

var testData []TestData = []TestData{
	TestData{10, 4, 5},
	TestData{3, 1, 3},
	TestData{4, 2, 2},
	TestData{5, 3, 5},
	TestData{20, 5, 4},
	TestData{1, 2, 1},
}

func main() {

	for _, data := range testData {
		// fmt.Println(data)

		r := Solution(data.a, data.b)
		if r == data.R {
			fmt.Printf("%v  %d pass\n", data, r)
		} else {
			fmt.Printf("%v %d FAIL!\n", data, r)
		}
	}
}
