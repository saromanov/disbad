package disbad

import (
	"sync"
)

type Node struct {
}
type leaderInfo struct {
	node         *Node
	replicaCount uint64
}

type Disbad struct {
	mutex   sync.RWMutex
	leaders map[string]leaderInfo
}

// New provides initialization of the grpc-server
func New() *Disbad {
	return &Disbad{
		leaders: map[string]leaderInfo{},
	}
}
