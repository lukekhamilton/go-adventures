package main

import "fmt"

// Channels ...
func Channels() {
	messages := make(chan string)

	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg) // Prints chan
}
