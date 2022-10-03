package main
		
import ("fmt"
		"bufio"
		"os"
		"strings"
		"strconv")

func main(){

	fmt.Println("Welcome to out pizza app")
	fmt.Println("Plear rate our pizza between 1 to 5")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	fmt.Println("Thenk for rating", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil{
		fmt.Println("Please Enter number")

	} else {
		fmt.Println("Added to 1 to your rating", numRating+1)
	}

}