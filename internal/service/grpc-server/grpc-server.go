package grpcserver

import (
	"context"
	"fmt"
	"net"
	"sync"

	uuid "github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/saromanov/disbad/internal/proto/master"
	"github.com/saromanov/disbad/internal/service"
)

type leaderInfo struct {
	node         *master.Node
	replicaCount uint64
}
type server struct {
	cfg     Config
	mutex   sync.RWMutex
	leaders map[string]leaderInfo
	srv     *grpc.Server
}

// New provides initialization of the grpc-server
func New(cfg Config) service.Service {
	return &server{
		cfg:     cfg,
		leaders: map[string]leaderInfo{},
		srv: grpc.NewServer(),
	}
}

// Run provides starting of the grpc server
func (s *server) Run(ctx context.Context, ready func()) error {
	logger := log.WithContext(ctx)
	listener, err := net.Listen("tcp", s.cfg.Address)
	if err != nil {
		logger.WithError(err).WithField("address", s.cfg.Address).Error("unable to listen tcp address")
		return err
	}

	master.RegisterMasterServer(s.srv, s)
	if err := s.srv.Serve(listener); err != nil {
		return err
	}
	return nil
}

func (s *server) JoinAsLeader(node *master.Node) *master.Cluster {
	newClusterID := uuid.New().String()

	s.mutex.Lock()
	s.leaders[newClusterID] = leaderInfo{node, 1}
	s.mutex.Unlock()
	return &master.Cluster{
		Id:                newClusterID,
		MasterGrpcAddress: node.GrpcAddress,
		MasterRaftAddress: node.RaftAddress,
	}

}

// GetMaster returns leader of the cluster
func (s *server) GetMaster(cluster *master.Cluster) (*master.Node, error) {
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

func (s *server) Shutdown(ctx context.Context) error {
	return nil
}
