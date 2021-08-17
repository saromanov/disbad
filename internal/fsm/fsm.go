package fsm

import (
	"encoding/json"

	"github.com/saromanov/disbad/internal/storage"
	"github.com/saromanov/disbad/internal/models"

	"github.com/hashicorp/raft"
)

type FSM struct {
	db storage.Storage
}

// New creates an instance of FSM
func New(path string) (*FSM, error) {
	var r storage.Storage
	err := r.Init(path)
	if err != nil {
		return nil, err
	}

	return &FSM{
		db: r,
	}, nil
}

// Get provides getting of the state from the storage
func (f *FSM) Get(key []byte) ([]byte, error) {
	data, err := f.db.Get(key)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Write provides writing state to the storage
func (f *FSM) Write(r *raft.Log) error {
	var l models.Log
	if err := json.Unmarshal(r.Data, &l); err != nil {
		return err
	}
	if err:= f.db.Set(l.Key, l.Value); err != nil {
		return err
	}
	return nil
}
