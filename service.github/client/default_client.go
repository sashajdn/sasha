package client

import (
	"context"
	"fmt"
)

func newDefaultGithubClient() *defaultGithubClient {
	return &defaultGithubClient{}
}

var _ GithubClient = &defaultGithubClient{}

type defaultGithubClient struct{}

func (d *defaultGithubClient) Ping(ctx context.Context) error {
	return fmt.Errorf("ping unimplemented")
}
