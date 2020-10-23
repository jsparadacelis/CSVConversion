package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type CompleteName struct {
	fname string
	lname string
}

func main() {
	var filename string
	sliceName := make([]CompleteName, 0)

	fmt.Println("Welcome, please insert the filename: ")
	fmt.Scan(&filename)
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		dataAsLines := strings.Split(string(data), "\n")
		for _, completeName := range dataAsLines {
			nameArr := strings.Split(completeName, " ")
			name := CompleteName{nameArr[0], nameArr[1]}
			sliceName = append(sliceName, name)
		}

		for i, value := range sliceName {
			fmt.Printf("Struct No. %v\n", i+1)
			fmt.Println("First name: ", value.fname, "\nLast name: ", value.lname)
		}
	}

}
