package main

import (
	"fmt"
	"sort"
)

func main() {

	var progLang = []string{"c", "c++", "java", "python", "Go"}
	fmt.Printf("type of progLAnf is %T \n", progLang)
	fmt.Println(progLang)

	progLang = append(progLang, "rust")
	fmt.Println(progLang)

	progLang = append(progLang[1:])
	fmt.Println(progLang)

	gpas := make([]int, 4)
	gpas[0] = 7
	gpas[1] = 8
	gpas[2] = 9
	gpas[3] = 10
	fmt.Println(gpas)

	gpas = append(gpas, 11, 5, 3)
	fmt.Println(gpas)

	sort.Ints(gpas)
	fmt.Println(gpas)

	// deleting element from slice
	var jsLibrary = []string{"react", "Anguler", "svelt", "NextJs", "Remix"}
	fmt.Println(jsLibrary)
	var index int = 2

	jsLibrary = append(jsLibrary[:index], jsLibrary[index+1:]...)
	fmt.Println(jsLibrary)

	// length and cap of slice

	fmt.Println("len of jsLibray slice is : ", len(jsLibrary))
	fmt.Println("capacity of jsLibray slice is : ", cap(jsLibrary))

}
