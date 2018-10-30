package main

import (
	"fmt"
)

// Findian ...
func Findian() {
	var str string
	fmt.Scan(&str)
	fmt.Println(str)
}

func XXX() {
	var xtemp int
	x1 := 0
	x2 := 1
	for x := 0; x < 5; x++ {
		xtemp = x2
		x2 = x2 + x1
		x1 = xtemp
	}
	fmt.Println(x2)
}
