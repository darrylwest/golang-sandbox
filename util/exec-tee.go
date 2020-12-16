/**
 * this example replicates what the unix tee command does
 */
package main

import (
    "bufio"
    "fmt"
	"log"
    "io"
    "os"
	"os/exec"
)

func main() {
    filename := "exec.log"
    log.Printf("open file: %s\n", filename)
    fout, err := os.Create(filename)
    if err != nil {
        log.Fatal(err)
    }
    defer fout.Close()

    cmd := exec.Command("test-logger.sh")

    // replace this with a stream reader/writer
    pr, pw := io.Pipe()
    defer pw.Close()

    reader := io.MultiWriter(fout, pw)

    cmd.Stdout = reader
    cmd.Stderr = reader

    go func() {
        defer pr.Close()

        lineReader := bufio.NewReader(pr)

        lc := 0
        for  {
            line, err := lineReader.ReadString('\n')
            if err != nil {
                break
            }

            lc++
            fmt.Printf("%d) %s", lc, line)
        }
    }()

    if err := cmd.Run(); err != nil {
        log.Fatal(err)
    }
}
