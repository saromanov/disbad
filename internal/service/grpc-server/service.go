package grpcserver

import (
	"context"
	"net"

	log "github.com/sirupsen/logrus"

	"github.com/saromanov/disbad/internal/disbad"
	"github.com/saromanov/disbad/internal/proto/master"
	"github.com/saromanov/disbad/internal/service"
)

type grpcServer struct {
	cfg    Config
	server *server
}

// New provides initialization of the grpc-server
func New(cfg Config, dis *disbad.Disbad) service.Service {
	return &grpcServer{
		cfg: cfg,
		server: &server{
			cfg: cfg,
			dis: dis,
		},
	}
}

// Run provides starting of the grpc server
func (s *grpcServer) Run(ctx context.Context, ready func()) error {
	logger := log.WithContext(ctx)
	s.server.Init(ctx, nil)
	listener, err := net.Listen("tcp", s.cfg.Address)
	if err != nil {
		logger.WithError(err).WithField("address", s.cfg.Address).Error("unable to listen tcp address")
		return err
	}

	logger.Info("starting of grpc server...")
	s.server.Init(ctx, nil)
	master.RegisterMasterServer(s.server.server, s.server)
	if err := s.server.server.Serve(listener); err != nil {
		return err
	}
	return nil
}

func (s *grpcServer) Shutdown(ctx context.Context) error {
	return nil
}
