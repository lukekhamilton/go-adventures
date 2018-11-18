package gobyexample

import (
	"fmt"
)

// Maps ...
func Maps() {
	fmt.Println("Hello world")

	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("maps: ", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("maps:", m)

	x, err := m["k2"]
	fmt.Println("prs:", x, err)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

}
