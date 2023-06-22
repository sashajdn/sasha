package dao

import (
	"context"

	"github.com/sashajdn/sasha/service.trade-engine/domain"
)

//  AddParticpantToTradeStrategy ...
func AddParticpantToTradeStrategy(ctx context.Context, tradeParticipant *domain.TradeStrategyParticipant) error {
	return nil
}

// ReadTradeStrategyParticipantByTradeStrategyID ...
func ReadTradeStrategyParticipantByTradeStrategyID(ctx context.Context, tradeStrategyID, userID string) (*domain.TradeStrategyParticipant, error) {
    return nil, nil
}
