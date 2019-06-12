package main

import (
	"fmt"
)

// Range ...
func Range() {
	fmt.Println("Hello world")

	nums := []int{2, 3, 4}
	sum := 0

	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}

	fmt.Println(kvs)

	for k, v := range kvs {
		// v = "pears"
		kvs[k] = "pears"
		fmt.Printf("%s -> %s\n", k, v)
	}

	fmt.Println(kvs)

	for k := range kvs {
		fmt.Println("key", k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
