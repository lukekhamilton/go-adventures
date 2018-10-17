package main

import (
	"fmt"
	"os"
	"strconv"
)

func fact(n int) int {
	fmt.Println("fact: ", n)
	if n == 0 {
		return 1
	}
	return n + fact(n-1)
}

// Recursion ...
func Recursion() {
	fmt.Printf("%T, %v\n", os.Args[1:], os.Args[1:])
	args := os.Args[1]
	if i, err := strconv.Atoi(args); err == nil {
		fmt.Printf("%T, %v\n", i, i)
		fmt.Println(fact(i))
	}
}
