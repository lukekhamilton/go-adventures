package main

import (
	"fmt"
)

func main() {
	fmt.Println("Please enter a sequence of 10 integers: ")
	seq := make([]int, 0, 10)
	var num int
	for i := 0; i < 10; i++ {
		fmt.Scan(&num)
		seq = append(seq, num)
	}
	fmt.Println("The sequence entered is: ")
	fmt.Println(seq)
	BubbleSort(seq)
	fmt.Println("Sequence after sorting:")
	fmt.Println(seq)
}

func BubbleSort(seq []int) {
	var swapped bool = true
	for swapped {
		swapped = false
		for j := 0; j < len(seq)-1; j++ {
			if seq[j] > seq[j+1] {
				swap(seq, j)
				swapped = true
			}
		}
		if !swapped {
			swapped = false
		}
	}
}

func swap(seq []int, i int) {
	var temp int
	temp = seq[i]
	seq[i] = seq[i+1]
	seq[i+1] = temp
}
