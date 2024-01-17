package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Nadia")
	want := "Hello Nadia!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
