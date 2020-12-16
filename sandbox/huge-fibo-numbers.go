package main

import (
  "fmt"
  "math/big"
)

func main() {
  a, b := big.NewInt(0), big.NewInt(1)
  
  var limit big.Int 
  limit.Exp(big.NewInt(10), big.NewInt(99), nil)
  
  for a.Cmp(&limit) < 0 {
    a.Add(a, b)
    a, b = b, a
  }
  fmt.Println(a) // 100-digits...
  
  fmt.Println("prime?", a.ProbablyPrime(20))
}

