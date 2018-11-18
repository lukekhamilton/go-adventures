package main

import "fmt"

// func foo(y *int) {
// 	*y = *y + 1
// }

// func passPointer() {

// 	x := 2
// 	foo(&x)
// 	fmt.Println(x)
// }

// func passArray(x [3]]int) int {
// 	return x[3]
// }

func sliceFoo(sli []int) {
	sli[0] = sli[0] + 1
}

func week1() {
	a := []int{1, 2, 3}
	sliceFoo(a)
	fmt.Println(a)
}
