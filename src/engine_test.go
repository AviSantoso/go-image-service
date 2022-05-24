package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

/**
 * Helper function to upload a multipart file. You can use this function in your
 * client code to upload a multipart file to this server.
 */
func UploadMultipartImage(uri string, key string, file_path string) (*http.Request, error) {
	body, writer := io.Pipe()

	req, err := http.NewRequest(http.MethodPost, uri, body)
	if err != nil {
		return nil, err
	}

	multipart_writer := multipart.NewWriter(writer)
	req.Header.Add("Content-Type", multipart_writer.FormDataContentType())

	err_channel := make(chan error)

	go func() {
		defer close(err_channel)
		defer writer.Close()
		defer multipart_writer.Close()

		w, err := multipart_writer.CreateFormFile(key, file_path)
		if err != nil {
			err_channel <- err
			return
		}

		in, err := os.Open(file_path)
		if err != nil {
			err_channel <- err
			return
		}
		defer in.Close()

		if written, err := io.Copy(w, in); err != nil {
			err_channel <- fmt.Errorf("error copying %s (%d bytes written): %v", file_path, written, err)
			return
		}

		if err := multipart_writer.Close(); err != nil {
			err_channel <- err
			return
		}
	}()

	return req, nil
}

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

/**
 * This test describes the steps that should be taken in order to correctly upload
 * and download image files to and from the server.
 */
func TestUploadDownloadImage(t *testing.T) {
	engine := getEngine()

	new_id, err := gonanoid.New()
	if err != nil {
		t.Errorf(err.Error())
	}

	// We send a POST request to this route with a multipart image as the body.
	// Ensure that the content type is multipart/form-data at the key 'file'.
	route := fmt.Sprintf("/api/v1/image/upload/%s", new_id)
	req, _ := UploadMultipartImage(route, "file", "./assets/test.png")

	w := httptest.NewRecorder()
	engine.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status to be 200 OK, but received %d instead.", w.Code)
	}

	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Errorf(err.Error())
	}

	var data map[string]interface{}
	err = json.Unmarshal([]byte(body), &data)
	if err != nil {
		t.Errorf(err.Error())
	}

	// We expect the received ID to be the same ID that we specified in the URI.
	id, ok := data["id"].(string)
	if !ok {
		t.Errorf("Expected id to be a string, but got %v", data["id"])
	}
	if id != new_id {
		t.Errorf("Expected the response id to be %s but received %s instead.", new_id, id)
	}
}
