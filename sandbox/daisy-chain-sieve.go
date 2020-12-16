package main

import (
    "fmt"
    "os"
)

// typical generator that starts at 2
func Generate(out chan<- int) {
    for i := 2; ; i++ {
        out <- i
    }
}

func Filter(in <-chan int, out chan<- int, prime int) {
    m := prime * prime
    for {
        i := <- in
        for i > m {
            m = m + prime
        }
        if i < m {
            out <- i
        }
    }
}

func Sieve(out chan<- int) {
    gen := make(chan int)
    go Generate(gen)

    p := <- gen
    out <- p

    base_primes := make(chan int)
    go Sieve(base_primes)
    bp := <- base_primes
    bq := bp * bp

    for {
        p = <- gen
        if p == bq {
            ft := make(chan int)
            go Filter(gen, ft, bp)
            gen = ft
            bp = <- base_primes
            bq = bp * bp
        } else {
            out <- p
        }
    }
}

func main() {
    sv := make(chan int)
    go Sieve(sv)
    limit := 100000

    file, err := os.Create("primes.txt")
    if err != nil {
        panic(err)
    }

    defer file.Close()

    for i := 0; i < limit; i++ {
        prime := <- sv
        fmt.Fprintf(file, "%d\n", prime)

        if i % 10000 == 0 {
            file.Sync()
        }

        // show the last 20...
        if i >= (limit - 20) {
            fmt.Printf("%4d ", prime)

            if (i + 1) % 20 == 0 {
                fmt.Println()
            }
        }

    }

    file.Sync()
}

