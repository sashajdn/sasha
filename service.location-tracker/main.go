package main

import (
	"context"
	"log"
	"os/signal"

	"github.com/sashajdn/sasha/libraries/environment"
	"github.com/sashajdn/sasha/libraries/mariana"
	"go.uber.org/zap"
)

const serviceName = "serivce.location-tracker"

func main() {
	ctx, cancel := signal.NotifyContext(context.Background())
	defer cancel()

	cfg, err := environment.LoadEnvironment()
	if err != nil {
		log.Fatalf("Failed to load environment: %v", err)
	}

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Failed to create zap logger: %v", err)
	}
	slogger := logger.Sugar().With(
		zap.String("service_name", serviceName),
		zap.String("environment", cfg.Metadata.Environment),
		zap.String("namespace", cfg.Metadata.Namespace),
	)

	srv := mariana.InitWithConfig(serviceName, cfg, slogger)
	// TODO: register proto.

	srv.Run(ctx)
}
