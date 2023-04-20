package client

import "context"

type GithubClient interface {
	Ping(ctx context.Context) error
}
