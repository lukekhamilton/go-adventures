package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// BubbleSort ...
func BubbleSort(i []int) {

}

func main() {
	fmt.Println("Please enter up to 10 ints then Press enter")
	fmt.Print(">")
	s := bufio.NewScanner(os.Stdin)
	var data []int

	s.Scan()
	input := s.Text()

	split := strings.Split(input, " ")

	for _, e := range split {
		i, err := strconv.Atoi(e)
		if err != nil {
			fmt.Println("Err: ", err)
		} else {
			fmt.Println(i)
			data = append(data, i)
		}
	}

	fmt.Println("Raw Data: ", data)

}
