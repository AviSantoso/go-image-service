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

func StorageCanReturnWhetherOrNotAnItemExists(t *testing.T, impl StorageInterface) {
	pathFail := "Goodbye/World"
	pathSuccess := "Hello, world!"
	item := []byte("Hello, world!")

	exists, err := impl.Has(pathSuccess)
	if err != nil {
		t.Error(err.Error())
	}
	if exists {
		t.Error("The item should not exist yet.")
	}

	err = impl.Store(pathSuccess, item)
	if err != nil {
		t.Error(err.Error())
	}

	exists, err = impl.Has(pathSuccess)
	if err != nil {
		t.Error(err.Error())
	}
	if !exists {
		t.Error("The item should exist.")
	}

	exists, err = impl.Has(pathFail)
	if err != nil {
		t.Error(err.Error())
	}
	if exists {
		t.Error("The item should not exist.")
	}
}

func StorageShouldReturnErrorWhenRetrievingNonExistentItems(t *testing.T, impl StorageInterface) {
	path := "Hello/World"

	item, err := impl.Retrieve(path)
	if err == nil {
		t.Error("This should return an error since the item does not exist.")
	}
	if item != nil {
		t.Error("This should return an empty item since the item does not exist.")
	}
}

func StorageShouldBeAbleToDeleteExistingItemsAndReturnErrorWhenDeletingNonExistentItems(t *testing.T, impl StorageInterface) {
	path := "Hello, world!"
	item := []byte("Hello, world!")

	err := impl.Delete(path)
	if err == nil {
		t.Error("The item should not exist yet.")
	}

	err = impl.Store(path, item)
	if err != nil {
		t.Error(err.Error())
	}

	err = impl.Delete(path)
	if err != nil {
		t.Error(err.Error())
	}

	err = impl.Delete(path)
	if err == nil {
		t.Error("This item should already be deleted and should not exist.")
	}

}
