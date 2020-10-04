package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	slice := make([]int, 0, 3)
	for {
		var numberToAdd string
		fmt.Println("Insert a number (scape X):")
		fmt.Scanln(&numberToAdd)
		if strings.ToLower(numberToAdd) == "x" {
			fmt.Println("Scape!")
			break
		} else {
			valueAsInt, _ := strconv.Atoi(numberToAdd)
			slice = append(slice, valueAsInt)
			sort.Ints(slice)
			fmt.Printf("Slice sorted: %v \n", slice)
		}
	}

}
