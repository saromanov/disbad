package disbad

import (
	"context"
	"fmt"
	"sync"

	uuid "github.com/google/uuid"
	"github.com/saromanov/disbad/internal/models"
)

type leaderInfo struct {
	node         *models.Node
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

func (d *Disbad) JoinMaster(ctx context.Context, grpcAddr, raftAddr string) (string, error) {
	newClusterID := uuid.New().String()

	d.mutex.Lock()
	d.leaders[newClusterID] = leaderInfo{&models.Node{
		ClusterID:   newClusterID,
		GrpcAddress: grpcAddr,
		RaftAddress: raftAddr,
	}, 1}
	d.mutex.Unlock()
	return newClusterID, nil
}

func (d *Disbad) GetMaster(ctx context.Context, clusterID string) (*models.Node, error) {
	d.mutex.RLock()
	defer d.mutex.RUnlock()
	leader, ok := d.leaders[clusterID]
	if !ok {
		return nil, fmt.Errorf("unable to get cluster leader")
	}
	if leader.node == nil {
		return nil, fmt.Errorf("unknown error. unable to get node")
	}
	return leader.node, nil
}

// ReadKey provides reading of the key
func (d *Disbad) ReadKey(ctx context.Context, clusterID, key string) (string, error) {
	if clusterID == "" {
		return "", fmt.Errorf("cluster id is not defined")
	}
	if key == "" {
		return "", fmt.Errorf("key is not defined")
	}
	return "", nil
}
