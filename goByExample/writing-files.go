package main

import (
	"io/ioutil"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	d1 := byte("hello\ngo\n")
	err := ioutil.WriteFile("here.txt", d1, 0644)
	checkErr(err)

}
