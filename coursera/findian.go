package main

import (
	"fmt"
)

var xxx = 123

// Findian ...
func Findian() {
	var str string
	fmt.Scan(&str)
	fmt.Println(str)
}

func XXX() {
	var xtemp int
	fmt.Println(xxx)
	x1 := 0
	x2 := 1
	for x := 0; x < 5; x++ {
		xtemp = x2
		x2 = x2 + x1
		x1 = xtemp
		xxx += x1 + x2 + xtemp
	}
	fmt.Println(x2)
}
