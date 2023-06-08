package handler

import (
	"context"

	locationtrackerproto "github.com/sashajdn/sasha/service.location-tracker/proto"
)

// Ping pings the location tracker service.
func (l *LocationTrackerService) Ping(
	ctx context.Context, req *locationtrackerproto.PingRequest,
) (*locationtrackerproto.PingResponse, error) {
	return &locationtrackerproto.PingResponse{}, nil
}
