package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name string
	var address string

	fmt.Println("Enter a name: ")
	fmt.Scan(&name)
	fmt.Println("Enter an address: ")
	fmt.Scan(&address)

	infoMap := make(map[string]string)
	infoMap["name"] = name
	infoMap["address"] = address

	mapAsJSON, _ := json.Marshal(infoMap)
	fmt.Print("JSON as byte[]: ", mapAsJSON, "\nJSON formatted: ", string(mapAsJSON))

}
