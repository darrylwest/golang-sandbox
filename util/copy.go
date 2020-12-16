package main

import (
	"fmt"
)

func main() {
	sliceA := []int{1, 2, 3}
	sliceB := [3]int{}

	fmt.Println("before copy, a:", sliceA, "b:", sliceB)

	copy(sliceB[:], sliceA)
	fmt.Println("after copy, a:", sliceA, "b:", sliceB)

	hashA := map[string]bool{"A": true, "B": false}
	hashB := map[string]bool{}
	fmt.Println("before copy, a:", hashA, "b:", hashB)

	for k, v := range hashA {
		hashB[k] = v
	}

	fmt.Println("after copy, a:", hashA, "b:", hashB)

}
