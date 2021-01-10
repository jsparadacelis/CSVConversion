package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func sortChunk(chunk []int, channel chan []int) {
	fmt.Println("Chunk to sort: ", chunk)
	sort.Ints(chunk)
	channel <- chunk
}

func getNumbersToSort() []int {
	var numbers []int
	fmt.Println("Please input some integers to sort separated by a comma: ")
	br := bufio.NewReader(os.Stdin)
	a, _, _ := br.ReadLine()
	ns := strings.Split(string(a), ",")

	for _, s := range ns {
		n, _ := strconv.Atoi(s)
		numbers = append(numbers, n)
	}
	return numbers
}

func main() {

	channel := make(chan []int)

	numbers := getNumbersToSort()

	wholeListSorted := make([]int, 0, 4)
	remainingChunks := len(numbers)

	if remainingChunks >= 4 {
		currentPos := 0
		for i := 4; i > 0; i-- {
			nextDistance := int(math.Ceil(float64(remainingChunks / i)))
			remainingChunks = remainingChunks - nextDistance
			chunk := numbers[currentPos : currentPos+nextDistance]

			go sortChunk(chunk, channel)

			resultFromChannel := <-channel
			wholeListSorted = append(wholeListSorted, resultFromChannel...)
			currentPos = currentPos + nextDistance
		}
		sort.Ints(wholeListSorted)
		numbers = wholeListSorted
	}
	sort.Ints(numbers)
	fmt.Println(numbers)

}
