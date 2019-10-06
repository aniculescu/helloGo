package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("English greeting a person by name", func(t *testing.T) {
		got := Hello("name")
		want := "Greeting, name"

		assertCorrectMessage(t, got, want)
	})

	t.Run("English greeting with default value", func(t *testing.T) {
		got := Hello("")
		want := "Greeting, default"

		assertCorrectMessage(t, got, want)
	})
}