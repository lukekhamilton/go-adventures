package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	data := []int{0, 8, 7, 6, 5, 4, 3, 2, 1, 9}
	wanted := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	got := BubbleSort(data)

	fmt.Println("wanted: ", wanted)
	fmt.Println("got: ", got)

	// for index := 0; index < len(got); index++ {
	// 	if len(got) != len(wanted) {
	// 		t.Error("Incorrect length")
	// 	}

	// 	if got[index] != wanted[index] {
	// 		t.Errorf("Slice index `%v` got `%v` and wanted `%v`", index, got[index], wanted[index])
	// 	}
	// }

	if reflect.DeepEqual(got, wanted) == false {
		t.Error("Slices not the same")
	}
}
