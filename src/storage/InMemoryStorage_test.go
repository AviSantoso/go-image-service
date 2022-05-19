package storage

import (
	"testing"
)

func TestInMemoryStorageRetrieve(t *testing.T) {
	storage := NewInMemoryStorage()
	T_StorageInterfaceRetrieve(t, storage)
}

func TestInMemoryStorageStore(t *testing.T) {
	storage := NewInMemoryStorage()
	T_StorageInterfaceStore(t, storage)
}

func TestInMemoryStorageHas(t *testing.T) {
	storage := NewInMemoryStorage()
	T_StorageInterfaceHas(t, storage)
}

func TestInMemoryStorageDelete(t *testing.T) {
	storage := NewInMemoryStorage()
	T_StorageInterfaceDelete(t, storage)
}
