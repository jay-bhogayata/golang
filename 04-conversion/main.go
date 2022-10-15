package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter rating for our services : ")

	input, _ := reader.ReadString('\n')

	fmt.Println(input)

	incNum, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		panic(err)
	} else {
		fmt.Println(incNum + 1)
	}

}
