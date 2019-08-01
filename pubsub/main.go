package main

import (
	"fmt"
	"log"

	"github.com/cskr/pubsub"
)

const topic = "data"

func publish(ps *pubsub.PubSub) {
	log.Println("publishing")

	for {
		ps.Pub("Message", topic)
	}
}

func main() {
	ps := pubsub.New(0)
	ch := ps.Sub(topic)
	go publish(ps)

	log.Println("subscripting")
	for i := 1; ; i++ {
		if i == 5 {
			go ps.Unsub(ch, topic)
		}

		if msg, ok := <-ch; ok {
			fmt.Printf("Received %s, %d times.\n", msg, i)
		} else {
			break
		}
	}
}
