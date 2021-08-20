package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/joeshaw/envdecode"
	_ "github.com/joho/godotenv/autoload"
	"github.com/oklog/run"
	"github.com/sirupsen/logrus"

	"github.com/saromanov/disbad/internal/service"
	grpc "github.com/saromanov/disbad/internal/service/grpc-server"
	"github.com/saromanov/disbad/internal/service/rest"
)

type config struct {
	GRPC grpc.Config
	Rest rest.Config
}

func main() {
	var cfg config
	if err := envdecode.StrictDecode(&cfg); err != nil {
		logrus.WithError(err).Fatal("Cannot decode config envs")
	}
	logger := logrus.New()
	logger.Level = logrus.DebugLevel
	logger.Formatter = &logrus.JSONFormatter{}
	ctx, cancel := context.WithCancel(context.Background())
	g := &run.Group{}
	{
		stop := make(chan os.Signal)
		signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
		g.Add(func() error {
			<-stop
			return nil
		}, func(error) {
			signal.Stop(stop)
			cancel()
			close(stop)
		})
	}

	st := grpc.New(cfg.GRPC)
	r := rest.New(cfg.Rest)

	s := service.Runner{}
	if err := s.SetupService(ctx, st, "grpc-server", g); err != nil {
		logger.WithError(err).Fatal("unable to setup service ")
	}
	if err := s.SetupService(ctx, r, "rest-server", g); err != nil {
		logger.WithError(err).Fatal("unable to setup service ")
	}
	logger.Info("Running of the service...")
	if err := g.Run(); err != nil {
		logger.WithError(err).Fatal("The service has been stopped with error")
	}
	logger.Info("Service is stopped")
}
