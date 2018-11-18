package gobyexample

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
		time.Sleep(100 * time.Millisecond)
	}
}

// Goroutines ...
func Goroutines() {

	for x := 0; x < 5; x++ {
		go f("goroutine1")

		go f("goroutine2")
	}
	go f("goroutine3")

	go f("goroutine4")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	fmt.Scanln()
	fmt.Println("done")
}
