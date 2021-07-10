package storage

// Storage defines interface for storage implementation
type Storage interface {
	Get([]byte)([]byte, error)
	Set([]byte, []byte) error
	Delete([]byte) error
}