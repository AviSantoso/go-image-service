package storage

import (
	"bytes"
	"testing"
)

func TestInMemoryStorageRetrieve(t *testing.T) {
	var output bytes.Buffer
	storage := NewInMemoryStorage(&output)
	T_StorageInterfaceRetrieve(t, storage)
}

func TestInMemoryStorageStore(t *testing.T) {
	var output bytes.Buffer
	storage := NewInMemoryStorage(&output)
	T_StorageInterfaceStore(t, storage)
}

func TestInMemoryStorageHas(t *testing.T) {
	var output bytes.Buffer
	storage := NewInMemoryStorage(&output)
	T_StorageInterfaceHas(t, storage)
}

func TestInMemoryStorageDelete(t *testing.T) {
	var output bytes.Buffer
	storage := NewInMemoryStorage(&output)
	T_StorageInterfaceDelete(t, storage)
}
