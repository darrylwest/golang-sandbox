package main

import (
	"fmt"
)

func reverse(slice []byte) {
	for left, right := 0, len(slice)-1; left < right; left, right = left+1, right-1 {
		slice[left], slice[right] = slice[right], slice[left]
	}
}

func main() {
	slice := []byte("this is a typical string")

	fmt.Printf("%s\n", slice)
	reverse(slice)
	fmt.Printf("%s\n", slice)
}
