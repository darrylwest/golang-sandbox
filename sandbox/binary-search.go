package main

import (
    "fmt"
)

// iterative but could be refactored for recursive
func binarySearch(list []int, target int) (int, bool) {
    low, high := 0, len(list) - 1

    for low <= high {
        mid := low + (high - low) / 2
        value := list[mid]

        switch {
        case value < target:
            low = mid + 1
        case value > target:
            high = mid - 1
        default: // found it
            return mid, true
        }
    }

    // not found...
    return -1, false
}

func main() {
    sortedList := []int{ 1, 3, 4, 6, 7, 9, 10, 11, 13 }
    fmt.Println(sortedList)

    count := 0
    for target := -2; target < 15; target++ {
        if idx, ok := binarySearch(sortedList, target); ok {
            fmt.Printf("%d found in list at index %d\n", target, idx)
            count++
        } 
    }

    if count == len(sortedList) {
        fmt.Printf("found all elements in list, count %d\n", count)
    } else {
        fmt.Printf("ERROR! should have found %d elements, found %d\n", len(sortedList), count)
    }
}

