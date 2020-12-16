package main

import (
	"fmt"
	"time"
)

var runnerA = func() {
	fmt.Printf("now A: %v\n", time.Now())
}

var runnerB = func() {
	fmt.Printf("now B: %v\n", time.Now())
}

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main() {

	jobs := []func(){runnerA, runnerB}

	for _, fn := range jobs {
		fn()
	}

	fn := squares()
	for i := 0; i < 10; i++ {
		fmt.Println(i+1, fn())
	}
}
