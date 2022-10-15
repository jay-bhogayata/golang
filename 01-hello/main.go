package main

import (
	"fmt"
	t "time" // here t is a alis of a time

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello world...")
	fmt.Println(quote.Go())
	fmt.Println(t.Now())
	take_input()

}

func take_input() {
	fmt.Println("Enter your name : ")
	var name string
	fmt.Scan(&name)
	fmt.Printf("welcome %v to world of golang", name)
}
