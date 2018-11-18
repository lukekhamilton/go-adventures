package gobyexample

import (
	"fmt"
	"time"
)

func worker(done chan bool, num int) {
	fmt.Printf("worker:%v running\n", num)
	time.Sleep(time.Second * 5)
	fmt.Printf("worker:%v finished\n", num)
	done <- true
}

// ChannelSync ...
func ChannelSync() {
	done := make(chan bool, 1)
	go worker(done, 1)
	go worker(done, 2)
	fmt.Println("main running")
	<-done
}
