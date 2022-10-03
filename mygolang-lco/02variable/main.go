package main

import "fmt"

const tocken string = "alsdfue"

func main(){
	var username string = "Vivek"	
	fmt.Println(username)	
	fmt.Printf("Variable type is %T \n",username)

	var smallValue uint8 = 255	
	fmt.Println(smallValue)	
	fmt.Printf("Variable type is %T \n",smallValue)

	var smallFloat float64 = 255.92347293420	
	fmt.Println(smallFloat)	
	fmt.Printf("Variable type is %T \n",smallFloat)

	var anotherVariable int
	fmt.Println(anotherVariable)	
	fmt.Printf("Variable type is %T \n",anotherVariable)

	//implicit type

	var website = "Blog.com"
	fmt.Println(website)

	// no var style

	numberOfUsers := 300
	fmt.Println(numberOfUsers)

	fmt.Println(tocken)
	
}
