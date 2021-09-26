package grpcserver

import (
	"context"
	"fmt"

	"github.com/saromanov/disbad/internal/disbad"
	"github.com/saromanov/disbad/internal/proto/master"
)

type leaderInfo struct {
	node         *master.Node
	replicaCount uint64
}
type server struct {
	cfg Config
	dis *disbad.Disbad
}

// Init provides starting of the grpc server
func (s *server) Init(ctx context.Context) (*master.Response, error) {
	s.dis = disbad.New()
	return nil, nil
}

func (s *server) JoinMaster(ctx context.Context, node *master.Node) (*master.Cluster, error) {
	clusterID, err := s.dis.JoinMaster(ctx, node.GrpcAddress, node.RaftAddress)
	if err != nil {
		return nil, fmt.Errorf("unable to join master: %v", err)
	}
	return &master.Cluster{
		Id:                clusterID,
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
