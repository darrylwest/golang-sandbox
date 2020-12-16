package main

import (
	"fmt"
)

// reverse in place
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

}

func main() {
	list := []int{1, 2, 3, 4}

	fmt.Println(list)
	reverse(list)
	fmt.Println(list)

}
