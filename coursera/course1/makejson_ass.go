package main

import (
	"encoding/json"
	"fmt"
)

func getPersonMap(name string, address string) map[string]string {
	personMap := make(map[string]string)
	personMap["name"] = name
	personMap["address"] = address
	fmt.Println("personMap ", personMap)
	return personMap
}

func main() {
	fmt.Println("input name")
	var name string
	fmt.Scan(&name)
	fmt.Println("input address")
	var address string
	fmt.Scan(&address)

	personMap := getPersonMap(name, address)

	barr, err := json.Marshal(personMap)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("json is ", string(barr))

	var p2 map[string]string
	err = json.Unmarshal(barr, &p2)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("origin obj converted from json is ", p2)
}
