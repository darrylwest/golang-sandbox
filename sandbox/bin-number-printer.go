package main

import "fmt"

type bin int

func (b bin) String() string {
	return fmt.Sprintf("%b", b)
}

func main() {
	list := []int{1041, 244, 42, 2015, 2016}

	for i := 0; i < len(list); i++ {
		n := list[i]
		fmt.Println(i+1, n, bin(n))
	}
}
