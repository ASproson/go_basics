package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Chris")
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want) // %q interpolates the passed variables
	}
}

// go test runs the tests
// test files must be named as s_test.go
// test functions must start with the word Test
// These functions take one argument only: t *testing.T, and to use testing we need to import it
// t.Errorf prints out a message during the test
