package badger

import (
	"github.com/dgraph-io/badger/v3"
)
type Badger struct {
	db *badger.DB
}