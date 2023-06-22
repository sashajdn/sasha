package handler

import (
	"context"

	"github.com/sashajdn/sasha/libraries/gerrors"
	"github.com/sashajdn/sasha/service.trade-account/dao"
	"github.com/sashajdn/sasha/service.trade-account/marshaling"
	tradeaccountproto "github.com/sashajdn/sasha/service.trade-account/proto"
)

// ReadVenueAccountByVenueAccountID ...
func (s *TradeAccountService) ReadVenueAccountByVenueAccountID(
	ctx context.Context, in *tradeaccountproto.ReadVenueAccountByVenueAccountIDRequest,
) (*tradeaccountproto.ReadVenueAccountByVenueAccountIDResponse, error) {
	switch {
	case in.VenueAccountId == "":
		return nil, gerrors.BadParam("missing_param.venue_account_id", nil)
	}

	errParams := map[string]string{
		"venue_account_id": in.VenueAccountId,
	}

	exchange, err := dao.ReadVenueAccountByVenueAccountID(ctx, in.VenueAccountId)
	if err != nil {
		return nil, gerrors.Augment(err, "failed_to_read_exchange_by_venue_account_id.dao", errParams)
	}

	proto, err := marshaling.VenueAccountDomainToProto(exchange)
	if err != nil {
		return nil, gerrors.Augment(err, "failed_to_read_exchange_by_venue_id.marshal_to_proto", errParams)
	}

	return &tradeaccountproto.ReadVenueAccountByVenueAccountIDResponse{
		VenueAccount: proto,
	}, nil
}
