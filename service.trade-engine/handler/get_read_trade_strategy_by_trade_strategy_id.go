package handler

import (
	"context"

	"github.com/sashajdn/sasha/libraries/gerrors"
	"github.com/sashajdn/sasha/service.trade-engine/dao"
	"github.com/sashajdn/sasha/service.trade-engine/marshaling"
	tradeengineproto "github.com/sashajdn/sasha/service.trade-engine/proto"
)

// ReadTradeStrategyByTradeStrategyID ...
func (s *TradeEngineService) ReadTradeStrategyByTradeStrategyID(
	ctx context.Context, in *tradeengineproto.ReadTradeStrategyByTradeStrategyIDRequest,
) (*tradeengineproto.ReadTradeStrategyByTradeStrategyIDResponse, error) {
	switch {
	case in.TradeStrategyId == "":
		return nil, gerrors.BadParam("missing_param.trade_strategy_id", nil)
	}

	errParams := map[string]string{
		"trade_strategy_id": in.TradeStrategyId,
	}

	trade, err := dao.ReadTradeStrategyByTradeStrategyID(ctx, in.TradeStrategyId)
	if err != nil {
		return nil, gerrors.Augment(err, "failed_to_read_trade_strategy_by_trade_strategy_id", errParams)
	}

	protoTrade := marshaling.TradeStrategyDomainToProto(trade)

	return &tradeengineproto.ReadTradeStrategyByTradeStrategyIDResponse{
		TradeStrategy: protoTrade,
	}, nil
}
