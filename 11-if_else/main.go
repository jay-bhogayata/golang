package main

import "fmt"

func main() {
	var num int = 0

	if num%2 == 0 {
		fmt.Println("number is even")
	} else {
		fmt.Println("number is odd")
	}

	if num := 5; num < 10 {
		fmt.Println("Number is a less then 10")
	} else {
		fmt.Println("Number is a not less then 10")
	}

	
}
/*
=> if {} , else if {} , else {}
*/