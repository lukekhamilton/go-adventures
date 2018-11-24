package main

import "testing"

func check(got string, want string, t *testing.T) {
	if got != want {
		t.Errorf("got: '%s' want: '%s'", got, want)
	}
}

func TestHello(t *testing.T) {
	got := Hello("Chris")
	want := "Hello, Chris"
	check(got, want, t)
}

func TestGoodBye(t *testing.T) {
	got := GoodBye("Luke")
	want := "Goodbye Luke, sorry to see you leave!"
	check(got, want, t)
}
