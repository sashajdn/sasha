package main

import (
	"context"

	"github.com/sashajdn/sasha/libraries/mariana"
	"github.com/sashajdn/sasha/service.trade-account/dao"
	"github.com/sashajdn/sasha/service.trade-account/handler"
	tradeaccountproto "github.com/sashajdn/sasha/service.trade-account/proto"
)

const (
	svcName = "service.trade-account"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Init Dao.
	if err := dao.Init(ctx, svcName); err != nil {
		panic(err)
	}

	// Init Mariana Server.
    // TODO: logger.
	srv := mariana.Init(svcName, nil)
	tradeaccountproto.RegisterAccountServer(srv.Grpc(), &handler.AccountService{})
	srv.Run(ctx)
}
