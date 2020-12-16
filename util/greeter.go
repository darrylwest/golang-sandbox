package main

import "fmt"

type Greeter struct {
	Format string
}

func (g Greeter) Greet(name string) string {
	if g.Format == "" {
		g.Format = "Hi %s"
	}

	return fmt.Sprintf(g.Format, name)
}

func main() {
	g := Greeter{
		Format: "Howdy %s",
	}

	fmt.Println(g.Greet("flarb"))
}
