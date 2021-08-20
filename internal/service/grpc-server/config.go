package grpcserver

// Config defines configuration for grpc
type Config struct {
	Address string `env:"GRPC_ADDRESS,required"`
}