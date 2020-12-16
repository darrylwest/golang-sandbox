package main

import (
	"fmt"

	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data, err := ioutil.ReadFile("./api/swagger/swagger.yaml")
	check(err)
	// fmt.Print(string(data))

	m := make(map[interface{}]interface{})
	err = yaml.Unmarshal(data, &m)
	check(err)

	yml, err := yaml.Marshal(&m)
	check(err)
	fmt.Printf("\n%s\n", string(yml))
}
