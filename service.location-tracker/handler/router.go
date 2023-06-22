package handler

import (
	locationtrackerproto "github.com/sashajdn/sasha/service.location-tracker/proto"
)

type LocationTrackerService struct {
	locationtrackerproto.UnsafeLocationtrackerServer
}
