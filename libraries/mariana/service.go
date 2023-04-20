package mariana

import (
	"context"
	"fmt"
	"net"
	"os"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/sashajdn/sasha/libraries/environment"
)

const (
	network            = "tcp"
	defaultServicePort = "8000"
)

// Server defines the interface for our base server setup.
type Service interface {
	Run(ctx context.Context)
	Grpc() *grpc.Server
	Name() string
}

// Init inits our base server.
// TODO: deprecate.
func Init(serviceName string, logger *zap.SugaredLogger) Service {
	logger = logger.With(zap.String("service_name", serviceName))

	return initService(serviceName, nil, logger)
}

// InitWithConfig ...
func InitWithConfig(serviceName string, cfg *environment.Environment, logger *zap.SugaredLogger) Service {
	logger = logger.With(zap.String("service_name", serviceName))

	return initService(serviceName, cfg, logger)
}

func initService(serviceName string, cfg *environment.Environment, logger *zap.SugaredLogger) *BaseService {
	grpcs := grpc.NewServer()

	reflection.Register(grpcs)
	s := &BaseService{
		s:      grpcs,
		name:   serviceName,
		logger: logger,
	}

	if cfg != nil {
		s.Config = cfg
	}

	return s
}

type BaseService struct {
	// Service Name
	Config *environment.Environment

	// Meta.
	name string

	// GRPC server.
	s      *grpc.Server
	logger *zap.SugaredLogger
}

// Run runs our base server.
func (s *BaseService) Run(ctx context.Context) {
	hostname, err := os.Hostname()
	if err != nil {
		panic(fmt.Sprintf("Failed to establish hostname: %v", err))
	}

	addr := formatAddr(hostname, defaultServicePort)

	listener, err := net.Listen(network, addr)
	if err != nil {
		panic(fmt.Sprintf("%s failed to listen on %s:%s: %v", s.name, network, addr, err))
	}

	if err := s.s.Serve(listener); err != nil {
		panic(err)
	}
}

// Grpc returns the underlying gRPC server.
func (s *BaseService) Grpc() *grpc.Server {
	return s.s
}

func (s *BaseService) Name() string { return s.name }

func formatAddr(serviceName, port string) string {
	return fmt.Sprintf("%s:%s", serviceName, port)
}
