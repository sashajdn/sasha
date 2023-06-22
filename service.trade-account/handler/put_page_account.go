package handler

import (
	"context"

	"github.com/monzo/terrors"

	"github.com/sashajdn/sasha/service.trade-account/dao"
	"github.com/sashajdn/sasha/service.trade-account/pager"
	tradeaccountproto "github.com/sashajdn/sasha/service.trade-account/proto"
)

func (s *TradeAccountService) PageAccount(
	ctx context.Context, in *tradeaccountproto.PageAccountRequest,
) (*tradeaccountproto.PageAccountResponse, error) {
	errParams := map[string]string{
		"account_id": in.UserId,
	}

	account, err := dao.ReadAccountByUserID(ctx, in.UserId)
	if err != nil {
		return nil, terrors.Augment(err, "Failed to read account", errParams)
	}

	var pagerType string
	switch in.Priority {
	case tradeaccountproto.PagerPriority_HIGH:
		pagerType = account.HighPriorityPager
	case tradeaccountproto.PagerPriority_LOW:
		pagerType = account.LowPriorityPager
	default:
		pagerType = account.LowPriorityPager
	}

	pager, err := pager.GetPagerByID(pagerType)
	if err != nil {
		errParams["pager_type"] = pagerType
		return nil, terrors.Augment(err, "Invalid pager type set to account", errParams)
	}

	identifier, err := getIdentifierFromAccount(account, pagerType)
	if err != nil {
		return nil, terrors.Augment(err, "Cannot page user; missing identifier on account", errParams)
	}

	if err := pager.Page(ctx, identifier, in.Content); err != nil {
		return nil, terrors.Augment(err, "Failed to page user", errParams)
	}

	return &tradeaccountproto.PageAccountResponse{}, nil
}
