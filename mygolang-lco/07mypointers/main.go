package main

import (
	"fmt"
)

func main() {

	fmt.Println("Welcome to class on pointers")

	//var ptr *int
	// ptr = 21
	var mynumber = 32

	var ptr = &mynumber
	fmt.Println("The value of pointer is ", ptr)
	fmt.Println("The value of pointer is ", *ptr)

	*ptr = *ptr + 2
	fmt.Println("New value is: ", mynumber)
		
}
