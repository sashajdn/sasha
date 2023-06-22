package handler

import (
	"context"

	"github.com/sashajdn/sasha/libraries/gerrors"
	"github.com/sashajdn/sasha/service.trade-account/dao"
	"github.com/sashajdn/sasha/service.trade-account/marshaling"
	tradeaccountproto "github.com/sashajdn/sasha/service.trade-account/proto"
)

// ReadAccount reads an account via the user ID, which is the discord ID.
func (a *TradeAccountService) ReadAccount(
	ctx context.Context, in *tradeaccountproto.ReadAccountRequest,
) (*tradeaccountproto.ReadAccountResponse, error) {
	switch {
	case in.UserId == "":
		return nil, gerrors.BadParam("missing_param.user_id", nil)
	}

	errParams := map[string]string{
		"user_id": in.UserId,
	}

	account, err := dao.ReadAccountByUserID(ctx, in.UserId)
	switch {
	case gerrors.Is(err, gerrors.ErrNotFound, "account_not_found"):
		return nil, gerrors.Augment(err, "failed_to_read_account.account_not_exist", errParams)
	case err != nil:
		return nil, gerrors.Augment(err, "failed_to_read_account", errParams)
	}

	return &tradeaccountproto.ReadAccountResponse{
		Account: marshaling.AccountDomainToProto(account),
	}, nil
}
