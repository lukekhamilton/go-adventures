package main

import (
	"encoding/json"
	"fmt"
)

// Person ...
type Person struct {
	name  string
	addr  string
	phone string
}

func main() {
	p1 := Person{name: "joe", addr: "st 1", phone: "12345"}
	// p2 := Person{"joe", "st 1", "12345"}

	barr, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	fmt.Printf("%+v\n", barr)
}
