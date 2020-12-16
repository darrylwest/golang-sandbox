package main

/**
 * from lesson 9, MaxProfit
 * see: https://en.wikipedia.org/wiki/Maximum_subarray_problem for Kadane's algorithm
 *
 */
import (
	"fmt"
)

func imax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func imin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// return the max profit or sell - buy
func kadane(list []int) int {
	profit := 0
	if len(list) < 1 {
		return profit
	}

	minbuy := list[0]
	for i := 1; i < len(list); i++ {
		profit = imax(profit, list[i]-minbuy)
		minbuy = imin(minbuy, list[i])
	}

	return profit
}

func Solution(A []int) int {
	return kadane(A)
}

type TestData struct {
	A []int
	R int
}

var testData []TestData = []TestData{
	TestData{A: []int{23171, 21011, 21123, 21366, 21013, 21367}, R: 356},
	TestData{A: []int{10, 7, 5, 8, 11, 9}, R: 6},
	TestData{A: []int{}, R: 0},
}

func main() {
	// use the kadane max subarray algorithm to solve...
	fmt.Println("from a stream of stock prices find the best buy/sell combination")

	tests := 0
	passed := 0

	for i := 0; i < len(testData); i++ {
		tests++
		A, R := testData[i].A, testData[i].R

		var passfail string

		r := Solution(A)
		if r == R {
			passed++
			passfail = "Ok"
		} else {
			passfail = "FAILED!"
		}

		fmt.Printf("%v = %d %s\n", A, r, passfail)
	}

	fmt.Printf("\nRan %d tests with %d passing, %d failing\n", tests, passed, tests-passed)
}
