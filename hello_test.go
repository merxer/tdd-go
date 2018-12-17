package main

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		result := Hello("")
		expected := "Hello, world"

		if result != expected {
			t.Errorf("Need %s but got %s", expected, result)
		}
	})

	t.Run("say 'Hello, World', when empty string", func(t *testing.T) {
		got := Hello("Pat")
		want := "Hello, Pat"

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
}

func TestHelloWithParameter(t *testing.T) {

}
