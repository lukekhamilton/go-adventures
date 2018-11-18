package main

import (
	"fmt"
	"strconv"
)

var channel = make(chan string)
var counter int

func hello(msg string) {
	counter++

	// time.Sleep(time.Millisecond * 500)
	channel <- msg + " " + strconv.Itoa(counter)
	return
}

func main() {

	var x = hello

	for {
		// fmt.Println("Starting...")
		go x("Hello World ThreadX")
		go x("Hello World ThreadY")
		go x("Hello World ThreadZ")
		fmt.Println(<-channel)
	}

	// go func(msg string) {
	// 	channel <- msg
	// }("Hello Threads")

}
