package main

// TODO : sqrt 5 method
import (
	"fmt"
	"math"
    "time"
)

func fibo(n int) []int {
	r := make([]int, n)
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
		r[i] = x
	}

	return r
}

func fibchan(n int) []int {
	var fib = func(ch chan<- int) {
		a, b := 0, 1
		for {
			a, b = b, a+b
			ch <- a
		}
	}

	r := make([]int, n)
	fchan := make(chan int)
	go fib(fchan)
	for i := 0; i < n; i++ {
		r[i] = <-fchan
	}

	return r
}

// TODO : fix this
func fiboSqrt5(n int) []int {

	sqrt5 := math.Sqrt(float64(5))

	p := (1 + sqrt5) / 2
	q := 1 / p

	r := make([]int, n)

    for i := 0; i < n; i++ {
        x := float64(i + 1)
	    fib := int((math.Pow(p, x) + math.Pow(q, x)) / sqrt5 + 0.5)
        r[i] = fib
    }

	return r
}

type fiboCalculator func(int) []int

func profile(fn fiboCalculator, n int) ([]int, float64) {
    t0 := time.Now().UnixNano()
    r := fn(n)
    t1 := time.Now().UnixNano()

    return r, float64(t1 - t0) / 1e6
}

func main() {
    n := 35
	fmt.Println(profile(fibo, n))
	fmt.Println(profile(fibchan, n))
    fmt.Println(profile(fiboSqrt5, n))
}

