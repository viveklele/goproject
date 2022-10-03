package main

import (
	"fmt"
)

func main() {

	fmt.Println("Maps in golang")

	language := make(map[string]string)

	language["JS"] = "javascript"
	language["RB"] = "Ruby"
	language["PY"] = "Python"

	fmt.Println("List of all languages:", language)
	fmt.Println("The value for RB is", language["JS"])

	// Delete word
	delete(language, "RB")

	fmt.Println("List of all languages:", language)

	// Loops are interesting in go lang

	for key, value := range language {
		fmt.Println(key, value)
	}

	for _, value := range language {
		fmt.Printf("The value is %v\n", value)
	}
}
