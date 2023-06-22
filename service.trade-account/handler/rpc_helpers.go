package handler

import (
	"context"
	"fmt"
	"time"

	"github.com/monzo/slog"

	"github.com/sashajdn/sasha/libraries/gerrors"
	discordproto "github.com/sashajdn/sasha/service.discord/proto"
	tradeaccountproto "github.com/sashajdn/sasha/service.trade-account/proto"
	tradeengineproto "github.com/sashajdn/sasha/service.trade-engine/proto"
)

func validateVenueCredentials(ctx context.Context, userID string, venueAccount interface{}) (bool, string, error) {
	var credentials *tradeengineproto.VenueCredentials
	switch t := venueAccount.(type) {
	case *tradeaccountproto.VenueAccount:
		credentials = &tradeengineproto.VenueCredentials{
			Venue:      t.Venue,
			ApiKey:     t.ApiKey,
			SecretKey:  t.SecretKey,
			Subaccount: t.SubAccount,
			Url:        t.Url,
			WsUrl:      t.WsUrl,
		}
	case *tradeaccountproto.InternalVenueAccount:
		credentials = &tradeengineproto.VenueCredentials{
			Venue:      t.Venue,
			ApiKey:     t.ApiKey,
			SecretKey:  t.SecretKey,
			Subaccount: t.SubAccount,
			Url:        t.Url,
			WsUrl:      t.WsUrl,
		}
	default:
		slog.Error(ctx, "Failed to validate venue credentials, invalid type: %T", t)
		return false, "", gerrors.Unimplemented("unimplemented_account_type", map[string]string{
			"account_type": fmt.Sprintf("%T", t),
		})
	}

	errParams := map[string]string{
		"venue": credentials.Venue.String(),
	}

	// Validate venue credentials.
	switch credentials.Venue {
	case tradeengineproto.VENUE_BITFINEX:
		return false, "", gerrors.Unimplemented("venue_unimplemented.bitfinex", nil)
	case tradeengineproto.VENUE_DERIBIT:
		return false, "", gerrors.Unimplemented("venue_unimplemented.deribit", nil)
	default:
		return false, "", gerrors.FailedPrecondition("failed_to_validate_credentials.invalid_venue_account", errParams)
	}
}

func notifyPulseChannel(ctx context.Context, userID, username string, timestamp time.Time) error {
	base := ":bear:    `NEW MEMBER`    :bear:"
	msg := `
UserID: %s
Username: %s
Timestamp: %v
`
	formattedMsg := fmt.Sprintf(msg, userID, username, timestamp)

	if _, err := (&discordproto.SendMsgToChannelRequest{
		ChannelId: discordproto.DiscordSatoshiAccountsPulseChannel,
		Content:   fmt.Sprintf("%s```%s```", base, formattedMsg),
	}).Send(ctx).Response(); err != nil {
		return gerrors.Augment(err, "failed_to_notify_account_pulse_channel", nil)
	}

	return nil
}
