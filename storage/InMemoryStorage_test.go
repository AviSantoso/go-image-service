package storage

import (
	"testing"
)

func TestInMemoryStorageCanStoreAndRetrieveAnItem(t *testing.T) {
	storage := NewInMemoryStorage()
	StorageCanStoreAndRetrieveAnItem(t, storage)
}
