package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/sacOO7/gowebsocket"
)

const (
	conn = "wss://api.bitfinex.com/ws/2"
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
		log.Println("Recieved connect error: ", err)
	}

	socket.OnTextMessage = func(message string, socket gowebsocket.Socket) {
		log.Println("Recieved message: " + message)
	}

	socket.OnBinaryMessage = func(data []byte, socket gowebsocket.Socket) {
		log.Println("Recieved binary data: ", data)
	}

	socket.OnPingReceived = func(data string, socket gowebsocket.Socket) {
		log.Println("Recieved ping: " + data)
	}

	socket.OnPongReceived = func(data string, socket gowebsocket.Socket) {
		log.Println("Recieved pong: " + data)
	}

	socket.OnDisconnected = func(err error, socket gowebsocket.Socket) {
		log.Println("Disconnected from server ")
		return
	}

	// socket.EnableLogging()
	socket.Connect()
	socket.SendText(`{ "event": "subscribe",  "channel": "ticker",  "symbol": "tBTCUSD" }`)

	for {
		select {
		case <-interrupt:
			log.Println("Interrupted. Shuting down connection")
			socket.Close()
			return
		}
	}
}
