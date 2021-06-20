package grpc-server

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