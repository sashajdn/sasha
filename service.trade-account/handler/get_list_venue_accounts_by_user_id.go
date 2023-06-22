package handler

import (
	"context"
	"strconv"

	"github.com/sashajdn/sasha/libraries/gerrors"
	"github.com/sashajdn/sasha/service.trade-account/dao"
	"github.com/sashajdn/sasha/service.trade-account/domain"
	"github.com/sashajdn/sasha/service.trade-account/marshaling"
	tradeaccountproto "github.com/sashajdn/sasha/service.trade-account/proto"
)

// ListVenueAccounts ...
func (s *TradeAccountService) ListVenueAccounts(
	ctx context.Context, in *tradeaccountproto.ListVenueAccountsRequest,
) (*tradeaccountproto.ListVenueAccountsResponse, error) {
	switch {
	case in.UserId == "":
		return nil, gerrors.FailedPrecondition("missing-param.user_id", nil)
	case !isValidActorUnmaskedRequest(in.ActorId, in.WithUnmaskedCredentials):
		return nil, gerrors.Unauthenticated("failed_to_list_venue_accounts_by_user_id.unauthenticated", map[string]string{
			"user_id":  "user_id",
			"actor_id": "actor_id",
		})
	}

	errParams := map[string]string{
		"user_id":  in.UserId,
		"actor_id": in.ActorId,
		"unmasked": strconv.FormatBool(in.WithUnmaskedCredentials),
	}

	// Determine the correct marshaller to user depending on whether the requester is allowed
	// unmasked credentials.
	var marshaller func([]*domain.VenueAccount) ([]*tradeaccountproto.VenueAccount, error)
	switch {
	case in.WithUnmaskedCredentials:
		marshaller = marshaling.VenueAccountDomainsToProtosUnmasked
	default:
		marshaller = marshaling.VenueAccountDomainsToProtos
	}

	venueAccounts, err := dao.ListVenueAccountsByUserID(ctx, in.UserId, in.GetActiveOnly())
	switch {
	case gerrors.Is(err, gerrors.ErrNotFound, "venue_accounts_not_found_for_user_id"):
		return nil, gerrors.Augment(err, "failed_to_list_venue_accounts_by_user_id", errParams)
	case err != nil:
		return nil, gerrors.Augment(err, "failed_to_read_venue_accounts_by_user_id", errParams)
	}

	protos, err := marshaller(venueAccounts)
	if err != nil {
		return nil, gerrors.Augment(err, "failed_to_list_venue_accounts.at_least_one_venue_account_has_an_unsupported_venue_account_type", errParams)
	}

	return &tradeaccountproto.ListVenueAccountsResponse{
		VenueAccounts: protos,
	}, nil
}
