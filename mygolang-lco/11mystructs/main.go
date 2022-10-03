package main

import (
	"fmt"
)

func main() {
	fmt.Println("Struct in Go lang")
	// No inheritance in golang; No super or parent

	Vivek := User{"Vivek", "vivek@123", true, 21}
	fmt.Println(Vivek)
	fmt.Printf("Vivek details are %+v\n", Vivek)
	fmt.Printf("Name is %v and email is %v", Vivek.Name, Vivek.Email)
		
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
