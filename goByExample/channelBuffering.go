package gobyexample

import "fmt"

func ChannelBuffering() {

	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages) // Prints chan
	fmt.Println(<-messages) // Prints chan
	fmt.Println(<-messages) // blocks chan
}
