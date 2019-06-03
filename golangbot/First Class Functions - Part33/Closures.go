package main

import "fmt"

func simple() {
	a := 5
	func() {
		fmt.Println("a: ", a)
	}()
}

func appendString() func(string) string {
	t := "Hello"
	return func(b string) string {
		t = t + " " + b
		return t
	}
}

func callAppendString() {
	a := appendString()
	b := appendString()

	fmt.Println(a("World"))
	fmt.Println(b("Everyone"))

	fmt.Println(a("Gopher"))
	fmt.Println(b("!"))

}

func main() {
	callAppendString()
}
