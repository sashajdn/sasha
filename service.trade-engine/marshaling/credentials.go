package marshaling

import (
	tradeaccountproto "github.com/sashajdn/sasha/service.trade-account/proto"
	tradeengineproto "github.com/sashajdn/sasha/service.trade-engine/proto"
)

// VenueAccountToVenueCredentials ...
func VenueAccountToVenueCredentials(venueAccount *tradeaccountproto.VenueAccount) *tradeengineproto.VenueCredentials {
	return &tradeengineproto.VenueCredentials{
		ApiKey:     venueAccount.ApiKey,
		SecretKey:  venueAccount.SecretKey,
		Subaccount: venueAccount.SubAccount,
		Venue:      venueAccount.Venue,
	}
}
