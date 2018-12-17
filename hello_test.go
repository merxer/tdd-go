package main

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	result := Hello()
	expected := "Hello, world"

	if result != expected {
		t.Errorf(`Need ${expected} but got ${result}`)
	}
}
