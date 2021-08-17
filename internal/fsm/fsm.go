package fsm

import (
	"github.com/saromanov/disbad/internal/storage"
)

type FSM struct {
	db storage.Storage
}

type Data struct {
	Operation string `json:"Operation"`
	Key       []byte `json:"Key"`
	Value     []byte `json:"Value"`
}

// New creates an instance of FSM
func New(path string) (*FSM, error) {
	var r storage.Storage
	err := r.Init(path)
	if err != nil {
		return nil, err
	}

	return &FSM{
		Db: &r,
	}, nil
}
