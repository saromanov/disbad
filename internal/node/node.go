package node

import (
	"encoding/json"
	"time"

	"github.com/saromanov/disbad/internal/fsm"

	"github.com/hashicorp/raft"
)

type Node struct {
	raft *raft.Raft
	fsm fsm.FSM
}

// Get provides getting of the node
func (n *Node) Get(key []byte) ([]byte, error) {
	if n.raft.State() != raft.Leader {
		return []byte{}, raft.ErrNotLeader
	}
	return n.fsm.Get(key)
}

func (n *Node) Set(key []byte, value []byte) error {
	return n.handle("set", key, value)
}

func (n *Node) Delete(key []byte) error {
	return n.handle("delete", key, []byte{})
}

func (n *Node) handle(operation string, key, value []byte) error {
	if n.raft.State() != raft.Leader {
		return raft.ErrNotLeader
	}

	var data fsm.LogData

	data.Operation = operation
	data.Key = key
	data.Value = value

	dataBuffer, err := json.Marshal(data)
	if err != nil {
		return err
	}

	f := n.raft.Apply(dataBuffer, 3*time.Second)

	return f.Error()
}