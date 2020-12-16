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

func isSorted(a []int) bool {
    for i := 1; i < len(a); i++ {
        if a[i] < a[i-1] {
            return false
        }
    }
    return true
}

func createScratchSpace(a []int) []int {
    return make([]int, len(a)/2+1)
}

func mergesort(a []int, scratch []int) {
    ln := len(a)
    if ln <= 1 {
        return
    }

    mid := ln >> 1
    mergesort(a[:mid], scratch) // left slice
    mergesort(a[mid:], scratch) // right slice

    if a[mid-1] <= a[mid] {
        return
    }

    copy(scratch, a[:mid])
    left, right := 0, mid
    for i := 0; ; i++ {
        if scratch[left] <= a[right] {
            a[i] = scratch[left]
            left++
            if left == mid {
                break
            }
        } else {
            a[i] = a[right]
            right++
            if right == ln {
                copy(a[i+1:], scratch[left:mid])
                break
            }
        }
    }

}

func main() {
    for _, a := range list {
        fmt.Println("  sort:", a)
        mergesort(a, createScratchSpace(a))
        if sort.IntsAreSorted(a) {
            fmt.Println("sorted:", a)
        } else {
            fmt.Println("ERROR! not sorted:", a)
        }
    }
}


