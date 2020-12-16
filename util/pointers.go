package main

import "fmt"

// POINTERS

func main() {

	// We pass the value of a variable to the function
	x := 0
	changeXVal(x)
	fmt.Println("x =", x, " no change when passed by value")

	// If we pass a reference to the variable we can change the value
	// in a function
	changeXValNow(&x)
	fmt.Println("x =", x, " changes when passed by reference")

	// Get the address x points to with &
	fmt.Println("Memory Address for x =", &x)

	// We can also generate a pointer with new

	yPtr := new(int)
	changeYValNow(yPtr)
	fmt.Println("y =", *yPtr, " value changes with passed pointer")

}

func changeXVal(x int) {

	// Has no effect on the value of x in main()
	x = 2

}

// * signals that we are being sent a reference to the value
func changeXValNow(x *int) {

	// Change the value at the memory address referenced by the pointer
	// * gives us access to the value the pointer points at

	*x = 2 // Store 2 in the memory address x refers to

}

func changeYValNow(yPtr *int) {

	*yPtr = 100

}
