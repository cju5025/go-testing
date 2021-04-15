package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	got := Hello("Colter")
	want := "Hello, Colter"

	if got != want {
		t.Errorf("got %q want %q,", got, want)
	}
}
