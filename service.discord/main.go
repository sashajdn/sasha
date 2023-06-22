package main

import (
	"context"

	"github.com/sashajdn/sasha/libraries/mariana"
	"github.com/sashajdn/sasha/service.discord/client"
	"github.com/sashajdn/sasha/service.discord/dao"
	"github.com/sashajdn/sasha/service.discord/handler"
	discordproto "github.com/sashajdn/sasha/service.discord/proto"
)

const (
	svcName = "s.discord"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Init dao
	if err := dao.Init(ctx, svcName); err != nil {
		panic(err)
	}

	// Init Client
	if err := client.Init(ctx); err != nil {
		panic(err)
	}

	// Init gRPC server
    // TODO: add logger
	s := mariana.Init(svcName, nil)
	discordproto.RegisterDiscordServer(s.Grpc(), &handler.DiscordService{})
	s.Run(ctx)
}
