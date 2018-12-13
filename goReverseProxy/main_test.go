package main

import "testing"

func TestGetListenAddress(t *testing.T) {
	got := getListenAddress()
	want := ":1338"
	if got != want {
		t.Errorf("got: %s, want: %s", got, want)
	}
}

func TestGetEnv(t *testing.T) {
	t.Run("Fallback", func(t *testing.T) {
		got := getEnv("TESTING", "FAIL")
		want := "FAIL"
		if got != want {
			t.Error("Failed dude!")
		}
	})

	t.Run("Success", func(t *testing.T) {
		got := getEnv("USER", "FAIL")
		want := "metta"
		if got != want {
			t.Errorf("Failed dude! got: %s, want: %s ", got, want)
		}
	})
}
