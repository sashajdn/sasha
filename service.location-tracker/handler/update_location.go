package handler

import (
	"context"
	"fmt"

	locationtrackerproto "github.com/sashajdn/sasha/service.location-tracker/proto"
)

// UpdateLocation updates the location given the `Location` input parameter.
func (l *LocationTrackerService) UpdateLocation(
	ctx context.Context, req *locationtrackerproto.UpdateLocationRequest,
) (*locationtrackerproto.UpdateLocationResponse, error) {
	if err := validateLocation(req.Location); err != nil {
		return nil, fmt.Errorf("invalid location")
	}

	return nil, fmt.Errorf("unimplemented update location method")
}

func validateLocation(l *locationtrackerproto.Location) error {
	return nil
}
