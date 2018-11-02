package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	m := make(map[string]string)

	fmt.Print("Enter name: ")
	scanner.Scan()
	m["name"] = scanner.Text()

	fmt.Print("Ender address: ")
	scanner.Scan()
	m["address"] = scanner.Text()

	jsonB, _ := json.Marshal(m)
	fmt.Printf("%+v\n", string(jsonB))

}
