package main

import (
	"encoding/csv"
	"encoding/json"
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
	resultList := make([]byte, 0, 5)
	for _, record := range allRecords[1:] {

		newSchool := NewSchool(record...)
		newSchoolJSON, _ := json.Marshal(*newSchool)
		resultList = append(resultList, newSchoolJSON...)
	}
	fmt.Println(string(resultList))

}
