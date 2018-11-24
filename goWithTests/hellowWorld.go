package main

import "fmt"

// Hello ...
func Hello(msg string) string {
	return "Hello, " + msg
}

// GoodBye ...
func GoodBye(msg string) string {
	msg = fmt.Sprintf("Goodbye %s, sorry to see you leave!", msg)
	return msg
}

func main() {
	fmt.Println(Hello("Luke"))
}
