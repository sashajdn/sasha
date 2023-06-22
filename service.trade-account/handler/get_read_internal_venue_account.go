package handler

import (
	"context"

	"github.com/sashajdn/sasha/libraries/gerrors"
	"github.com/sashajdn/sasha/service.trade-account/dao"
	"github.com/sashajdn/sasha/service.trade-account/marshaling"
	tradeaccountproto "github.com/sashajdn/sasha/service.trade-account/proto"
	tradeengineproto "github.com/sashajdn/sasha/service.trade-engine/proto"
)

// ReadInternalVenueAccount ...
func (s *TradeAccountService) ReadInternalVenueAccount(
	ctx context.Context, in *tradeaccountproto.ReadInternalVenueAccountRequest,
) (*tradeaccountproto.ReadInternalVenueAccountResponse, error) {
	switch {
	case in.Venue == tradeengineproto.VENUE_UNREQUIRED:
		return nil, gerrors.BadParam("missing_param.venue", nil)
	case in.ActorId == "":
		return nil, gerrors.BadParam("missing_param.actor_id", nil)
	}

	if ok := isValidActorID(in.ActorId); !ok {
		return nil, gerrors.Unauthenticated("failed_to_read_internal_venue_account.unauthenticated", nil)
	}

	errParams := map[string]string{
		"actor_id":           in.ActorId,
		"venue":              in.Venue.String(),
		"subaccount":         in.Subaccount,
		"venue_account_type": in.VenueAccountType.String(),
	}

	internalVenueAccount, err := dao.ReadInternalVenueAccount(ctx, in.Venue.String(), in.Subaccount, in.VenueAccountType.String())
	if err != nil {
		return nil, gerrors.Augment(err, "failed_to_read_internal_venue_account.dao", errParams)
	}

	proto, err := marshaling.InternalVenueAccountDomainToProto(internalVenueAccount)
	if err != nil {
		return nil, gerrors.Augment(err, "failed_to_read_internal_venue_account.marshal", errParams)
	}

	return &tradeaccountproto.ReadInternalVenueAccountResponse{
		InternalVenueAccount: proto,
	}, nil
}
