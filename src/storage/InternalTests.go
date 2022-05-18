package storage

import (
	"testing"
)

func isByteArrEqual(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func StorageCanStoreAndRetrieveAnItem(t *testing.T, impl StorageInterface) {
	path := "Hello/World"
	item := []byte("Hello, world!")

	err := impl.Store(path, item)
	if err != nil {
		t.Error(err.Error())
	}

	received, err := impl.Retrieve(path)
	if err != nil {
		t.Error(err.Error())
	}

	if !isByteArrEqual(item, received) {
		t.Error("Failed to store and retrieve the same item.")
	}
}

func StorageThrowsWhenAnItemPathAlreadyExists(t *testing.T, impl StorageInterface) {
	path := "Hello/World"
	item := []byte("Hello, world!")

	err := impl.Store(path, item)
	if err != nil {
		t.Error(err.Error())
	}

	err = impl.Store(path, item)
	if err == nil {
		t.Error("Should fail to store an item to the same path.")
	}
}
