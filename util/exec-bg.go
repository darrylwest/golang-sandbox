//
// A very simple way to create a daemon; 
//
// TODO enhance this with [start|status|stop] and write the pid to a file...
//
// @author darryl.west <darryl.west@raincitysoftware.com>
// @created 2017-09-04 15:15:03
//

package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {

    cmd := exec.Command("ls", "-l")
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    if err := cmd.Start(); err != nil {
        fmt.Println(err)
        return
    }

    if err := cmd.Wait(); err != nil {
        fmt.Println(err)
        return 
    }

    fmt.Printf("pid: %d, ok...\n", cmd.Process.Pid)
}
