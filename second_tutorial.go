package main

import "fmt"

// two slashes are used for commenting

/*

Multi-line comments are done this way

*/

// VARIABLES

//Unsigned integers are positive numbers
//unint8 0 to 255
//unint16 0 to 65535 and so on
//int8 -128 to 127
//int16 -32768 to 32767 and so on
//unint unsigned, either 32 or 64 bits
//ing signed, either 32 or 64 bits

//float32 32-bit floating-point nubers 1.23555
//float64 64-bit floating-point nubers 1.23555

//byte alias for unint8
//rune alias for int32

//bool true or false
//strings text "text looks like this"

//How to declare variable?

//One example is explicit declaration
//We have to DECLARE with - var | variableName | type
//Creates a location in memory to hold the declared TYPE
//Location called greeting and hold a string

var greeting string

func main() {
	//Store a value in our variable aka INITIALIZATION
	greeting = "What it is"
	fmt.Println(greeting)
	//IMPLICIT DECLARATION - means declare and initialize at the same time using :=
	a := 30
	fmt.Println(a)
	year := "2021"
	myString := "Happy New Year! it's "
	myBool := 5 > 8

	//we can also use var with imlicit declaration

	var myFloat float32 = 5.8 + 2.4
	fmt.Println(myString + year)
	fmt.Println(myBool)
	fmt.Println(myFloat)
}