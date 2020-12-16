package main


import "fmt"

func sum(numbers ...int) int {
    s := 0
    for _, v := range numbers {
        s += v
    }

    return s
}

func main() {
    list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
    x := sum(list...)
    fmt.Printf("sum of %v = %d\n", list, x)
}

