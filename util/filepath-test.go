package main

/*
    This implementation is a patch to fix windows/dos path name problems with filepath.ToSlash() and filepath.Base().  The objective is to
    extract the base filename then remove the extension.  It fixes a windows problem with ToSlash().
*/

import (
    "fmt"
    "os"
    "strings"
)

func main() {
    stat, _ := os.Stat("/Users/darwest/bolt/TestRunnerService/README.md")

    fmt.Println(strings.Replace(stat.Name(), ".md", "", 1))
}
