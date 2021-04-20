package main

import (
	"testing"
)

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got string, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {

		got := Hello("Colter", english)
		want := "Hello, Colter"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", english)
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Amber", spanish)
		want := "Hola, Amber"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Amber", french)
		want := "Bonjour, Amber"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Italian", func(t *testing.T) {
		got := Hello("Amber", italian)
		want := "Ciao, Amber"
		assertCorrectMessage(t, got, want)
	})
}
