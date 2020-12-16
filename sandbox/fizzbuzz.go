package main

import (
    "fmt"
)

func FizzBuzz(n int) []string {
    r := make([]string, n)

    for i := 1; i <= n; i++ {
        s := fmt.Sprintf("%d", i)

        switch {
        case i % 3 == 0 && i % 5 == 0:
            s += "FizzBuzz"
        case i % 3 == 0:
            s += "Fizz"
        case i % 5 == 0:
            s += "Buzz"
        }

        r[i-1] = s
    }

    return r
}

func main() {
    r := FizzBuzz(100)
    fmt.Println(r)
}
