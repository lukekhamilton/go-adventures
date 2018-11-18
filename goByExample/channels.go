package gobyexample

import "fmt"

// Channels ...
func Channels() {
	messages := make(chan string)

	go func() { messages <- "ping" }()

	// msg := <-messages
	fmt.Println(<-messages) // Prints chan
	fmt.Println(<-messages) // blocks chan
}
