package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	var name, address string

	fmt.Printf("\nEnter Your Name: ")
	if scanner.Scan() {
		name = scanner.Text()
	}

	fmt.Printf("\nEnter Your Address: ")
	if scanner.Scan() {
		address = scanner.Text()
	}
	//fmt.Printf("\nName: %s",name)
	//fmt.Printf("\nAddress: %s\n",address)

	personal := make(map[string]string)

	personal["name"] = name
	personal["address"] = address

	//fmt.Printf("\n%v",personal)
	//fmt.Printf("\n%s",personal[name])

	jPerson, _ := json.Marshal(personal)
	fmt.Printf("\n%s", string(jPerson))

}
