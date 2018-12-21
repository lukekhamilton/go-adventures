package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my website")
		log.Print("hit from: %v!", r)
	})

	fs := http.FileServer(http.Dir("static"))
	http.Handle("static", http.StripPrefix("static", fs))

	log.Print("Started Server! 127.0.0.1:80")

	http.ListenAndServe(":80", nil)
}
