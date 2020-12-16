package main

import (
	"fmt"
	"math"
)

func calcRow(n int) []int {
	// row always has n + 1 elements
	row := make([]int, n+1, n+1)

	// set the edges
	row[0], row[n] = 1, 1

	// calculate values for the next n-1 columns
	for i := 0; i < int(n/2); i++ {
		x := int(row[i] * (n - i) / (i + 1))

		row[i+1], row[n-1-i] = x, x
	}

	return row
}

func calcSum(row []int) int {
	var sum int
	for _, v := range row {
		sum += v
	}

	return sum
}

func main() {

	for n := 0; n < 25; n++ {
		row := calcRow(n)
		sum := calcSum(row)
		pow := int(math.Pow(2, float64(n)))
		fmt.Printf("n = %d, row = %v, sum: = %d, pow: %v %v\n", n, row, sum, pow, (sum == pow))
	}

}
