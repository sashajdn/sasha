package handler

import (
	githubproto "github.com/sashajdn/sasha/service.github/proto"
)

type GithubService struct {
	*githubproto.UnimplementedGithubServer
}
