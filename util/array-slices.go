package main

import "fmt"

type MyType struct {
    name string
}


func show(ss []MyType) {
    fmt.Println(ss, ss[len(ss) - 1])
}

func main() {
	x := []string{"a", "b", "c"}

	for i, v := range x {
		fmt.Println(i, v)
	}

	n := 9
	row := make([]int, n+1, n+1)
	row[0], row[n] = 1, 1

	fmt.Println(row)

    maxslice := 2
    ss := []MyType{}

    ss = append(ss, MyType{ "fred" })
    ss = append(ss, MyType{ "sally" })

    show(ss)

    ss = append(ss[1:maxslice], MyType{ "sam" })
    show(ss)

    ss = append(ss[1:maxslice], MyType{ "larry" })
    show(ss)

    ss = append(ss[1:maxslice], MyType{ "moe" })
    show(ss)

    // now reverse...
	x = []string{"a", "b", "c", "d", "e", "f"}
    fmt.Println(x)
    n = len(x)
    for i := 0; i < (n / 2); i++ {
        j := n - i - 1
        x[i], x[j] = x[j], x[i]
    }
    fmt.Println(x)
}
