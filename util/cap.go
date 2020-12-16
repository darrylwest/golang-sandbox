package main

import "fmt"

func main() {
    const n = 7
    var a = []uint{n:n}
    fmt.Println(a, len(a),  cap(a))
}
