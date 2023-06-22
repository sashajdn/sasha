package main

import (
	"context"

	"github.com/sashajdn/sasha/libraries/mariana"
	"github.com/sashajdn/sasha/service.trade-engine/dao"
	"github.com/sashajdn/sasha/service.trade-engine/handler"
	tradeengineproto "github.com/sashajdn/sasha/service.trade-engine/proto"
)

const (
	svcName = "s.trade-engine"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Init Dao
	if err := dao.Init(ctx, svcName); err != nil {
		panic(err)
	}

	// Init Mariana Server
    // TODO: pass logger.
	srv := mariana.Init(svcName, nil)
	tradeengineproto.RegisterTradeengineServer(srv.Grpc(), &handler.TradeEngineService{})
	srv.Run(ctx)
}
