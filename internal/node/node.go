package node

import (
	"encoding/json"
	"time"

	"github.com/hashicorp/raft"
)

type Node struct {
	Raft *raft.Raft
}

func (n *Node) Get(key []byte) ([]byte, error) {
	if n.Raft.State() != raft.Leader {
		return []byte{}, raft.ErrNotLeader
	}
	return n.Fsm.Get(key)
}

func (n *Node) Set(key []byte, value []byte) error {
	if n.Raft.State() != raft.Leader {
		return raft.ErrNotLeader
	}

	var data fsm.LogData

	data.Operation = "set"
	data.Key = key
	data.Value = value

	dataBuffer, err := json.Marshal(data)
	if err != nil {
		return err
	}

	f := n.Raft.Apply(dataBuffer, 5*time.Second)

	return f.Error()
}

func (n *Node) Delete(key []byte) error {
	if n.Raft.State() != raft.Leader {
		return raft.ErrNotLeader
	}

	var data fsm.LogData

	data.Operation = "delete"
	data.Key = key
	data.Value = []byte{}

	dataBuffer, err := json.Marshal(data)
	if err != nil {
		log.Fatal("RavelNode: Unable to marhsal key value")
		return err
	}

	f := n.Raft.Apply(dataBuffer, 3*time.Second)

	return f.Error()
}