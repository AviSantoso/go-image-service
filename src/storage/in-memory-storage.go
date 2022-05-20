package storage

import (
	"errors"
	"fmt"
	"io"

	"github.com/AviSantoso/go-image-service/logger"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

type InMemoryStorage struct {
	items  map[string][]byte
	logger logger.Logger
}

func NewInMemoryStorage(writer io.Writer) InMemoryStorage {
	id, _ := gonanoid.New()
	logger := logger.New(writer, "storage/in-memory-storage", id)
	storage := InMemoryStorage{items: map[string][]byte{}, logger: logger}
	return storage
}

func (storage InMemoryStorage) Has(path string) (bool, error) {
	return storage.items[path] != nil, nil
}

func (storage InMemoryStorage) Store(path string, data []byte) error {
	storage.logger.Info(fmt.Sprintf("Storing data at path %s with length %d", path, len(data)))

	if storage.items[path] != nil {
		return errors.New("The item with path '" + path + "' already exists.")
	}
	storage.items[path] = data
	return nil
}

func (storage InMemoryStorage) Retrieve(path string) ([]byte, error) {
	storage.logger.Info(fmt.Sprintf("Retrieving data at path %s", path))

	if storage.items[path] == nil {
		return nil, errors.New("The item with path '" + path + "' does not exist.")
	}
	return storage.items[path], nil
}

func (storage InMemoryStorage) Delete(path string) error {
	storage.logger.Info(fmt.Sprintf("Deleted data at path %s", path))

	if storage.items[path] == nil {
		return errors.New("The item with path '" + path + "' does not exist.")
	}
	storage.items[path] = nil
	return nil
}
