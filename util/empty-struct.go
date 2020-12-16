package main

type Lamp struct{}

func (l Lamp) Off() {
	println("lamp off")
}

func (l Lamp) On() {
	println("lamp on")
}

func main() {
	var lamp Lamp
	lamp.On()
	lamp.Off()

	Lamp{}.On()
	Lamp{}.Off()
}
