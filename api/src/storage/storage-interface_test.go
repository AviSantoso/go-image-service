package storage

import (
	"testing"
)

/*
	Given that an item exists,
	When the item is retrieved,
	Then the implementation should return the item and a nil error
	And a nil item and an error otherwise
*/
func T_StorageInterfaceRetrieve(t *testing.T, impl StorageInterface) {
	path := "Hello/World"
	item := []byte("Hello, world!")

	err := impl.Store(path, item)
	if err != nil {
		t.Errorf(err.Error())
	}

	received, err := impl.Retrieve(path)
	if err != nil {
		t.Errorf(err.Error())
	}

	if !_isByteArrEqual(item, received) {
		t.Errorf("Failed to store and retrieve the same item.")
	}

	failPath := "Goodbye/World"

	item, err = impl.Retrieve(failPath)
	if err == nil {
		t.Errorf("This should return an error since the item does not exist.")
	}
	if item != nil {
		t.Errorf("This should return an empty item since the item does not exist.")
	}
}

/*
	Given that a path exists,
	When an item is stored at the same path,
	Then the impl should return an error
	And a nil error otherwise
*/
func T_StorageInterfaceStore(t *testing.T, impl StorageInterface) {
	path := "Hello/World"
	item := []byte("Hello, world!")

	err := impl.Store(path, item)
	if err != nil {
		t.Errorf(err.Error())
	}

	err = impl.Store(path, item)
	if err == nil {
		t.Errorf("Should fail to store an item to the same path.")
	}
}

/*
	Given that a path exists,
	When we check for its existence,
	Then the impl should tell us true
	And false otherwise
*/
func T_StorageInterfaceHas(t *testing.T, impl StorageInterface) {
	pathFail := "Goodbye/World"
	pathSuccess := "Hello, world!"
	item := []byte("Hello, world!")

	exists, err := impl.Has(pathSuccess)
	if err != nil {
		t.Errorf(err.Error())
	}
	if exists {
		t.Errorf("The item should not exist yet.")
	}

	err = impl.Store(pathSuccess, item)
	if err != nil {
		t.Errorf(err.Error())
	}

	exists, err = impl.Has(pathSuccess)
	if err != nil {
		t.Errorf(err.Error())
	}
	if !exists {
		t.Errorf("The item should exist.")
	}

	exists, err = impl.Has(pathFail)
	if err != nil {
		t.Errorf(err.Error())
	}
	if exists {
		t.Errorf("The item should not exist.")
	}
}

/*
	Given that an item exists,
	When it is deleted,
	Then the impl should not return an error
	And an error if it does not exist
*/
func T_StorageInterfaceDelete(t *testing.T, impl StorageInterface) {
	path := "Hello, world!"
	item := []byte("Hello, world!")

	err := impl.Delete(path)
	if err == nil {
		t.Errorf("The item should not exist yet.")
	}

	err = impl.Store(path, item)
	if err != nil {
		t.Errorf(err.Error())
	}

	err = impl.Delete(path)
	if err != nil {
		t.Errorf(err.Error())
	}

	err = impl.Delete(path)
	if err == nil {
		t.Errorf("This item should already be deleted and should not exist.")
	}
}

/**
 * Any storage implementation should return an error if the item has a content
 * length of zero.
 */
func T_StorageInterfaceZeroLengthContent(t *testing.T, impl StorageInterface) {
	path := "Hello, world!"
	item := []byte("")

	err := impl.Store(path, item)
	if err == nil {
		t.Errorf("Expected interface to return an error if the item length is zero.")
	}
}

func _isByteArrEqual(a, b []byte) bool {
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
