package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// BubbleSort ...
func BubbleSort(i []int) []int {
	fmt.Println("------------------------------------------------------")
	fmt.Println("Input Data: ", i)

	sortMe := true

	for sortMe {
		sortMe = false

		for index := 0; index < len(i)-1; index++ {

			if i[index] > i[index+1] {
				Swap(i, index)
				sortMe = true
			}
		}
	}

	fmt.Println("Sorted Data: ", i)
	fmt.Println("------------------------------------------------------")
	return i
}

// Swap ...
func Swap(s []int, i int) {
	s[i], s[i+1] = s[i+1], s[i]
}

func inputData() []int {
	var data []int
	s := bufio.NewScanner(os.Stdin)

	fmt.Println("Please enter some ints then press enter!")
	fmt.Print(">")

	s.Scan()
	input := s.Text()

	split := strings.Split(input, " ")

	for _, e := range split {
		i, err := strconv.Atoi(e)
		if err != nil {
			fmt.Printf("Dude `%s` isn't a int!\n", e)
		} else {
			data = append(data, i)
		}
	}
	return data
}

func main() {
	data := inputData()
	BubbleSort(data)
}
