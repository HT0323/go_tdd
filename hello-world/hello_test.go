package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Masa", "")
		want := "Hello, Masa"

		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello, World", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Aki", "Spanish")
		want := "Hola, Aki"
		assertCorrectMessage(t, got, want)
	})

}
