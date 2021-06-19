package grpc-server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

type server struct {
	cfg Config
}

// New provides initialization of the grpc-server
func New(cfg Config) service.Service {
	return &rest{
		cfg: cfg,
	}
}

func (s *rest) Run(ctx context.Context, ready func()) error {
	logger := log.WithContext(ctx)
	listener, err := net.Listen("tcp", s.cfg.Address)
	if err != nil {
		logger.WithError(err).Error("unable to listen tcp address")
		return err
	}

	clusterAdminGRPCServer = NewClusterAdminGRPCServer()
	RavelClusterAdminPB.RegisterRavelClusterAdminServer(clusterAdminGRPCServer.Server, clusterAdminGRPCServer)
	if err := clusterAdminGRPCServer.Server.Serve(listener); err != nil {
		return err
	}
	return nil
}

func (s *rest) Shutdown(ctx context.Context) error {
	return nil
}
