package main

import (
	"fmt"
	"time"
)

var startTime time.Time

func uptime() time.Duration {
	return time.Since(startTime)
}

func init() {
	startTime = time.Now()
}

func main() {
	fmt.Println("started")

	time.Sleep(time.Second * 1)
	fmt.Printf("uptime %s\n", uptime())

	time.Sleep(time.Second * 5)
	fmt.Printf("uptime %s\n", uptime())
}
