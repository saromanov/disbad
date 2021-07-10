package grpcserver

import (
	"context"
	"net"

	log "github.com/sirupsen/logrus"

	"github.com/saromanov/disbad/internal/proto/master"
	"github.com/saromanov/disbad/internal/service"
)

type server struct {
	cfg Config
}

// New provides initialization of the grpc-server
func New(cfg Config) service.Service {
	return &server{
		cfg: cfg,
	}
}

// Run provides starting of the grpc server
func (s *server) Run(ctx context.Context, ready func()) error {
	logger := log.WithContext(ctx)
	listener, err := net.Listen("tcp", s.cfg.Address)
	if err != nil {
		logger.WithError(err).Error("unable to listen tcp address")
		return err
	}

	clusterAdminGRPCServer = New()
	master.RegisterMasterServer(clusterAdminGRPCServer.Server, clusterAdminGRPCServer)
	if err := clusterAdminGRPCServer.Server.Serve(listener); err != nil {
		return err
	}
	return nil
}

func (s *server) Shutdown(ctx context.Context) error {
	return nil
}
