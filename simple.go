package main

import (
	"fmt"
	"log"
	"net/http"
)

func NetPrint(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Hello Gophers!\n")
}

func main() {
	http.HandleFunc("/", NetPrint)
	err := http.ListenAndServe("localhost:12345", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
