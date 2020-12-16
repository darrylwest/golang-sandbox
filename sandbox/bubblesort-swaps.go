package main

import (
    "fmt"
)

var (
    // test data
    A = []int{170, 45, 75, -90, -802, 24, 2, 66}
    B = []int{23,76,99,58,97,57,35,89,51,38,95,92,24,46,31,24,14,12,57,78,4}
    C = []int{88,18,31,44,4,0,8,81,14,78,20,76,84,33,73,75,82,5,62,70,12,7,1}
    D = []int{1, 2, 3}
    E = []int{3, 2, 1}
    list = [][]int{ A, B, C, D, E }
)

func isSorted(a []int) bool {
    for i := 1; i < len(a); i++ {
        if a[i] < a[i-1] {
            return false
        }
    }
    return true
}

type stats struct {
    swaps int
    first int
    last int
}

func (s stats)show() {
    fmt.Printf("Array is sorted in %d swaps.\n", s.swaps)
    fmt.Println("First Element:", s.first)
    fmt.Println("Last Element:", s.last)
}

func bubbleSort(a []int) stats {
    last := len(a)
    stats := stats{}

    for {
        swaps := 0;

        for i := 1; i < last; i++ {
            if a[i] < a[i-1] {
                // swap
                a[i], a[i-1] = a[i-1], a[i]
                swaps++
            }
        }

        if swaps == 0 {
            stats.first = a[0]
            stats.last = a[len(a)-1]

            return stats
        }

        stats.swaps += swaps

        last--
    }
}

func main() {
    for _, a := range list {
        fmt.Println("  sort:", a)
        stats := bubbleSort(a)

        if isSorted(a) {
            fmt.Println("sorted:", a)
        } else {
            fmt.Println("not sorted!", a)
        }

        stats.show()
    }
}


