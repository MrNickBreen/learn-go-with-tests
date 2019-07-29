package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Nicko")
	want := "Hello, Nicko"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
