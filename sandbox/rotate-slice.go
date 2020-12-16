package main

import "fmt"

// rotate in place
func rotate(list []int) {
	n := len(list) - 1
	first := list[0]
	for i := 0; i < n; i++ {
		list[i] = list[i+1]
	}

	list[n] = first
}

func main() {
	list := []int{1, 2, 3, 4, 5, 6, 7}

	fmt.Println(list)
	rotate(list)
	fmt.Println(list)
	rotate(list)
	fmt.Println(list)
	rotate(list)
	fmt.Println(list)
}
