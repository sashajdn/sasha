package orderrouter

import (
	"context"
	"strings"

	"github.com/monzo/slog"

	"github.com/sashajdn/sasha/libraries/gerrors"
	tradeengineproto "github.com/sashajdn/sasha/service.trade-engine/proto"
)

// RouteAndExecuteNewOrder routes order to the correct exchange & executes.
func RouteAndExecuteNewOrder(
	ctx context.Context,
	_ *tradeengineproto.Order,
	venue tradeengineproto.VENUE,
	instrumentType tradeengineproto.INSTRUMENT_TYPE,
	_ *tradeengineproto.VenueCredentials,
) (*tradeengineproto.Order, error) {

	errParams := map[string]string{
		"venue_id":        strings.ToLower(venue.String()),
		"instrument_type": strings.ToLower(instrumentType.String()),
	}

	switch venue {
	default:
		slog.Error(ctx, "Failed to route order: venue, instrument pair not implemented: %+v", errParams)
		return nil, gerrors.Unimplemented("failed_to_route_and_execute_order.venue_unimplemented", errParams)
	}
}
