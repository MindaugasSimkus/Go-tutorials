package main

import "fmt"

// POINTERS
// Pointer holds the memory address of another variable

// & has one meaning
// & symbol means: THE ADDRESS OF the variable it is next to
// Example: If a := 25 and we set b := &a
// The value of b in the memory address of a
// And it type is a pointer to an int - *int

// * has two meanings
// 1: The * symbol next to VARIABLE means: Get the value of the variable that this
// pointer is pointing to aka DEREFERENCING

// Example: If we set n:= *b
// The value of n is 25. Its the value stored at the memory address that b is pointing to

// 2: The * symbol next to a TYPE means the variable being declared
// is a pointer and it points to an address holding the type followed by the *

// Example: var *string myName
// myName is variable which holds the memory address of a string variable

func main1() {
	a := 25
	b := &a
	fmt.Println(b)
	c:= &b
	// Could have written
	// var c *(*int) = &b
	fmt.Println(*c) // that is Pointers!
}
	// why use pointers?
	// One use case is to alter the variables passed into FUNCTIONS
	// When we call a function that takes an argument, that argument is copied ti the function:

func zero(x int) {
	x = 0
}

func zero2(x *int) {
	*x = 0
}

func main2() {
		x := 5
		z := 5
		zero(x)
		zero2(&z)
		fmt.Println(x)
		fmt.Println(z)
}

// Pointers can reveal if a value has actually been set or if its just the default value

	var isSeventeen bool
	
func main() {

	// default type for booleans is false
	// default value for interfaces, slices, channels, maps, functions and POINTERS is nil!

	var isSeventeenPtr *bool
	fmt.Println(isSeventeen)
	fmt.Println(isSeventeenPtr)

	input := false

	isSeventeenPtr = &input

	fmt.Println(isSeventeenPtr)
	fmt.Println(*isSeventeenPtr)
}