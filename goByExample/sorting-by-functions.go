package gobyexample

import (
	"fmt"
	"sort"
)

type byLength []string

func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

// SortingByFunctions ...
func SortingByFunctions() {

	fruits := []string{"peace", "banana", "kiwi"}
	fmt.Println(fruits)

	sort.Strings(fruits)
	fmt.Println(fruits)

	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}
