package main

import "testing"

func getEnvTest(t *testing.T) {
	t.Run("Fallback", func(t *testing.T) {
		got := getEnv("TESTING", "FAIL")
		want := "FAIL"
		if got != want {
			t.Error("Failed dude!")
		}
	})

	t.Run("Success", func(t *testing.T) {
		got := getEnv("USER", "FAIL")
		want := "mett12a"
		if got != want {
			t.Errorf("Failed dude! got: %s, want: %s ", got, want)
		}
	})
}
