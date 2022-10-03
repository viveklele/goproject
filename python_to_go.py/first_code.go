package main

import "fmt"

type UserData struct {
		i int
		name string
}

func main() {
	name := "Vivek"
	nameList := make([]UserData)
	
	var userData = UserData{
		name : name,
		
	}
	nameList = append(nameList, userData)


	// get user name
	for i := 1; i <= 10; i++ {
		fmt.Printf("%v%v\n",name, i)
		nameList = append(nameList )
		fmt.Println(nameList)
	}

	// // get user email
	// for i := 1; i <= 10; i++{
	// 	fmt.Printf("")

	// }

}
