package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	m := make(map[string]string)

	fmt.Print("Enter name: ")
	scanner.Scan()
	m["name"] = scanner.Text()

	// fmt.Printf("name: `%v`\n", m["name"])

	fmt.Print("Ender address: ")
	scanner.Scan()
	m["addr"] = scanner.Text()

	// fmt.Printf("addr: `%v`\n", m["addr"])

}
