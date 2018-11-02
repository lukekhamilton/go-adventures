package main

import "fmt"

func main() {
	channel := make(chan string)

	go func(msg string) {
		channel <- msg
	}("Hello Threads")

	fmt.Println(<-channel)

}
