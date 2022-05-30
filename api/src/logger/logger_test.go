package logger

import (
	"bytes"
	"strings"
	"testing"
)

func TestLogger(t *testing.T) {
	var output bytes.Buffer
	path := "MyPath"
	id := "MyId"
	logger := New(&output, path, id)

	message := "Hello, world!"
	level := "info"
	logger.Info(message)

	res := output.String()

	if !strings.Contains(res, message) {
		t.Errorf("Expected '%s' to contain '%s'.", res, message)
	}

	if !strings.Contains(res, level) {
		t.Errorf("Expected '%s' to contain '%s'.", res, level)
	}

	output = bytes.Buffer{}

	message = "Armageddon!"
	level = "error"
	logger.Error(message)

	res = output.String()

	if !strings.Contains(res, message) {
		t.Errorf("Expected '%s' to contain '%s'.", res, message)
	}

	if !strings.Contains(res, level) {
		t.Errorf("Expected '%s' to contain '%s'.", res, level)
	}
}

