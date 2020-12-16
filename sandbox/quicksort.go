package main

import (
    "fmt"
    "sort"
    "math/rand"
)

func partition(a sort.Interface, first, last, pivotIndex int) int {
    a.Swap(first, pivotIndex)
    left := first+1
    right := last

    for left <= right {
        for left <= last && a.Less(left, first) {
            left++
        }
        for right >= first && a.Less(first, right) {
            right--
        }
        if left <= right {
            a.Swap(left, right)
            left++
            right--
        }
    }

    a.Swap(first, right)
    return right
}

func quicksortHelper(a sort.Interface, first, last int) {
    if first >= last {
        return
    }

    pivotIndex := partition(a, first, last, rand.Intn(last - first + 1) + first)
    quicksortHelper(a, first, pivotIndex-1)
    quicksortHelper(a, pivotIndex+1, last)
}

func quicksort(a sort.Interface) {
    quicksortHelper(a, 0, a.Len() - 1)
}

var (
    // test data
    A = []int{170, 45, 75, -90, -802, 24, 2, 66}
    B = []int{23,76,99,58,97,57,35,89,51,38,95,92,24,46,31,24,14,12,57,78,4}
    C = []int{88,18,31,44,4,0,8,81,14,78,20,76,84,33,73,75,82,5,62,70,12,7,1}
    list = [][]int{ A, B, C }
)

func main() {
    for _, a := range list {
        fmt.Println("  sort:", a)
        quicksort(sort.IntSlice(a))
        if sort.IntsAreSorted(a) {
            fmt.Println("sorted:", a)
        } else {
            fmt.Println("ERROR! not sorted!", a)
        }
    }
}


