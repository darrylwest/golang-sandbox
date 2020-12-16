package main

import (
    "fmt"
    "sort"
)

var (
    // test data
    A = []int{170, 45, 75, -90, -802, 24, 2, 66}
    B = []int{23,76,99,58,97,57,35,89,51,38,95,92,24,46,31,24,14,12,57,78,4}
    C = []int{88,18,31,44,4,0,8,81,14,78,20,76,84,33,73,75,82,5,62,70,12,7,1}
    list = [][]int{ A, B, C }
)

func bubbleSort(a []int) {
    last := len(a)
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
            return
        }

        last--
    }
}

func main() {
    for _, a := range list {
        fmt.Println("  sort:", a)
        bubbleSort(a)
        if sort.IntsAreSorted(a) {
            fmt.Println("sorted:", a)
        } else {
            fmt.Println("not sorted!", a)
        }
    }
}


