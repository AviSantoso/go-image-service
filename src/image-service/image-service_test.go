package imageservice

import (
	"bytes"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/AviSantoso/go-image-service/storage"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

/**
 * Get started with using the ImageService through uploading and downloading
 * the test image within this folder. You can also copy this and try your own.
 */
func TestImageServiceGettingStarted(t *testing.T) {
	var output bytes.Buffer
	var storage = storage.NewInMemoryStorage(&output)
	service := New(&output, storage)

	// A test image is provided in this folder to demo functionality
	img, err := ioutil.ReadFile("test.png")
	if err != nil {
		t.Errorf(err.Error())
	}

	// We need to specify the id the image should be stored at
	id, err := gonanoid.New()
	if err != nil {
		t.Errorf(err.Error())
	}

	// Using the UploadImage method, we can store the image in the service
	res := service.UploadImage(id, img)

	// On a successful upload, the status is 200 OK
	if res.Status != http.StatusOK {
		t.Errorf("Expected status to be OK but got %d", res.Status)
	}

	// The data exposes a field called 'id' that is a string and not empty
	outId, ok := res.Data["id"].(string)
	if !ok || len(outId) == 0 {
		t.Errorf("Expected id to be a string and not be empty.")
	}

	// The id field is the same ID used as input
	if outId != id {
		t.Errorf("Expected outputId to be %s, but got %s", id, outId)
	}

	// The data exposes a field called 'message' that is a string and not empty
	message, ok := res.Data["id"].(string)
	if !ok || len(message) == 0 {
		t.Errorf("Expected message to be a string and not be empty.")
	}

	// At any time later, we can retrieve the image in bytes using the DownloadImage method
	res = service.DownloadImage(id)

	// On a successful download, the status is 200 OK
	if res.Status != http.StatusOK {
		t.Errorf("Expected status to be 200 OK, but got %d", res.Status)
	}

	// The data exposes a field called 'id' that is a string and not empty
	outId, ok = res.Data["id"].(string)
	if !ok || len(outId) == 0 {
		t.Errorf("Expected id to be a string and not be empty.")
	}

	// The id field is the same ID used as input
	if outId != id {
		t.Errorf("Expected outputId to be %s, but got %s", id, outId)
	}

	// The data exposes a field called 'message' that is a string and not empty
	message, ok = res.Data["id"].(string)
	if !ok || len(message) == 0 {
		t.Errorf("Expected message to be a string and not be empty.")
	}

	// The data exposes a field called 'image' that is a byte array and not empty
	image, ok := res.Data["image"].([]byte)
	if !ok || len(image) == 0 {
		t.Errorf("Expected image to be a byte array and not be empty.")
	}
}
