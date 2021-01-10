package main

import (
	"fmt"
	"time"
)

var count int

func addAndPrint() {
	fmt.Println(count)
	count++
}

func main() {

	count = 1
	addAndPrint()
	addAndPrint()
	time.Sleep(1 * time.Second)
}
