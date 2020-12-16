package main

import (
    "fmt"
    "syscall"
    "os"
)

var stat syscall.Statfs_t

func main() {
    wd, err := os.Getwd()
    if err != nil {
        panic(err)
    }

    syscall.Statfs(wd, &stat)

    fmt.Println(stat.Bavail * uint64(stat.Bsize))
}
