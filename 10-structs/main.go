package main

import "fmt"

type User struct {
	Name       string
	UserName   string
	Email      string
	IsPaidUser bool
}

func main() {

	jay := User{"jay", "jay07", "jay@gmail.com", true}

	fmt.Println("jay's details are : ", jay)
	fmt.Printf("jay's details are : %+v \n", jay)
	fmt.Println("jay's email is  : ", jay.Email)

}

// grouping together related variables is done using a struct
// Struct names begin with a capital letter in Go
/* type NameOfStrcut struct {

}
*/
// Go allows us to rely on default values as well. We can omit fields:
// https://www.codecademy.com/courses/learn-go-loops-arrays-maps-and-structs/lessons/structs-in-go/exercises/go-functions-that-access-a-struct
