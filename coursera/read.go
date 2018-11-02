package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type name struct {
	fname string
	lname string
}

var names []name

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Please enter a filename: ")
	scanner.Scan()
	filename := scanner.Text()

	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	fscanner := bufio.NewScanner(file)

	for fscanner.Scan() {
		line := fscanner.Text()
		sline := strings.Split(line, " ")
		names = append(names, name{sline[0], sline[1]})
	}

	for _, name := range names {
		fmt.Printf("fname: %s, \t lname: %s\n", name.fname, name.lname)
	}

}
