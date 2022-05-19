package storage

import (
	"testing"
)

func TestInMemoryStorageCanStoreAndRetrieveAnItem(t *testing.T) {
	storage := NewInMemoryStorage()
	StorageCanStoreAndRetrieveAnItem(t, storage)
}

func TestInMemoryStorageThrowsWhenAnItemPathAlreadyExists(t *testing.T) {
	storage := NewInMemoryStorage()
	StorageThrowsWhenAnItemPathAlreadyExists(t, storage)
}

func TestInMemoryStorageCanReturnWhetherOrNotAnItemExists(t *testing.T) {
	storage := NewInMemoryStorage()
	StorageCanReturnWhetherOrNotAnItemExists(t, storage)
}

func TestInMemoryStorageShouldReturnErrorWhenRetrievingNonExistentItems(t *testing.T) {
	storage := NewInMemoryStorage()
	StorageShouldReturnErrorWhenRetrievingNonExistentItems(t, storage)
}

func TestInMemoryStorageShouldBeAbleToDeleteExistingItemsAndReturnErrorWhenDeletingNonExistentItems(t *testing.T) {
	storage := NewInMemoryStorage()
	StorageShouldBeAbleToDeleteExistingItemsAndReturnErrorWhenDeletingNonExistentItems(t, storage)
}
