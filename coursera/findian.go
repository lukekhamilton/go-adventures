package coursera

import (
	"fmt"
	"strings"
)

var searchChars = []string{"i", "a", "n"}

const (
	i = "i"
	a = "a"
	n = "n"
)

// Main ...
func Main() {
	var str string
	var found = false

	fmt.Print("Enter data: ")
	fmt.Scanln(&str)
	fmt.Printf("Data: %+v\n", str)

	if strings.HasPrefix(str, "i") {
		found = true
	}

	if strings.HasSuffix(str, "n") {
		found = true
	}

	if strings.Contains(str, "a") {
		found = true
	}

	if found {
		fmt.Printf("%+v\n", "Found")
	}
}
