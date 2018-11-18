package gosandpit

import "fmt"

// SliceAdv ...
func SliceAdv() {
	arr := [...]string{"a", "b", "c", "d", "e", "f", "g"}
	fmt.Printf("%+v\n", arr)
	fmt.Printf("%T\n", arr)

	s1 := arr[1:3]
	s2 := arr[2:5]

	fmt.Printf("%+v\n", s1)
	fmt.Printf("%+v\n", s2)

	fmt.Printf("%+v\n", len(s1))
	fmt.Printf("%+v\n", cap(s1))
	fmt.Printf("%+v\n", cap(arr))

	// arr = append(arr, "h")
	fmt.Printf("%+v\n", arr)

	// slice literals
	sli := []int{1, 2, 3}
	fmt.Printf("%T\n", sli)
	fmt.Printf("%v\n", sli)

	sli = append(sli, 4)
	fmt.Printf("%v\n", sli)

	// make slice with make
	sli1 := make([]int, 10)
	fmt.Printf("%+v\n", sli1)

	sli1 = append(sli1, 1)
	fmt.Printf("%+v\n", sli1)

	sli2 := make([]int, 10, 15)
	fmt.Printf("%+v\n", sli2)

	sli3 := make([]int, 0, 3)
	sli3 = append(sli3, 100)
	sli3 = append(sli3, 101)
	sli3 = append(sli3, 102)
	sli3 = append(sli3, 103)
	fmt.Println(sli3)

}
