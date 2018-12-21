package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/sacOO7/gowebsocket"
)

const (
	conn = "ws://stream.binance.com:9443"
)

func main() {
	log.Println("Starting up...")
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	socket := gowebsocket.New(conn)

	socket.OnConnected = func(socket gowebsocket.Socket) {
		log.Println("Connected to server")
	}

	socket.OnConnectError = func(err error, socket gowebsocket.Socket) {
		log.Println("Recieved connect error ", err)
	}

	socket.EnableLogging()
	socket.Connect()

	for {
		select {
		case <-interrupt:
			log.Println("Interrupted. Shuting down connection")
			socket.Close()
			return
		}
	}
}
