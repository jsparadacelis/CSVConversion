package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {

	fileContent, err := os.Open("C2ImportSchoolSample.csv")
	if err == nil {
		fmt.Println("..... Read Successful .....", &fileContent)
	} else {
		fmt.Println("..... Couldn't read .....", err)
	}
	defer fileContent.Close()

	reader := csv.NewReader(fileContent)
	allRecords, _ := reader.ReadAll()
	headers := allRecords[0]
	fmt.Println("headers", headers)
	// resultList := make([]School, 0, 5)
	for _, record := range allRecords[1:] {

		newSchool := NewSchool(record...)
		fmt.Println(newSchool)
	}
}
