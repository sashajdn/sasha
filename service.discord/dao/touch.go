package dao

import (
	"context"

	"github.com/sashajdn/sasha/service.discord/domain"
)

// Exists checkcks
func Exists(ctx context.Context, idempotencyKey string) (*domain.Touch, bool, error) {
	// var sql     = `SELECT * FROM s_discord_touches WHERE idempotency_key=$1`

    return nil, false, nil
}

// Update updates existing touch via merging with mutation & persisting.
func Update(ctx context.Context, mutation *domain.Touch) (*domain.Touch, error) {
    // var sql = ` UPDATE s_discord_touches SET idempotency_key=$1, updated=$2, sender_id=$3`

    // Check key.
    // Merge updates.
    // return merged update.

    return nil, nil
}

func Create(ctx context.Context, touch *domain.Touch) (*domain.Touch, error) {
	// var sql = ` INSERT INTO s_discord_touches (idempotency_key, updated, sender_id) VALUES ($1, $2, $3)`

	return touch , nil
}
