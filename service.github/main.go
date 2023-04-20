package main

import (
	"context"
	"log"
	"os/signal"

	"go.uber.org/zap"

	"github.com/sashajdn/sasha/libraries/environment"
	"github.com/sashajdn/sasha/libraries/mariana"
	"github.com/sashajdn/sasha/service.github/handler"
	githubproto "github.com/sashajdn/sasha/service.github/proto"
)

const serviceName = "service.github"

func main() {
	ctx, cancel := signal.NotifyContext(context.Background())
	defer cancel()

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Failed to create zap logger: %v", err)
	}
	slogger := logger.Sugar()

	cfg, err := environment.LoadEnvironment()
	if err != nil {
		log.Fatalf("Failed to load environment: %v", err)
	}

	// Init Mariana Server
	srv := mariana.InitWithConfig(serviceName, cfg, slogger)
	githubproto.RegisterGithubServer(srv.Grpc(), &handler.GithubService{})

	srv.Run(ctx)
}
