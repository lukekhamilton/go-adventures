package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	slice := make([]int, 0, 3)
	msg := "Please enter an int (X to exit): "

	fmt.Printf(msg)
	// input := "luke was here"
	scanner := bufio.NewScanner(os.Stdin)
	// scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		input := scanner.Text()
		if input == "X" {
			break
		}

		x, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("No a valid int")
		} else {
			slice = append(slice, x)
		}
		sort.Ints(slice)
		fmt.Printf("%+v\n", slice)
		fmt.Printf(msg)
	}
	fmt.Printf("%+v\n", scanner.Text())
}
