package main

import (
	"fmt"
	"time"
)

func main2 () {
	x := 5
	if x < 10 {
		fmt.Println("x is less than 10")
	}
	fmt.Println("This will run no matter what")
}

func main3 () {
	x := 5
	if x == 10 {
		fmt.Println("x is equal to 10")
	} else {
		fmt.Println("x is not equal to 10")
	}
	fmt.Println("This will run no matter what")
}

func main4 () {
	x := 5
	if x > 3 {
		a := x // a is only accessible in this if block
		fmt.Println("x and a are lager than 3", a, x)
	} else {
		fmt.Println("a is not larger than 3")
	}
}

func main5 () {
	x := 5
	if a := x - 3; a > 3 { // variable a is only accessible in this if/else statement block
		fmt.Println("x and a are larger than 3", a, x)
	} else {
		fmt.Println(a)
	}
}

func main6 () {
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}

func main7 () {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

func main8 () {
	names := []string{"Sam", "Tom", "Joe"} // slice of names
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
}

func main9 () {
	names := []string{"Sam", "Tom", "Joe"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
		if i == 1 {
			break // breaks loop
		}
	}
}

func main10 () {
	names := []string{"Sam", "Tom", "Joe"}
	for i := 0; i < len(names); i++ {
		if names[i] == "Sam" {
			continue // allows ti skip Sam
		}
		fmt.Println(names[i])
	}
}

func main11 () {
	i := 0
	for i < 5 {  // instead of while statment we can use for in this way
		fmt.Println(i)
		i = i + 1
	}
}

func main12 () {
	i := 0
	for { // not needed condition here, can be written in a block
		fmt.Println(i)
		i = i + 1
		if i > 5 {
			break
		}
	}
}

func main13 () {
	 names := []string{"Sam", "Tom", "Joe"}
	 for k, v := range names {
	 	fmt.Println(k, v)
	 }
}

func main14 () {
	switch time.Now().Weekday() {
	case time.Saturday:
		fmt.Println("Today is Saturday.")
	case time.Sunday:
		fmt.Println("Today is Sunday.")
	default: // in non case statements are fullfiled, default value is used.
		fmt.Println("Today is a weekday.")
	}
}

func main15 () {
		switch hour := time.Now().Hour(); {// missing expression means "true"
	case hour < 12:
		fmt.Println("Good morning!")
	case hour < 17:
		fmt.Println("Good afternoon!")
	default: 
		fmt.Println("Good evening!")
	}
}

func main() {
	switch 2 { // Golang runs only the chase that matches, fallthrough lets to run all other cases after the matching case
	case 1:
		fmt.Println("1")
		fallthrough
	case 2:
		fmt.Println("2")
		fallthrough
	case 3:
		fmt.Println("3")
	}
}