package storage

type StorageInterface interface {
	Has(path string) (bool, error)
	Store(path string, data []byte) error
	Retrieve(path string) ([]byte, error)
	Delete(path string) error
}
