package gobyexample

import (
	"fmt"
)

func Make() {
	fmt.Println("go make keyword")

	x := make([]int, 0, 10)
	fmt.Println("x:", x)
	fmt.Printf("x: %v\n", x)
	fmt.Printf("type: %T", x)
}
