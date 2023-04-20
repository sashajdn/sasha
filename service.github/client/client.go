package client

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

var (
	cli  GithubClient
	once sync.Once
)

var (
	ErrClientNotInitialized     = errors.New("github client not initialized")
	ErrClientAlreadyInitialized = errors.New("github client already initialized")
)

func Init() error {
	var err error
	once.Do(func() {
		cli = newDefaultGithubClient()

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err = cli.Ping(ctx); err != nil {
			return
		}
	})

	if err != nil {
		return fmt.Errorf("init github client: %w", err)
	}

	return nil
}

func WithMock() error {
	var err error
	once.Do(func() {
		cli = newMockClient()
	})

	if err != nil {
		return fmt.Errorf("init github client: %w", err)
	}

	return nil
}

type GithubClient interface {
	Ping(ctx context.Context) error
}

func Ping(ctx context.Context) error {
	if cli == nil {
		return ErrClientNotInitialized
	}

	return cli.Ping(ctx)
}
