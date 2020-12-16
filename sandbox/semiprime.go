package main

import (
    "fmt"
    "time"
)

// return true/false if n is a semiprime
var cache = make(map[int]bool)

func IsSemiprime(n int) bool {
    count := 0
    for i := 2; i <= n; i++ {
        for n % i == 0 {
            if count == 2 {
                return false
            }
            count++
            n /= i
        }
    }

    return count == 2
}

func main() {

    t0 := time.Now().UnixNano()
    for n := 1; n <= 10000; n++ {
        semi := IsSemiprime(n)
        if semi {
            cache[n] = true
            // fmt.Println(n, "->", semi)
        }
    }
    t1 := time.Now().UnixNano()
    fmt.Printf("elapsed %f seconds\n", float64(t1 - t0) / 1e9 )

    t0 = time.Now().UnixNano()
    for n := 1; n <= 30000; n++ {
        semi := cache[n]
        if semi {
            // cache[n] = true
            // fmt.Println(n, "->", semi)
        }
    }
    t1 = time.Now().UnixNano()
    fmt.Printf("elapsed %f seconds\n", float64(t1 - t0) / 1e9 )
}
