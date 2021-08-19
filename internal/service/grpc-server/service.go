package grpcserver

import (
	"context"
	"net"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/saromanov/disbad/internal/proto/master"
	"github.com/saromanov/disbad/internal/service"
)

type grpcServer struct {
	cfg    Config
	server *server
}

// New provides initialization of the grpc-server
func New(cfg Config) service.Service {
	return &grpcServer{
		cfg: cfg,
		server: &server{
			leaders: map[string]leaderInfo{},
			srv:     grpc.NewServer(),
		},
	}
}

// Run provides starting of the grpc server
func (s *grpcServer) Run(ctx context.Context, ready func()) error {
	logger := log.WithContext(ctx)
	listener, err := net.Listen("tcp", s.cfg.Address)
	if err != nil {
		logger.WithError(err).WithField("address", s.cfg.Address).Error("unable to listen tcp address")
		return err
	}

	master.RegisterMasterServer(s.server.srv, s.server)
	if err := s.server.srv.Serve(listener); err != nil {
		return err
	}
	return nil
}

func (s *grpcServer) Shutdown(ctx context.Context) error {
	return nil
}
