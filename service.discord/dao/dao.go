package dao

import (
	"context"
	"sync"
)

var (
	mu sync.Mutex
)

// Init creates the database connection.
func Init(ctx context.Context, serviceName string) error {
	return nil
}

// WithMock uses a mock db.
func WithMock() {
}
