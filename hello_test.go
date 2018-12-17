package main

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("Need %s but got %s", want, got)
		}
	}
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("")
		want := "Hello, world"
		assertCorrectMessage(t, want, got)
	})

	t.Run("say 'Hello, World', when empty string", func(t *testing.T) {
		got := Hello("Pat")
		want := "Hello, Pat"
		assertCorrectMessage(t, want, got)
	})
}
