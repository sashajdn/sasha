package main

import (
	"context"
	"log"
	"os/signal"

	"go.uber.org/zap"

	"github.com/sashajdn/sasha/libraries/environment"
	"github.com/sashajdn/sasha/libraries/mariana"

	"github.com/sashajdn/sasha/service.openai/dao"
	"github.com/sashajdn/sasha/service.openai/handler"
	openaiproto "github.com/sashajdn/sasha/service.openai/proto"
)

const serviceName = "service.openai"

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

	if err := dao.Init(serviceName, cfg.Cassandra, slogger); err != nil {
		logger.With(zap.Error(err)).Fatal("Failed to init dao")
	}

	// Init Mariana Server.
	srv := mariana.InitWithConfig(serviceName, cfg, slogger)
	openaiproto.RegisterOpenaiServer(srv.Grpc(), &handler.OpenAIService{})

	srv.Run(ctx)
}
