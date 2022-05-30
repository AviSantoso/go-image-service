package main

import "testing"

func TestHello(t *testing.T) {
	name := "World"
	message := Hello(name)

	if message != "Hello, World!" {
		t.Errorf("Expected 'Hello World!' but got '%s'", message)
	}
}
