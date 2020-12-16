package main

import (
	"fmt"
)

// standard recursion (stack problems)
func recur(n int) int {
	if n < 1 {
		return n
	}

	fmt.Println(n, n)

	return n + recur(n-1)
}

// tail recursion (still no optimization)
func tail(n int, accumulator int) int {
	accumulator += n
	fmt.Println(n, accumulator)

	if n < 1 {
		return accumulator
	}

	return tail(n-1, accumulator)
}

// a standard loop
func loop(n int) int {
	acc := 0
	for i := n; i > 0; i-- {
		acc += i
		fmt.Println(i, acc)
	}

	return acc
}

func main() {
	answer := recur(5)
	fmt.Printf("recur: %d\n", answer)

	answer = tail(5, 0)
	fmt.Printf("tail: %d\n", answer)

	loop(5)
	fmt.Printf("loop: %d\n", answer)

}
