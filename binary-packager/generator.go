
package main

import (
    "net/http"
    "github.com/shurcooL/vfsgen"
)


func main() {
    opts := vfsgen.Options{
        Filename: "assets.go",
    }

    var fs http.FileSystem = http.Dir("public")
    if err := vfsgen.Generate(fs, opts); err != nil {
        panic(err)
    }

}
