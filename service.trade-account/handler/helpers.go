package handler

import (
	"github.com/sashajdn/sasha/libraries/gerrors"
	"github.com/sashajdn/sasha/service.trade-account/domain"
	tradeaccountproto "github.com/sashajdn/sasha/service.trade-account/proto"
	tradeengineproto "github.com/sashajdn/sasha/service.trade-engine/proto"
)

func getIdentifierFromAccount(account *domain.Account, pagerType string) (string, error) {
	errParams := map[string]string{
		"user_id": account.UserID,
	}

	switch pagerType {
	case tradeaccountproto.PagerType_DISCORD.String():
		return account.UserID, nil

	case tradeaccountproto.PagerType_EMAIL.String():
		if account.Email == "" {
			return "", gerrors.FailedPrecondition("failed_to_get_identifier_from_account.email", errParams)
		}
		return account.Email, nil
	case tradeaccountproto.PagerType_PHONE.String():
		if account.PhoneNumber == "" {
			return "", gerrors.FailedPrecondition("failed_to_get_identifier_from_account.phone_number", errParams)
		}
		return account.PhoneNumber, nil
	case tradeaccountproto.PagerType_SMS.String():
		if account.PhoneNumber == "" {
			return "", gerrors.FailedPrecondition("failed_to_get_identifier_from_account.sms", errParams)
		}
		return account.PhoneNumber, nil
	}

	errParams["pager_type"] = pagerType
	return "", gerrors.FailedPrecondition("failed_to_get_identifier_from_account.unknown_pager_type", errParams)
}

func isValidActorID(actorID string) bool {
	switch actorID {
	case tradeaccountproto.ActorSystemPayments, tradeaccountproto.ActorManual, tradeaccountproto.ActorSystemTradeEngine:
		return true
	default:
		return false
	}
}

func isValidActorUnmaskedRequest(actorID string, isRequestingUnmaskedCredentials bool) bool {
	if !isRequestingUnmaskedCredentials {
		return true
	}

	if actorID != tradeaccountproto.ActorSystemTradeEngine {
		return false
	}

	return true
}

func validateVenueAccount(venueAccount *tradeaccountproto.VenueAccount) error {
	if venueAccount == nil {
		return gerrors.BadParam("missing_param.venue_account", nil)
	}

	switch {
	case venueAccount.ApiKey == "":
		return gerrors.BadParam("missing_param.api_key", nil)
	case venueAccount.SecretKey == "":
		return gerrors.BadParam("missing_param.secret_key", nil)
	}

	switch venueAccount.Venue {
	case tradeengineproto.VENUE_UNREQUIRED:
		return gerrors.BadParam("missing_param.venue_account.venue", nil)
	}

	return nil
}

func validateInternalVenueAccount(internalVenueAccount *tradeaccountproto.InternalVenueAccount) error {
	if internalVenueAccount == nil {
		return gerrors.BadParam("missing_param.internal_venue_account", nil)
	}

	switch {
	case internalVenueAccount.ApiKey == "":
		return gerrors.BadParam("missing_param.internal_venue_account.api_key", nil)
	case internalVenueAccount.SecretKey == "":
		return gerrors.BadParam("missing_param.internal_venue_account.secret_key", nil)
	case internalVenueAccount.Url == "":
		return gerrors.BadParam("missing_param.internal_venue_account.url", nil)
	case internalVenueAccount.WsUrl == "":
		return gerrors.BadParam("missing_param.internal_venue_account.ws_url", nil)
	}

	switch internalVenueAccount.Venue {
	case tradeengineproto.VENUE_UNREQUIRED:
		return gerrors.BadParam("missing_param.internal_venue_account.venue", nil)
	}

	return nil
}
