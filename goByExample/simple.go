package gobyexample

import (
	"fmt"
	"log"
	"net/http"
)

// NetPrint ...
func NetPrint(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Hello Gophers!\n")
	fmt.Println("Req: ", req)
}

// SimpleWWW ...
func SimpleWWW() {
	fmt.Println("http server starting on port:12345")
	http.HandleFunc("/", NetPrint)
	err := http.ListenAndServe("localhost:12345", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
