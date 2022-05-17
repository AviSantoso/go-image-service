package storage

import "errors"

type InMemoryStorage struct {
	items map[string][]byte
}

func NewInMemoryStorage() InMemoryStorage {
	storage := InMemoryStorage{items: map[string][]byte{}}
	return storage
}

func (storage InMemoryStorage) Has(path string) (bool, error) {
	return storage.items[path] != nil, nil
}

func (storage InMemoryStorage) Store(path string, data []byte) error {
	if storage.items[path] != nil {
		return errors.New("The item with path '" + path + "' already exists.")
	}
	storage.items[path] = data
	return nil
}

func (storage InMemoryStorage) Retrieve(path string) ([]byte, error) {
	if storage.items[path] == nil {
		return nil, errors.New("The item with path '" + path + "' does not exist.")
	}
	return storage.items[path], nil
}

func (storage InMemoryStorage) Delete(path string) error {
	if storage.items[path] == nil {
		return errors.New("The item with path '" + path + "' does not exist.")
	}
	storage.items[path] = nil
	return nil
}
