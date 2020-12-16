// basic insertion sort implementation
package main

import (
	"fmt"
	"sort"
	// "github.com/pkg/profile"
)

func insertionSort1(a []int) {
	for i := 1; i < len(a); i++ {
		value := a[i]
		j := i - 1
		for j >= 0 && a[j] > value {
			a[j+1] = a[j]
			j = j - 1
		}

		a[j+1] = value
	}
}

func insertionSort(a sort.Interface) {
	for i := 1; i < a.Len(); i++ {
		for j := i; j > 0 && a.Less(j, j-1); j-- {
			a.Swap(j-1, j)
		}
	}
}

func main() {
	// defer profile.Start().Stop()

	list := []int{31, 41, 59, 26, 53, 58, 97, 23, 84}
	fmt.Println(list)

	insertionSort1(list)
	fmt.Println(list)

	list = []int{31, 41, 59, 26, 53, 58, 97, 23, 84}
	insertionSort(sort.IntSlice(list))

	fmt.Println(list)
}
