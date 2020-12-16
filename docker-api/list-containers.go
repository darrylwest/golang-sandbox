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
    "strings"

    "github.com/docker/docker/api/types"
    "github.com/docker/docker/client"
)

func main() {
    cli, err := client.NewEnvClient()
    if err != nil {
        panic(err)
    }

    opts := types.ContainerListOptions{}
    opts.All = true
    containers, err := cli.ContainerList(context.Background(), opts);
    if err != nil {
        panic(err)
    }

    for _, container := range containers {
        if strings.Contains(container.Names[0], "service-tester-") {
            fmt.Printf("%s : %s : %s : %s\n", container.ID, container.State, container.Status, container.Names[0])
            fmt.Printf("\tImage  : %s Command: %s\n", container.Image, container.Command)
            fmt.Printf("\tNetwork: %v\n", container.NetworkSettings)
            fmt.Printf("\tMounts : %v\n", container.Mounts)
            fmt.Printf("\tPorts  : %v count: %d public: %d\n", container.Ports, len(container.Ports), container.Ports[0].PublicPort)
            fmt.Println("")
        }
    }
}
