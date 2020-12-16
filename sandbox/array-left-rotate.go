package main

import (
    "fmt"
)

/*
    rotate an array to the left n times

    1 <= n <= 1e5
    1 <= d <= n -> so rotates is always less than length
    1 <= a <= 1e6 -> absolute values shouldn't matter in this case
*/

// brute force
func rotateLeftBrute(n int, A []int) []int {
    rotates := n // % len(A)

    for i := 0; i < rotates; i++ {
        A = append(A[1:], A[0])
    }

    return A
}

func rotateLeft(n int, A []int) []int {
    ln := len(A)
    result := make([]int, ln)
    j := n
    for i := 0; i < ln; i++ {
        result[i] = A[(i + j) % ln]
    }

    return result
}

func testit(rotateCount int, A, expect []int) bool {
    results := rotateLeft(rotateCount, A)
    maxp := 10

    if len(A) < 10 {
        fmt.Printf("input: %v results: %v ", A, results)
    } else {
        fmt.Printf("input: %v results: %v ", A[:maxp], results[:maxp])
    }

    errorCount := 0
    for i, n := range expect {
        value := results[i]
        if n != value {
            fmt.Printf("item %d expected %d got %d\n", i, n, value)
            errorCount++
        }
    }

    if errorCount == 0 {
        fmt.Println("ok...")
    } else {
        fmt.Printf("FAILED! %d errors...")
    }

    return errorCount == 0
}

func main() {
    A := []int{ 1, 2, 3, 4, 5 }
    n := 4
    expect := []int{ 5, 1, 2, 3, 4 }

    testit(n, A, expect)

    n = 1
    expect = []int{ 2, 3, 4, 5, 1 }
    testit(n, A, expect)

    n = 2
    expect = []int{ 3, 4, 5, 1, 2 }
    testit(n, A, expect)

    // zero rotations
    expect = []int{ 1, 2, 3, 4, 5 }
    testit(0, A, expect)

    // max array size rotated n / 2
    A = make([]int, 1e5)
    expect = make([]int, 1e5)
    n = len(A) / 2

    for i,_ := range A {
        value := (i * 10) + 1
        j := (n + i) % len(A)

        A[i], expect[j] = value, value
    }
    testit(n, A, expect)
}

