package main

import "fmt"

func main() {

	var progLang [6]string

	progLang[0] = "c"
	progLang[1] = "c++"
	progLang[2] = "java"
	progLang[3] = "python"
	progLang[5] = "Golang"

	fmt.Println(progLang)      // notic space b/w python and golang bcz it initalized 4th index with empty string
	fmt.Println(len(progLang)) // 6 (declared value)
	progLang[0] = "c#"         // reassignment
	fmt.Println(progLang[0])
	fmt.Println(progLang[0:3])

	var cloudProviders = [4]string{"AWS", "Azure", "Gcp", "IBM"}
	fmt.Println(cloudProviders)

	countryList := [...]string{"India", "USA", "UK", "Russia"} // compiler automatic determine it's size
	fmt.Println(countryList)
}

/*
=>  array is a collection of data elements of the same type can access each element by an index. fixed size
=> Due to Go being a pass by value language, modifying a normal array parameter won’t create permanent change Sometimes this can be useful in performing local calculations.

*/
