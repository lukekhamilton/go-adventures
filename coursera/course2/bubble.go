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
	fmt.Println("------------------------------------------------------")
	end := len(i) - 1

	for index := 0; index < len(i)-1; index++ {
		fmt.Println("i: ", i[index])
		fmt.Println("Index: ", index)

		if end == 0 {
			break
		}

		if i[index] > i[index+1] {
			Swap(i, index)
		}

		end--
	}

	fmt.Println("------------------------------------------------------")
}

func Swap(s []int, i int) {
	s[i], s[i+1] = s[i+1], s[i]
}

func main() {
	fmt.Println("Please enter some ints then press enter!")
	fmt.Print(">")
	s := bufio.NewScanner(os.Stdin)
	var data []int

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

	fmt.Println("Raw int Data: ", data)
	BubbleSort(data)
	fmt.Println("Sorted int Data: ", data)
}
