package execution

import (
	"context"
	"fmt"
	"time"

	discordproto "github.com/sashajdn/sasha/service.discord/proto"
	tradeaccountproto "github.com/sashajdn/sasha/service.trade-account/proto"

	"github.com/sashajdn/sasha/libraries/gerrors"
	"github.com/sashajdn/sasha/libraries/util"
	"github.com/sashajdn/sasha/service.trade-engine/marshaling"
	tradeengineproto "github.com/sashajdn/sasha/service.trade-engine/proto"
)

func readVenueCredentials(ctx context.Context, userID string, venue tradeengineproto.VENUE) (*tradeengineproto.VenueCredentials, error) {
	// Read venue account.
	rsp, err := (&tradeaccountproto.ReadVenueAccountByVenueAccountDetailsRequest{
		Venue:          venue,
		UserId:         userID,
		ActorId:        tradeaccountproto.ActorSystemTradeEngine,
		RequestContext: tradeaccountproto.RequestContextOrderRequest,
	}).Send(ctx).Response()
	if err != nil {
		return nil, gerrors.Augment(err, "failed_to_read_venue_credentials", nil)
	}

	// Marshal venue credentials.
	return marshaling.VenueAccountToVenueCredentials(rsp.GetVenueAccount()), nil
}

func readVenueAccountBalance(ctx context.Context, venue tradeengineproto.VENUE, tradeStrategy *tradeengineproto.TradeStrategy, credentials *tradeengineproto.VenueCredentials) (float64, error) {
	errParams := map[string]string{
		"venue": venue.String(),
	}

	switch venue {
	default:
		return 0, gerrors.Unimplemented("failed_to_read_venue_account_balance.unimplemented.venue", errParams)
	}
}

func notifyUser(ctx context.Context, msg string, userID string) error {
	content := `:wave: <@%s> WARNING FROM TRADE ENGINE:

%s
`

	formattedContent := fmt.Sprintf(content, userID, msg)
	if _, err := (&discordproto.SendMsgToPrivateChannelRequest{
		UserId:         userID,
		SenderId:       tradeengineproto.TradeEngineActorSatoshiSystem,
		Content:        formattedContent,
		IdempotencyKey: fmt.Sprintf("%s-%s-%s", userID, util.Sha256Hash(msg), time.Now().UTC().Truncate(10*time.Minute)),
	}).Send(ctx).Response(); err != nil {
		return gerrors.Augment(err, "failed_to_notify_user", nil)
	}

	return nil
}
