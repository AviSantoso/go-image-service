package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

/**
 * We expose this route so that we can easily check the health status of the
 * instance by accessing any / route outside of the api routes.
 */
func TestHelloWorld(t *testing.T) {
	// For example, let's send a request to '/hello'
	engine := getEngine()
	req, _ := http.NewRequest("GET", "/", nil)

	w := httptest.NewRecorder()
	engine.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status to be 200 OK, but received %d instead.", w.Code)
	}

	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Errorf(err.Error())
	}

	// We expect the response to be 'Hello, there!'. This tells us the server
	// is listening, port bound, and accessible from the developer's machine.
	expected := "Hello, there!"
	if !strings.Contains(string(body), expected) {
		t.Errorf("Expected body to contain '%s', but received %s instead.", expected, body)
	}
}
