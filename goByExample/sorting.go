package gobyexample

import (
	"fmt"
	"sort"
)

// Sorting ...
func Sorting() {
	strs := []string{"c", "a", "b"}
	fmt.Println("Strings: ", strs)
	fmt.Printf("%T\n", strs)
	sort.Strings(strs)
	fmt.Println("Strings: ", strs)

	ints := []int{7, 2, 4}
	fmt.Println("Ints: ", ints)
	sort.Ints(ints)
	fmt.Println("Ints: ", ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)
}
