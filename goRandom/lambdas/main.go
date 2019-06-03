package main

import "fmt"

func main() {
	addOne := func(x int) int {
		return x + 1
	}

	fmt.Println("addOne: ", addOne(5))

	// use lambda as a value
	v := []int{1, 2, 3, 4, 5}
	apply(v, addOne)

	fmt.Println("v: ", v)
}

func apply(arr []int, f func(int) int) {
	for i, a := range arr {
		arr[i] = f(a)
	}
}
