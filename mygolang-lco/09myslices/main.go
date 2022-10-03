package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("Welcome to vide on slices")

	var fruitList = []string{"Mango", "Tomato", "Nothing"}
	fmt.Printf("The data type of fruitList is: %T\n", fruitList)
	fmt.Printf("The lenth of fruitList is %v\n", len(fruitList))

	// Adding data in slices
	fruitList = append(fruitList, "Apple")

	fmt.Println(fruitList)

	for i := 0; i < 3; i++ {
		fmt.Println("The list of Fruits is:", fruitList[i])
	}
	// slicing of slice
	fruitList = append(fruitList[:3])
	fmt.Println("The list of Fruits is:", fruitList)

	// data type and size
	highScore := make([]int, 4)

	highScore[0] = 234
	highScore[1] = 414
	highScore[2] = 532
	highScore[3] = 624

	highScore = append(highScore, 123, 532, 631)
	fmt.Println(highScore)

	sort.Ints(highScore)
	fmt.Println(highScore)

	// how ro remove a value from slices base on  index

	var cources = []string{"Python", "Java", "C++", "reactjs", "javascript", "ruby"}
	var index int = 2

	cources = append(cources[:index], cources[index+1:]...)
	fmt.Println(cources)

}
