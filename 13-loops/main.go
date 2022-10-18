package main

import "fmt"

func main() {

	i := 1
	for i <= 5 {
		fmt.Println(i)
		i++

	}

	for i := 5; i <= 10; i++ {
		fmt.Println(i)
	}

	for i := 1; i <= 50; i++ {
		if i == 12 {
			break // print till 11
		}
		fmt.Println(i)
	}

	for i := 1; i <= 50; i++ {
		if i == 12 {
			continue // skip 12
		}
		fmt.Println(i)
	}

	progLang := []string{"Js", "Go", "Py", "Rb", "c", "c++"}

	for i := 0; i < len(progLang); i++ {
		fmt.Println(progLang[i])
	}

	for i := 0; i < len(progLang); i++ {

		if progLang[i] == "Js" {
			continue
		}

		fmt.Println(progLang[i])
		goto js // at the end of loop execute
	}

js:
	fmt.Println("js is dynamic type language")
}
