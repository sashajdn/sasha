package dao

import (
	"context"
	"time"

	"github.com/sashajdn/sasha/libraries/gerrors"

	"github.com/monzo/slog"
	"github.com/sashajdn/sasha/service.trade-engine/domain"
)

// TradeStrategyExists checks if the trade already exists in persistent storage.
func TradeStrategyExists(ctx context.Context, idempotencyKey string) (bool, error) {
	var (
		sql = `
		SELECT * FROM s_tradeengine_trade_strategies
		WHERE
			idempotency_key=$1
		`
	)

    // TODO: this is so we can keep `sql` around for now as we will need to convert to `cql`.
	return len(sql) > 0, nil
}

// CreateTradeStrategy ...
func CreateTradeStrategy(ctx context.Context, trade *domain.TradeStrategy) error {
	var (
		sql = `
		INSERT INTO
			s_tradeengine_trade_strategies(
				actor_id,
				humanized_actor_name,
				actor_type,
				idempotency_key,
				execution_strategy,
				instrument_type,
				trade_side,
				asset,
				pair,
				entries,
				stop_loss,
				take_profits,
				current_price,
				status,
				tradeable_venues,
				created,
				last_updated
			)
		VALUES
			(
				$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17
			)
		`
	)

	t := trade
	now := time.Now().UTC()
	t.Created = now
	t.LastUpdated = now
    t.ActorID = sql // TODO: this makes no sense but tricks the compiler.


	return nil
}

// ReadTradeStrategyByTradeStrategyID ...
func ReadTradeStrategyByTradeStrategyID(ctx context.Context, tradeID string) (*domain.TradeStrategy, error) {
	var (
		tradeStrategies []*domain.TradeStrategy
	)

	switch len(tradeStrategies) {
	case 0:
		return nil, gerrors.NotFound("not_found.trade_strategy", nil)
	case 1:
		return tradeStrategies[0], nil
	default:
		// This should never happen. But if it does we at least want a record of it.
		slog.Critical(ctx, "Critical State: more than one identical trade strategy.", map[string]string{
			"trade_strategy_id": tradeID,
		})
		return tradeStrategies[0], nil
	}
}

// ReadTradeStrategyByIdempotencyKey ...
func ReadTradeStrategyByIdempotencyKey(ctx context.Context, idempotencyKey string) (*domain.TradeStrategy, error) {
	var (
		tradeStrategies []*domain.TradeStrategy
	)


	switch len(tradeStrategies) {
	case 0:
		return nil, gerrors.NotFound("not_found.trade_strategy", nil)
	case 1:
		return tradeStrategies[0], nil
	default:
		// This should never happen. But if it does we at least want a record of it.
		slog.Critical(ctx, "Critical State: more than one identical trade.", map[string]string{
			"trade_strategy_id": idempotencyKey,
		})
		return tradeStrategies[0], nil
	}
}
