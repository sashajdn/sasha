package handler

import (
	tradeaccountproto "github.com/sashajdn/sasha/service.trade-account/proto"
)

// TradeAccountService ...
type TradeAccountService struct {
	*tradeaccountproto.UnimplementedTradeaccountServer
}
