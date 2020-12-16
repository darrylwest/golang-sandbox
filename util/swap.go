package main

func main() {

	a := 34
	b := 55

	println(a, b)

	b, a = a, b
	println(a, b)
}
