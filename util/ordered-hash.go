package main

import (
	"fmt"
)

func main() {
	hash := make(map[string]int)

	hash["first"] = 1
	hash["third"] = 3
	hash["second"] = 2

	var text [3]string

	for k, v := range hash {
		idx := v - 1

		text[idx] = k
	}

	fmt.Println(text)

}
