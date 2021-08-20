package rest

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/saromanov/disbad/internal/service"
)

type rest struct {
	cfg Config
}

func New(cfg Config) service.Service {
	return &rest{
		cfg: cfg,
	}
}

// Run provides starting of the grpc server
func (s *rest) Run(ctx context.Context, ready func()) error {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		
	})
	r.Run(s.cfg.Address)
	return nil
}

func (s *rest) Shutdown(ctx context.Context) error {
	return nil
}
