package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is passed", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("says Hello in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("says Hello in French", func(t *testing.T) {
		got := Hello("Francois", "French")
		want := "Bonjour, Francois"
		assertCorrectMessage(t, got, want)
	})

	t.Run("says Hello in Japanese", func(t *testing.T) {
		got := Hello("Yuki", "Japanese")
		want := "Konnichiwa, Yuki"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper() // Informs compilter this is a helper, error messages now report func error line rather than this func
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

// go test runs the tests
// test files must be named as s_test.go
// test functions must start with the word Test
// These functions take one argument only: t *testing.T, and to use testing we need to import it
// t.Errorf prints out a message during the test
