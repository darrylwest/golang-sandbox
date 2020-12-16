package main

import (
    "flag"
	"fmt"

	"io/ioutil"
    "os"

	"gopkg.in/yaml.v2"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}


func main() {
    var file string
    flag.StringVar(&file, "f", "", "set the yaml filename ")
    flag.Parse()

    if file == "" {
        flag.PrintDefaults()
        os.Exit(1)
    }

	data, err := ioutil.ReadFile(file)
	check(err)
	// fmt.Print(string(data))

	m := make(map[interface{}]interface{})
	err = yaml.Unmarshal(data, &m)
	check(err)

	yml, err := yaml.Marshal(&m)
	check(err)
	fmt.Printf("\n%s\n", string(yml))
}
