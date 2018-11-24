package main

import (
	"testing"
)

func TestGenDisplaceFn1(t *testing.T) {

	expected := 24.0
	fn := GenDisplaceFn(2.0, 5.0, 0.0)
	got := fn(3)

	// fmt.Println(got)
	if expected != got {
		t.Error("Failed. Got: ", got)
	}
}

func TestGenDisplaceFn2(t *testing.T) {

	expected := 52.0
	fn := GenDisplaceFn(10.0, 2.0, 1.0)
	got := fn(3)

	if expected != got {
		t.Error("Failed. Got: ", got)
	}
}

func TestGenDisplaceFn3(t *testing.T) {

	expected := 136.0
	fn := GenDisplaceFn(10.0, 2.0, 1.0)
	got := fn(5)

	if expected != got {
		t.Error("Failed. Got: ", got)
	}
}
