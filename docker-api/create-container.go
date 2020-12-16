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
    "io"
    "os"
    "strings"
    "time"

    "github.com/docker/docker/api/types"
    "github.com/docker/docker/api/types/container"
    "github.com/docker/docker/client"
)

func main() {
    cli, err := client.NewEnvClient()
    if err != nil {
        panic(err)
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    name := "alpine-hello"
    image := "alpine:latest"

    /*
    _, err = cli.ImagePull(ctx, image, types.ImagePullOptions{})
    if err != nil {
        panic(err)
    }
    */
    
    if list, err := cli.ContainerList(ctx, types.ContainerListOptions{ All:true }); err == nil {

        for _, c := range list {
            fmt.Printf("%s\n", c.Names[0])

            if strings.HasSuffix(c.Names[0], name) {
                id := c.ID
                fmt.Printf("%s %s\n", id, c.Names[0])

                // stop and wait...
                if c.State == "running" {
                    cli.ContainerStop(ctx, id, nil)

                    cli.ContainerWait(ctx, id, "")
                } else {
                    fmt.Printf("remove %s\n", name)
                    cli.ContainerRemove(ctx, image, types.ContainerRemoveOptions{})
                }

                break
            }
        }
    }

    conf := &container.Config{
        Image: image,
        Cmd: []string{"echo", "howdy"},
    }
    
    resp, err := cli.ContainerCreate(ctx, conf, nil, nil, name)
    if err != nil {
        panic(err)
    }

    fmt.Printf("%v\n", resp)
    if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
        panic(err)
    }

    out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
    if err != nil {
        panic(err)
    }

    io.Copy(os.Stdout, out)
}
