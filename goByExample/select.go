package main

import (
	"fmt"
	"time"
)

// Select ...
func Select() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			time.Sleep(1 * time.Second)
			c1 <- "one"
		}
	}()

	go func() {
		for {
			time.Sleep(2 * time.Second)
			c2 <- "two"
		}
	}()

	for {
		select {
		case msg1 := <-c1:
			fmt.Println("received: ", msg1)
		case msg2 := <-c2:
			fmt.Println("received: ", msg2)
		}
	}
}

func main() {
	Select()
}
