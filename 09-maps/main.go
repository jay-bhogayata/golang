package main

import (
	"fmt"
)

func main() {

	// maps : kv pair's

	// variableName := make(map[keyType]valueType)

	/* variableName := map[keyType]valueType{
		name1: value1,
		name2: value2,
		name3: value3,
	}
	*/

	// yourMap[newKey] = newValue

	progLang := make(map[string]string)
	progLang["go"] = "Golang"
	progLang["py"] = "python"
	progLang["js"] = "javascript"
	progLang["sh"] = "shellscript"

	fmt.Println(progLang)
	fmt.Println(progLang["sh"])

	delete(progLang, "py")
	fmt.Println(progLang)

	for key, value := range progLang {
		fmt.Printf("Key is %v and value is %v \n", key, value)
	}

}

/*
A map is an unordered collection of keys and values.

*/
