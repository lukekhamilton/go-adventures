package main

import "fmt"

// Passing functions as arguments to other functions
func simple1(a func(a, b int) int) {
	fmt.Println(a(60, 7))
}

func callSimple1() {
	f := func(a, b int) int {
		return a + b
	}

	simple1(f)
}

// Returning functions from other functions
func simple2() func(a, b int) int {
	return func(a, b int) int {
		return a + b
	}
}

func callSimple2() {
	s := simple2()
	fmt.Println(s(60, 7))
}

func main() {
	callSimple2()
}
