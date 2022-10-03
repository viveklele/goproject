package main

import (
	"fmt"
)

func main(){

	fmt.Println("Welcome to array in golang")

	// in array you must have to provide data
	var fruitList [4] string

	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[3] = "Nothing"

	fmt.Println("My fruitList is :", fruitList)
	fmt.Println("My fruitList is :", len(fruitList))

	var vegitabel = [3] string{"potato", "beans", "Kobi"}
	
	for i:=0; i<3; i++{
	fmt.Println("My fruitList is :", vegitabel[i])

	}
}