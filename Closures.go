package main

import "fmt"

func intSeq() func() int {
	i := 0

	return func() int {
		i++
		return i
	}
}

func Closures() {
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// Return's a memory address
	// fmt.Println(intSeq())
	// fmt.Println(intSeq())

	newInts := intSeq()
	fmt.Println(newInts())
	fmt.Println(newInts())
	fmt.Println(newInts())
}
