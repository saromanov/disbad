package badger

import (
	"fmt"

	"github.com/dgraph-io/badger/v3"
	"github.com/saromanov/disbad/internal/storage"
)

//Badger defines struct for db
type Badger struct {
	db *badger.DB
}

// New defines initialization of Badger
func New() (storage.Storage, error) {
	opts := badger.DefaultOptions("")
	db, err := badger.Open(opts)
	if err != nil {
		return nil, fmt.Errorf("unable to initialize badger: %v", err)
	}
	return &Badger{
		db: db,
	}, nil
}

// Set adds new key-value pair
func (b *Badger) Set(key, value []byte) error {
	err := b.db.Update(func(txn *badger.Txn) error {
		return txn.Set(key, value)
	})

	if err != nil {
		return fmt.Errorf("unable to set data: %v", err)
	}
	return nil
}

// Get getting value by the key
func (b *Badger) Get(key []byte) ([]byte, error) {
	var valCopy []byte
	err := b.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err != nil {
			return err
		}
		valCopy, err = item.ValueCopy(nil)
		return err
	})

	if err != nil {
		return nil, fmt.Errorf("unable to get data: %v", err)
	}
	return valCopy, nil
}

// Delete removes key-value pair
func (b *Badger) Delete(key []byte) error {
	err := b.db.Update(func(txn *badger.Txn) error {
		return txn.Delete(key)
	})
	if err != nil {
		return fmt.Errorf("unable to delete data: %v", err)
	}
	return nil
}
