package handler

import (
	"context"

	"github.com/monzo/slog"

	"github.com/sashajdn/sasha/libraries/gerrors"
	"github.com/sashajdn/sasha/service.trade-account/dao"
	"github.com/sashajdn/sasha/service.trade-account/marshaling"
	tradeaccountproto "github.com/sashajdn/sasha/service.trade-account/proto"
)

// UpdateAccount ...
func (s *TradeAccountService) UpdateAccount(
	ctx context.Context, in *tradeaccountproto.UpdateAccountRequest,
) (*tradeaccountproto.UpdateAccountResponse, error) {
	// Validation.
	switch {
	case in.UserId == "":
		return nil, gerrors.BadParam("missing_param.user_id", nil)
	case (in.IsAdmin || in.IsFutures) && !isValidActorID(in.ActorId):
		// Here if the user is setting a users futures or admin level; then we require a certain actor id.
		return nil, gerrors.BadParam("missing_param.actor_id", map[string]string{
			"actor_id": in.ActorId,
		})
	}

	errParams := map[string]string{
		"user_id": in.UserId,
	}

	// Create mutation in domain.
	mutation := marshaling.UpdateAccountProtoToDomain(in)

	// Apply mutation.
	account, err := dao.UpdateAccount(ctx, mutation)
	if err != nil {
		return nil, gerrors.Augment(err, "failed_to_update_account", errParams)
	}

	slog.Info(ctx, "Updated account %s: %v", account.UserID, account.Updated)

	return &tradeaccountproto.UpdateAccountResponse{
		Account: marshaling.AccountDomainToProto(account),
	}, nil
}
