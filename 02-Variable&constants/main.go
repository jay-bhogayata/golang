package main

import (
	"fmt"
)

const Publicvar string = "Go" // public var 1st latter is cap. cnt be accesible outside of the pkg also

const os string = "Linux" // can not access outside this pkg

func main() {
	var name string = "Jay"
	fmt.Printf("My name is %v \n", name)
	fmt.Printf("type of name var is %T \n", name)

	var num int = 15
	fmt.Printf("the given number is a %d \n", num)
	fmt.Printf("type of num var is %T \n", num)
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 (contains only +ve value)

	var gpa float32 = 8.95
	fmt.Printf("My gpa is 2nd sem is %f \n", gpa)
	fmt.Printf("type of gpa var is %T \n", gpa)
	// float32 float64

	// complex64 complex128

	var isPass bool = true
	fmt.Println(isPass)
	fmt.Printf("type of isPass var is %T \n", isPass)

	// implicite type
	var country = "india"
	fmt.Println(country)

	// walrus opeartor
	country2 := "USA"
	fmt.Println(country2)

	fmt.Println(os)

	// zero values
	var a int
	var b string
	var c float64
	var d bool

	fmt.Printf("var a %T = %+v\n", a, a)
	fmt.Printf("var b %T = %q\n", b, b)
	fmt.Printf("var c %T = %+v\n", c, c)
	fmt.Printf("var d %T = %+v\n \n", d, d)

	const pi = 3.14
	//pi = 3.15
	//not possible
	fmt.Println(pi)

}

//  := is for declaration + assignment, whereas = is for assignment only.
// variables are case sensitive, the case of the first letter of a variable has special meaning in Go. If a variable starts with an uppercase letter, then that variable is accessible outside the package it was declared in (or exported). If a variable starts with a lowercase letter, then it is only available within the package it is declared in.
