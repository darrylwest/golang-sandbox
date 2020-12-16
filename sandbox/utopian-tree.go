package main
import "fmt"

/*

test with: echo "14 2 4 6 8 10 12 14 16 18 20 22 24 26 60" | go run utopian-tree.go
2 3 11
4 7 111
6 15 1111
8 31 11111
10 63 111111
12 127 1111111
14 255 11111111
16 511 111111111
18 1023 1111111111
20 2047 11111111111
22 4095 111111111111
24 8191 1111111111111
26 16383 11111111111111
60 2147483647 1111111111111111111111111111111

*/

var debug = true

// formula is ...
func calcHeight(cycles int) int {
    n := uint64( cycles )
    return ((1 << (( n >> 1) + 1 )) - 1) << (n % 2)
}

func main() {
    var count, n int
    fmt.Scanf("%v", &count)
    
    for count > 0 {
        fmt.Scanf("%v", &n)
        if debug {
            height := calcHeight(n)
            fmt.Printf("%d %d %b\n", n, height, height)
        } else {
            fmt.Println(calcHeight(n))
        }
        
        count--
    }
}

