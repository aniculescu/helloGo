package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("name")
	want := "hello name"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
