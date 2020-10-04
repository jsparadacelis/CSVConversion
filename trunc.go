package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello, insert your float number")
	var number float64
	fmt.Scan(&number)
	fmt.Printf("Integer part of your number is: %v", math.Trunc(number))
}
