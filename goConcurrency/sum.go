package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateInts() {
	s1 := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(s1)
	aa := make(chan int)
	bb := make(chan int)
	cc := make(chan int)

	go sum(aa, bb, cc)

	for {
		aa <- rand.Int()
		bb <- rand.Int()
		cc <- rand.Int()
	}
}

func sum(a chan int, b chan int, c chan int) {
	for i := 0; ; i++ {
		aa, bb, cc := <-a, <-b, <-c
		sum := aa + bb + cc
		fmt.Printf("sum[%v]: %v\n", i, sum)
	}
}

func main() {
	generateInts()
}
