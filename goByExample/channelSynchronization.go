package main

import (
	"fmt"
	"time"
)

func worker(done chan bool, num int) {
	fmt.Print("working...:", num)
	time.Sleep(time.Second * 5)
	fmt.Println("finished")

	done <- true
}

func ChannelSync() {
	done := make(chan bool, 1)
	go worker(done, 1)
	go worker(done, 2)
	fmt.Println("here")
	<-done
}
