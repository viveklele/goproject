package main

import ("fmt"
		"time"
)

func main(){
	fmt.Println("Welcome to time study of golang")

		time.Sleep(1)
		presentTime := time.Now()
		fmt.Println(presentTime.Format("01-02-2006 Monday 15:04:05"))

		createDate := time.Date(2020, time.August, 10, 23, 23, 0, 0, time.UTC)
		fmt.Println(createDate.Format("01-02-2006 Monday"))
}	