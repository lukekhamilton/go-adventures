package main

import "fmt"

func callStyle1() {
	a := func() {
		fmt.Println("hello world first class function")
	}

	a()
	fmt.Printf("%T", a)
}

func callStyle2() {
	func() {
		fmt.Println("hello world first class function")
	}()
}

func passArgs() {
	func(n string) {
		fmt.Println("Welcome ", n)
	}("Gophers")
}

func main() {
	passArgs()
}
