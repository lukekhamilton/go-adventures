package main

import "fmt"

func main() {
	var x float32 = 1.234
	var xx float64 = 1.234
	var xxx = 1.234

	fmt.Println(x)
	fmt.Printf("%+T\n", x)

	fmt.Println(xx)
	fmt.Printf("%+T\n", xx)

	fmt.Println(xxx)
	fmt.Printf("%+T\n", xxx)
}
