package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("being polite is important, so we do say: ", func(t *testing.T) {
		got := Hello("Elodie", "English")
		want := "Hello Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello World' when an empty string in supplied ", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello World"
		assertCorrectMessage(t, got, want)

	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola Elodie"
		assertCorrectMessage(t, got, want)

	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Elodie", "French")
		want := "Bonjour Elodie"
		assertCorrectMessage(t, got, want)

	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
