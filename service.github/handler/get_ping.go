package handler

import (
	githubproto "github.com/sashajdn/sasha/service.github/proto"
	"golang.org/x/net/context"
)

func (g *GithubService) Ping(
	ctx context.Context, req *githubproto.PingRequest,
) (*githubproto.PingResponse, error) {
	return nil, nil
}
