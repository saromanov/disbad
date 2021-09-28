package grpcserver

import (
	"context"
	"fmt"
	"sync"

	"github.com/google/uuid"
	"github.com/saromanov/disbad/internal/disbad"
	"github.com/saromanov/disbad/internal/proto/master"
	"google.golang.org/grpc"
)

type leaderInfo struct {
	node         *master.Node
	replicaCount uint64
}
type server struct {
	mutex   sync.RWMutex
	cfg     Config
	dis     *disbad.Disbad
	server  *grpc.Server
	leaders map[string]leaderInfo
}

// Init provides starting of the grpc server
func (s *server) Init(ctx context.Context, c *master.Cluster) (*master.Response, error) {
	s.dis = disbad.New()
	s.server = grpc.NewServer()
	s.leaders = map[string]leaderInfo{}
	return nil, nil
}

func (s *server) JoinMaster(ctx context.Context, node *master.Node) (*master.Cluster, error) {
	newClusterID := uuid.New().String()

	s.mutex.Lock()
	s.leaders[newClusterID] = leaderInfo{&master.Node{
		ClusterId:   newClusterID,
		GrpcAddress: node.GrpcAddress,
		RaftAddress: node.RaftAddress,
	}, 1}
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
	s.mutex.Lock()
	_, ok := s.leaders[node.ClusterId]
	if !ok {
		return nil, fmt.Errorf("unable to bind cluster id")
	}
	delete(s.leaders, node.ClusterId)
	s.mutex.Unlock()
	return nil, nil
}

func (s *server) JoinExistingCluster(ctx context.Context, node *master.Node) (*master.Cluster, error) {
	return nil, nil
}

// GetMaster returns leader of the cluster
func (s *server) GetMaster(ctx context.Context, cluster *master.Cluster) (*master.Node, error) {
	node, err := s.dis.GetMaster(ctx, cluster.Id)
	if err != nil {
		return nil, fmt.Errorf("unable to get node: %v", err)
	}
	return &master.Node{
		RaftAddress: node.RaftAddress,
		GrpcAddress: node.GrpcAddress,
		ClusterId:   node.ClusterID,
		Id:          node.ID,
	}, nil
}
