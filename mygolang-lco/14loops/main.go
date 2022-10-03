package main

import "fmt"

func main() {
	fmt.Println("loops for golang")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Friday", "Saturday"}
	fmt.Println(days)

	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])

	}
	// Another type of loop using range function
	for i := range days {
		fmt.Println(days[i])
	}
	// For loop using index
	for index, day := range days {
		fmt.Printf("Index value is %v Day value is %v\n", index, day)
	}

	for _, day := range days {
		fmt.Printf("Day value is %v\n", day)
	}

	roughValue := 1

	for roughValue < 10 {

		if roughValue == 2 {
			//Provides an unconditional jump from the goto to a labeled statement in the same function.
			goto google
		}

		fmt.Println(roughValue)

	google:
		fmt.Println("From goto statment")
		if roughValue == 5 {
			roughValue++
			continue
		}
		if roughValue == 6 {
			break
		}
		roughValue++
	}

}
