package main

import "testing"

func TestHello(t *testing.T) {
	name := "World"
	message, err := Hello(name)

	if err != nil {
		t.Error(err.Error())
	}

	if message != "Hello, World!" {
		t.Errorf("Expected 'Hello World!' but got '%s'", message)
	}
}
