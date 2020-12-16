package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    file, err := os.Open("zipcodes.csv")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    count := 0
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        fmt.Println(scanner.Text())

        count++
    }

    fmt.Println(count, "lines read")
    if err := scanner.Err(); err != nil {
        fmt.Println(err)
    }
}
