package handler

import (
	"context"

	"github.com/monzo/terrors"

	"github.com/sashajdn/sasha/service.trade-account/dao"
	"github.com/sashajdn/sasha/service.trade-account/domain"
	"github.com/sashajdn/sasha/service.trade-account/marshaling"
	tradeaccountproto "github.com/sashajdn/sasha/service.trade-account/proto"
)

// ListAccounts returns a list of all given accounts.
func (s *TradeAccountService) ListAccounts(
	ctx context.Context, in *tradeaccountproto.ListAccountsRequest,
) (*tradeaccountproto.ListAccountsResponse, error) {
	var (
		accounts []*domain.Account
		err      error
	)

	switch {
	case in.IsFuturesMember:
		accounts, err = dao.ListFuturesMembers(ctx)
	default:
		accounts, err = dao.ListAccounts(ctx)
	}

	if err != nil {
		return nil, terrors.Augment(err, "Failed to list accounts", nil)

	}

	var protoAccounts []*tradeaccountproto.Account
	for _, account := range accounts {
		protoAccounts = append(protoAccounts, marshaling.AccountDomainToProto(account))
	}

	return &tradeaccountproto.ListAccountsResponse{
		Accounts: protoAccounts,
	}, nil
}
