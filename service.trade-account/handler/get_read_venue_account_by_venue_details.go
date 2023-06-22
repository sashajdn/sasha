package handler

import (
	"context"

	"github.com/sashajdn/sasha/libraries/gerrors"
	"github.com/sashajdn/sasha/service.trade-account/dao"
	"github.com/sashajdn/sasha/service.trade-account/marshaling"
	tradeaccountproto "github.com/sashajdn/sasha/service.trade-account/proto"
	tradeengineproto "github.com/sashajdn/sasha/service.trade-engine/proto"
)

// ReadVenueAccountByVenueAccountDetails ...
func (s *TradeAccountService) ReadVenueAccountByVenueAccountDetails(
	ctx context.Context, in *tradeaccountproto.ReadVenueAccountByVenueAccountDetailsRequest,
) (*tradeaccountproto.ReadVenueAccountByVenueAccountDetailsResponse, error) {
	switch {
	case in.UserId == "":
		return nil, gerrors.BadParam("missing_param.user_id", nil)
	case in.ActorId == "":
		return nil, gerrors.BadParam("missing_param.actor_id", nil)
	case in.Venue == tradeengineproto.VENUE_UNREQUIRED:
		return nil, gerrors.BadParam("missing_param.venue", nil)
	case in.RequestContext == "":
		return nil, gerrors.BadParam("missing_param.request_context", nil)
	}

	errParams := map[string]string{
		"user_id":    in.UserId,
		"actor_id":   in.ActorId,
		"venue":      in.Venue.String(),
		"subaccount": in.Subaccount,
	}

	// Validate the request venue credentials.
	if err := validateVenueAccountDetailsRequest(in); err != nil {
		return nil, gerrors.Augment(err, "failed_to_read_venue_account_by_venue_details", errParams)
	}

	// Read.
	venueAccount, err := dao.ReadVenueAccountByVenueAccountDetails(ctx, in.Venue.String(), in.UserId)
	if err != nil {
		return nil, gerrors.Augment(err, "failed_to_read_venue_account_by_venue_details.dao", errParams)
	}

	// Marshal.
	proto, err := marshaling.VenueAccountDomainToProtoUnmasked(venueAccount)
	if err != nil {
		return nil, gerrors.Augment(err, "failed_to_read_venue_account_by_venue_details.marshal_to_proto", errParams)
	}

	return &tradeaccountproto.ReadVenueAccountByVenueAccountDetailsResponse{
		VenueAccount: proto,
	}, nil
}

func validateVenueAccountDetailsRequest(in *tradeaccountproto.ReadVenueAccountByVenueAccountDetailsRequest) error {
	switch in.RequestContext {
	case tradeaccountproto.RequestContextOrderRequest, tradeaccountproto.RequestContextUserRequest:
	default:
		return gerrors.FailedPrecondition("invalid_request_context", nil)
	}

	switch in.Venue {
	case  tradeengineproto.VENUE_BITFINEX, tradeengineproto.VENUE_DERIBIT:
	default:
		return gerrors.Unimplemented("venue.unimplemented", nil)
	}

	switch in.ActorId {
	case in.UserId:
	case tradeaccountproto.ActorSystemTradeEngine, tradeaccountproto.ActorSystemPayments:
	default:
		return gerrors.FailedPrecondition("bad_actor", nil)
	}

	return nil
}
