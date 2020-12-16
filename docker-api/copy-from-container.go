package main

import (
    // "bytes"
    "context"
    "fmt"

    // "os"
    // "io"
    "io/ioutil"

    // "github.com/docker/docker/api/types"
    "github.com/docker/docker/client"
)

func main() {
    ctx := context.Background()
    cli, err := client.NewEnvClient()
    if err != nil {
        panic(err)
    }

    id := "6233e06f19d36f8a3adcf6ae645bd77524ed05dcb7160905ce68439fe8ab5e90"
    path := "/opt/bolt/config/config.json"

    reader, stats, err := cli.CopyFromContainer(ctx, id, path)
    if err != nil {
        panic(err)
    }

    defer reader.Close()
    fmt.Printf("stats: %v\n\n", stats)

    content, err := ioutil.ReadAll(reader)
    if err != nil {
        panic(err)
    }

    fmt.Printf("size: %d\n", len(bytes))
    fmt.Printf("content: %s\n", string(bytes))

}
