package main

import (
    "encoding/csv"
    "fmt"
    "io"
    "os"
)

func main() {
    file, err := os.Open("zipcodes.csv")
    // file, err := os.Open("/Users/dpw/MyDownloads/uszipsv1.4.csv")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    count := 0
    reader := csv.NewReader(file)
    for {
        record, err := reader.Read()
        if err != nil {
            if err == io.EOF {
                fmt.Println("read complete...")
                break
            } else {
                panic(err)
            }
        }

        // fmt.Println(record)
        if count != 0 {
            fmt.Printf("%s %s %s\n", record[0], record[4], record[3])
        }

        count++
    }

    fmt.Println(count, "lines/records read")
}
