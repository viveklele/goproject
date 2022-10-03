package main

import (
	"fmt"
)

func main() {
	fmt.Println("if else in golang")
	var result string

	loginCount := 32

	if loginCount < 10 {
		result = "Regular result"
	} else if loginCount > 10 {
		result = "Ok"
	} else {
		result = "Something else"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}
	if num := 3; num < 10 {
		fmt.Println("Num is less then 10")
	} else {
		fmt.Println("Num is greter then 10")
	}
} 