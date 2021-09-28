package rest

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/saromanov/disbad/internal/service"
	"github.com/saromanov/disbad/internal/service/rest/models"
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
	r.GET("/api/v1/keys", func(c *gin.Context) {
		
	})
	r.POST("/api/v1/keys", func(c *gin.Context) {
		var req models.KeysRequest
		if err := c.Bind(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
	})
	r.Run(s.cfg.Address)
	return nil
}

func (s *rest) Shutdown(ctx context.Context) error {
	return nil
}
