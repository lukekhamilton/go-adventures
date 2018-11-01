package main

import "fmt"

func ex1() {
	x := []int{4, 8, 5}
	y := -1
	for _, elt := range x {
		if elt > y {
			y = elt
		}
	}
	fmt.Print(y)
}

func ex2() {
	x := [...]int{4, 8, 5}
	y := x[0:2]
	z := x[1:3]
	fmt.Print(z)
	y[0] = 1
	z[1] = 3
	fmt.Print(x)
}

func ex3() {
	x := [...]int{1, 2, 3, 4, 5}
	y := x[0:2]
	fmt.Println(y)
	z := x[1:4]
	fmt.Println(z)
	fmt.Print(len(y), cap(y), len(z), cap(z))
}

func ex4() {
	x := map[string]int{
		"ian": 1, "harris": 2}
	for i, j := range x {
		if i == "harris" {
			fmt.Print(i, j)
		}
	}
}

type P struct {
	x string
	y int
}

func main() {
	b := P{"x", -1}

	a := [...]P{P{"a", 10}, P{"b", 2}, P{"c", 3}}
	fmt.Printf("a:%v", a)
	fmt.Printf("a:%T\n", a)
	for _, z := range a {
		fmt.Printf("z:%v\n", z)
		if z.y > b.y {
			b = z
		}
	}
	fmt.Println(b.x)
}
