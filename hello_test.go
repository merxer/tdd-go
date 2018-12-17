package main

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	result := Hello("world")
	expected := "Hello, world"

	if result != expected {
		t.Errorf("Need %s but got %s", expected, result)
	}
}

func TestHelloWithParameter(t *testing.T) {
	got := Hello("Pat")
	want := "Hello, Pat"

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
