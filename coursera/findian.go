package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func ProcessINPUT(rd io.Reader) string {
	var found = false

	fmt.Print("Enter data: ")
	in := bufio.NewReader(rd)

	line, err := in.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	line = strings.Trim(line, "\n")
	line = strings.ToLower(line)
	fmt.Printf("Data: '%v'\n", line)

	if strings.HasPrefix(line, "i") && strings.HasSuffix(line, "n") && strings.Contains(line, "a") {
		found = true
	}

	if found {
		fmt.Printf("%+v\n", "Found!")
		return "Found!"
	}

	fmt.Printf("%+v\n", "Not Found!")
	return "Not Found!"
}

// Main ...
func main() {
	for {
		ProcessINPUT(os.Stdin)
	}
}
