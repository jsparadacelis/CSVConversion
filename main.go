package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type School struct {
	id         int
	name       string
	phone      string
	address    string
	city       string
	state      string
	zip        string
	schoolType string
	principal  string
	website    string
	other      string
	image      string
	active     bool
}

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
	resultList := make([]School, 0, 5)
	for index, record := range allRecords[1:] {
		fmt.Println("index", index, "record", record)
	}
}
