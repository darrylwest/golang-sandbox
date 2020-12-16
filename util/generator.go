package main

import "fmt"

// a simple finite integer generator where n = from..to

func Generator(from, to int, out chan<- int) {

	for n := from; n < to; n++ {
		out <- n
	}

	close(out)
}

func main() {

	ch := make(chan int)
	go Generator(0, 10, ch)

	for {
		n, ok := <-ch
		if !ok {
			return
		}

		fmt.Printf("%d\n", n)
	}
}
