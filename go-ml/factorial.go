package main

import (
    "fmt"
    "math/big"
    "flag"
)

const LIMIT = 100

func factorial(n int64) *big.Int {
    one := big.NewInt(1)
    if n < 2 {
        return one
    }

    result := big.NewInt(n)
    for ; n > 1; n-- {
        result = result.Mul(result, big.NewInt(n-1))
    }

    return result
}

func main() {
    nptr := flag.Int("n", 0, "number of factorials to calc")
    qptr := flag.Bool("q", false, "quiet...")

    flag.Parse()

    if *nptr > 0 {
        a := factorial(int64(*nptr))
        if *qptr {
            fmt.Println(a)
        } else {
            fmt.Printf("factorial of %d is %d\n", *nptr, a)
        }
    } else {
        for n := 2; n <= LIMIT; n++ {
            a := factorial(int64(n))
            fmt.Printf("factorial of %d is %d\n", n, a)
            // fmt.Printf("factorial of %d is %x\n", n, a)
        }
    }
}
