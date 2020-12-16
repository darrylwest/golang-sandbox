package main

import (
	"fmt"
)

func signum(x int) int {
	switch {
	case x > 0:
		return 1
	default:
		return 0
	case x < 0:
		return -1
	}
}
func main() {
	for i := -5; i < 6; i++ {
		fmt.Printf("i=%d, return = %d\n", i, signum(i))
	}
}
