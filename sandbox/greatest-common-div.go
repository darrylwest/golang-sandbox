package main

import (
	"fmt"
    "math/big"
)

type BinaryEuclidianGCD func(a, b, res int) int

func gcd(a, b int) int {
	// binary euclidian algorithm
	var bgcd BinaryEuclidianGCD

	bgcd = func(a, b, res int) int {
		switch {
		case a == b:
			return res * a
		case a%2 == 0 && b%2 == 0:
			return bgcd(a/2, b/2, 2*res)
		case a%2 == 0:
			return bgcd(a/2, b, res)
		case b%2 == 0:
			return bgcd(a, b/2, res)
		case a > b:
			return bgcd(a-b, b, res)
		default:
			return bgcd(a, b-a, res)
		}
	}

	return bgcd(a, b, 1)
}

func libgcd(a, b int) int64 {
    x, y := big.NewInt(int64(a)), big.NewInt(int64(b))
    
    z := x.GCD(nil, nil, x, y).Int64()

    return z
}

func main() {
	type pair struct{ a, b int }
	pairs := []pair{
		pair{24, 9},
		pair{75, 15},
		pair{15, 9},
		pair{33, 77},
		pair{49865, 69811},
		pair{97, 101},
		pair{1071, 1029},
		pair{3528, 3780},
		pair{100, 150},
		pair{15, 75},
	}

	for _, p := range pairs {
		a, b := p.a, p.b
		fmt.Printf("%d, %d gcd: %d\n", a, b, gcd(a, b))
		fmt.Printf("%d, %d gcd: %d\n", a, b, libgcd(a, b))
	}

}
