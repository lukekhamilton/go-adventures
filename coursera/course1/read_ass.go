package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Person struct {
	fname string
	lname string
}

func main() {
	fmt.Print("Enter the filename: ")
	var filename string
	_, _ = fmt.Scan(&filename)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()

	var persons = make([]Person, 0, 5)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var fname, lname string
		fmt.Sscanf(scanner.Text(), "%s %s", &fname, &lname)
		persons = append(persons, Person{fname: fname, lname: lname})
	}
	for _, person := range persons {
		fmt.Printf("%s %s\n", person.fname, person.lname)
	}
}
