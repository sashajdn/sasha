package client

import "context"

func newMockClient() *mockClient {
	return &mockClient{}
}

var _ GithubClient = &mockClient{}

type mockClient struct{}

func (m *mockClient) Ping(ctx context.Context) error { return nil }
