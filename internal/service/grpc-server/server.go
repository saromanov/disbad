package grpcserver

import (
	"sync"
	"google.golang.org/grpc"
)

type Server struct {
	mutex            sync.Mutex
	srv           *grpc.Server
}

// New provides initialization of grps server
func New() *Server{
	return &Server {
		srv: grpc.NewServer(),
	}
}

// Join provides joining of the node
func (s *Server) Join() error {
	return nil
}

func (s *Server) Leave() error {
	return nil
}