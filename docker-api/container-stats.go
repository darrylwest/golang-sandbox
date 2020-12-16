//
// list containers: see https://godoc.org/github.com/moby/moby/client 
//
// @author darryl.west <darwest@ebay.com>
// @created 2017-07-03 12:48:30
//

package main

import (
    "context"
    "fmt"

    "github.com/docker/docker/api/types"
    "github.com/docker/docker/client"
)

func main() {
    ctx := context.Background()
    cli, err := client.NewEnvClient()
    if err != nil {
        panic(err)
    }

    opts := types.ContainerListOptions{}
    opts.All = true
    containers, err := cli.ContainerList(ctx, opts);
    if err != nil {
        panic(err)
    }

    for _, container := range containers {
        fmt.Printf("%s %s %s\n", container.ID[:10], container.State, container.Names[0])
        stats, err := cli.ContainerStats(ctx, container.ID, false)
        if err != nil {
            panic(err)
        }

        fmt.Printf("%v\n", stats)
    }
}
