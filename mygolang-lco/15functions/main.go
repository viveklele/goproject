package main

import "fmt"

func main() {
	fmt.Println("Welcome to function in golang")
	greeter()

	result := adder(23, 22)
	fmt.Println("The sum of two numbers is: ", result)

	proRes, myMessage := proAdder(2, 4, 4, 32, 342, 2, 4)	
	fmt.Println("Pro result is: ", proRes)
	fmt.Println("My message is:", myMessage)
}

func greeter() {
	fmt.Println("Hello from greeter")
}

func adder(num1 int, num2 int) int {
	return num1 + num2
}

// This syntax is use when you don't know number of inputs
// If you have more then one value to return write data type in round brackets
func proAdder(values ...int) (int, string) {
	total := 0

	for _, value := range values {
		total += value
	}
	return total, "Hi pro result function"
}