package gosandpit

import "fmt"

func send() {
	// messages <- "ping"
}

func receive() {

}

// ChannelAdv ...
func ChannelAdv() {
	messages := make(chan string)

	// go func() { messages <- "ping" }()

	go send()
	go receive()

	msg := <-messages
	fmt.Println(msg)
}
