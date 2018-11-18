package pkg

import "fmt"

// X ...
var X = "Hello World"

// Main will print stuff to stdout
func Main() {
	fmt.Printf("%+v\n", X)
}

// Hello will return a string
func Hello() string {
	return X
}
