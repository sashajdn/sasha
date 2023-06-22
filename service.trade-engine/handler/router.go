package handler

import (
	tradeengineproto "github.com/sashajdn/sasha/service.trade-engine/proto"
)

// TradeEngineService ...
type TradeEngineService struct {
	*tradeengineproto.UnimplementedTradeengineServer
}
