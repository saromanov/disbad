package grpcserver

import (
	"context"
	"fmt"
	"sync"

	uuid "github.com/google/uuid"

	"github.com/saromanov/disbad/internal/proto/master"
)

type leaderInfo struct {
	node         *master.Node
	replicaCount uint64
}
type server struct {
	cfg     Config
	mutex   sync.RWMutex
	leaders map[string]leaderInfo
}

// Inuit provides starting of the grpc server
func (s *server) Init(ctx context.Context, c *master.Cluster) (*master.Response, error) {

	return nil, nil
}

func (s *server) JoinMaster(ctx context.Context, node *master.Node) (*master.Cluster, error) {
	newClusterID := uuid.New().String()

	s.mutex.Lock()
	s.leaders[newClusterID] = leaderInfo{node, 1}
	s.mutex.Unlock()
	return &master.Cluster{
		Id:                newClusterID,
		MasterGrpcAddress: node.GrpcAddress,
		MasterRaftAddress: node.RaftAddress,
	}, nil

}

func (s *server) UpdateMaster(ctx context.Context, node *master.Node) (*master.Response, error) {
	return nil, nil
}

func (s *server) LeaveCluster(ctx context.Context, node *master.Node) (*master.Response, error) {
	return nil, nil
}

func (s *server) JoinExistingCluster(ctx context.Context, node *master.Node) (*master.Cluster, error) {
	return nil, nil
}

// GetMaster returns leader of the cluster
func (s *server) GetMaster(ctx context.Context, cluster *master.Cluster) (*master.Node, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	leader, ok := s.leaders[cluster.Id]
	if !ok {
		return nil, fmt.Errorf("unable to get cluster leader")
	}
	if leader.node == nil {
		return nil, fmt.Errorf("unknown error. unable to get node")
	}
	return leader.node, nil
}
