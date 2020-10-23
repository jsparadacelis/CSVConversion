package main

import "fmt"

func swap(indexToChange int, arrToSort *[]int) {
	aux := (*arrToSort)[indexToChange+1]
	(*arrToSort)[indexToChange+1] = (*arrToSort)[indexToChange]
	(*arrToSort)[indexToChange] = aux
}

func bubbleSort(arrToSort *[]int) {
	for i := 0; i < len(*arrToSort); i++ {
		for j := 0; j < len(*arrToSort)-i-1; j++ {
			if (*arrToSort)[j] > (*arrToSort)[j+1] {
				swap(j, arrToSort)
			}
		}
	}
}

func main() {

	var arrToSort []int
	count := 0
	fmt.Println("Hi! Please, insert ten numbers")
	for count < 10 {
		var numberToArr int
		fmt.Printf("\nInsert #%v: ", count+1)
		fmt.Scan(&numberToArr)
		count++
		arrToSort = append(arrToSort, numberToArr)
	}
	fmt.Printf("Array inserted: %v", arrToSort)
	bubbleSort(&arrToSort)
	fmt.Printf("\nArray sorted: %v", arrToSort)
}
