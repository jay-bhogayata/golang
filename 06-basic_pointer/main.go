package main

import "fmt"

func main() {
	myNumber := 23

	var ptr = &myNumber

	fmt.Println("Value of actual pointer is ", ptr)
	fmt.Println("Value of actual pointer is ", *ptr)

	*ptr = *ptr + 2
	fmt.Println("New value is: ", myNumber)

	num := 1
	fmt.Println(num)
	add100(num)
	fmt.Println(num) // print 1 bcz copy of num is passed in a function

	a := 10
	p1 := &a
	fmt.Println(a)
	*p1 = 20
	fmt.Println(a)

}
func add100(num int) {
	num += 100
}
