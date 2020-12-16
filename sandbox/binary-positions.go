package main

import (
    "fmt"
    "strconv"
)

func majorIndex(number string) int {
    var zero byte = 48
    var one byte = 49
    n, err := strconv.Atoi(number)
    if err != nil {
        panic(err)
    }
    s := []byte(fmt.Sprintf("%b", n))

    pos := 1
    sum := 0
    right := len(s) - 1
    last := s[right]
    right--

    for i := right; i >= 0; i-- {
        chr := s[i]
        if last == one && chr == zero {
            sum += pos
        }

        last = chr
        pos++
    }
 
    return sum
}

func main() {
    fmt.Println(83, majorIndex("83"))
}
