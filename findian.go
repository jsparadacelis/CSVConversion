package main

import (
	"fmt"
	"strings"
)

func main() {
	var stringToTest string
	fmt.Println("Hello, insert your string:")
	fmt.Scanln(&stringToTest)
	if strings.HasPrefix(stringToTest, "i") && strings.HasSuffix(stringToTest, "n") && strings.Contains(stringToTest, "a") {
		fmt.Print("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
