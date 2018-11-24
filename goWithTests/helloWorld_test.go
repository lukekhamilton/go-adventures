package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Chris")
	want := "Hello, Chris"
	if got != want {
		t.Errorf("got: '%s' want: '%s'", got, want)
	}
}

func TestHelloNext(t *testing.T) {
	assertCorrectMessage := func(got, want string, t *testing.T) {
		t.Helper()
		if got != want {
			t.Errorf("got: '%s' want: '%s'", got, want)
		}
	}

	t.Run("saging Hello to people", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"
		assertCorrectMessage(got, want, t)
	})

	t.Run("say Hello, world when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		assertCorrectMessage(got, want, t)
	})
}

func TestGoodBye(t *testing.T) {
	got := GoodBye("Luke")
	want := "Goodbye Luke, sorry to see you leave!"
	if got != want {
		t.Errorf("got: '%s' want: '%s'", got, want)
	}
}
