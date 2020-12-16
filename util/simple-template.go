package main

import (
    "fmt"
    "os"
    "text/template"
)

func main() {
    const letter = `Dear {{.}}, How are you?`

    tmpl, err := template.New("letter").Parse(letter)
    if err != nil {
        fmt.Println(err.Error())
    }
    tmpl.Execute(os.Stdout, "Professor Doolittle")
}
