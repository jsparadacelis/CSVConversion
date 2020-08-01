package main

import (
	"fmt"
	"strconv"
)

func main() {
	testNumber := 12345
	resultSum := giveMeSum(testNumber)
	fmt.Println("La suma de las cifras de 12345 es 15", resultSum == 15)
	testNumberBinary := 10
	intToBinary := intToBinary(testNumberBinary)
	fmt.Println("10 en binario es 1010", intToBinary == "1010")

}

func giveMeSum(testNumber int) int {

	var totalSum int
	for testNumber > 0 {
		lastCipher := testNumber % 10
		testNumber = testNumber / 10
		totalSum = totalSum + lastCipher
	}
	return totalSum
}

func intToBinary(testNumber int) string {
	var resultBinary []int
	var resultAsString string
	for testNumber > 1 {
		resultBinary = append(resultBinary, testNumber%2)
		testNumber = testNumber / 2
	}
	resultBinary = append(resultBinary, testNumber)
	for i := len(resultBinary) - 1; i > -1; i-- {
		resultAsString = resultAsString + strconv.Itoa(resultBinary[i])
	}
	return resultAsString
}
