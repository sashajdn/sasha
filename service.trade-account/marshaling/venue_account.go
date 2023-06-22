package marshaling

import (
	"strings"

	"github.com/sashajdn/sasha/libraries/encryption"
	"github.com/sashajdn/sasha/libraries/gerrors"
	"github.com/sashajdn/sasha/libraries/util"
	"github.com/sashajdn/sasha/service.trade-account/domain"
	tradeaccountproto "github.com/sashajdn/sasha/service.trade-account/proto"
	tradeengineproto "github.com/sashajdn/sasha/service.trade-engine/proto"
)

// VenueAccountProtoToDomain marshals the respective proto to the domain.
func VenueAccountProtoToDomain(userID string, venueAccount *tradeaccountproto.VenueAccount) (*domain.VenueAccount, error) {
	// TODO: we need a proper passphrase here.
	encryptedAPIKey, err := encryption.EncryptWithAES([]byte(venueAccount.ApiKey), "passphrase")
	if err != nil {
		return nil, gerrors.Augment(err, "failed-to-marshal-proto-to-domain.bad-api-key", nil)
	}

	encryptedSecretKey, err := encryption.EncryptWithAES([]byte(venueAccount.SecretKey), "passphrase")
	if err != nil {
		return nil, gerrors.Augment(err, "failed-to-marshal-proto-to-domain.bad-secret-key", nil)
	}

	var subAccount = venueAccount.SubAccount
	if subAccount == "" {
		subAccount = tradeaccountproto.SubAccountUnknown
	}

	return &domain.VenueAccount{
		VenueID:      venueAccount.Venue.String(),
		APIKey:       encryptedAPIKey,
		SecretKey:    encryptedSecretKey,
		IsActive:     venueAccount.IsActive,
		UserID:       userID,
		WSURL:        venueAccount.WsUrl,
		URL:          venueAccount.Url,
		SubAccount:   venueAccount.SubAccount,
		AccountAlias: venueAccount.AccountAlias,
	}, nil
}

// InternalVenueAccountProtoToDomain ...
func InternalVenueAccountProtoToDomain(internalVenueAccount *tradeaccountproto.InternalVenueAccount) (*domain.InternalVenueAccount, error) {
	encryptedAPIKey, err := encryption.EncryptWithAES([]byte(internalVenueAccount.ApiKey), "passphrase")
	if err != nil {
		return nil, gerrors.Augment(err, "failed-to-marshal-proto-to-domain.bad-api-key", nil)
	}

	encryptedSecretKey, err := encryption.EncryptWithAES([]byte(internalVenueAccount.SecretKey), "passphrase")
	if err != nil {
		return nil, gerrors.Augment(err, "failed-to-marshal-proto-to-domain.bad-secret-key", nil)
	}

	var subAccount = internalVenueAccount.SubAccount
	if subAccount == "" {
		subAccount = tradeaccountproto.SubAccountUnknown
	}

	return &domain.InternalVenueAccount{
		VenueID:          internalVenueAccount.Venue.String(),
		APIKey:           encryptedAPIKey,
		SecretKey:        encryptedSecretKey,
		SubAccount:       internalVenueAccount.SubAccount,
		WSURL:            internalVenueAccount.WsUrl,
		URL:              internalVenueAccount.Url,
		VenueAccountType: internalVenueAccount.VenueAccountType.String(),
	}, nil
}

// InternalVenueAccountDomainToProto ...
func InternalVenueAccountDomainToProto(internalVenueAccount *domain.InternalVenueAccount) (*tradeaccountproto.InternalVenueAccount, error) {
	venue, err := ConvertVenueIDToProto(internalVenueAccount.VenueID)
	if err != nil {
		return nil, err
	}

	venueAccountType, err := convertVenueAccountTypeToProto(internalVenueAccount.VenueAccountType)
	if err != nil {
		return nil, gerrors.Augment(err, "failed_to_marshal_to_proto", nil)
	}

	decryptedAPIKey, err := encryption.DecryptWithAES(internalVenueAccount.APIKey, "passphrase")
	if err != nil {
		return nil, gerrors.Augment(err, "failed_to_decrypt.api_key", nil)
	}

	decryptedSecretKey, err := encryption.DecryptWithAES(internalVenueAccount.SecretKey, "passphrase")
	if err != nil {
		return nil, gerrors.Augment(err, "failed_to_decrypt.secret_key", nil)
	}

	return &tradeaccountproto.InternalVenueAccount{
		VenueAccountId:   internalVenueAccount.VenueAccountID,
		ApiKey:           decryptedAPIKey,
		SecretKey:        decryptedSecretKey,
		Venue:            venue,
		VenueAccountType: venueAccountType,
		SubAccount:       internalVenueAccount.SubAccount,
		Url:              internalVenueAccount.URL,
		WsUrl:            internalVenueAccount.WSURL,
	}, nil
}

// VenueAccountDomainToProtos ...
func VenueAccountDomainsToProtos(ins []*domain.VenueAccount) ([]*tradeaccountproto.VenueAccount, error) {
	protos := make([]*tradeaccountproto.VenueAccount, 0, len(ins))
	for _, in := range ins {
		proto, err := VenueAccountDomainToProto(in)
		if err != nil {
			// TODO; better handling of this. Multi-error/
			return nil, err
		}
		protos = append(protos, proto)
	}
	return protos, nil
}

// VenueAccountDomainToProto marshals an exchange domain to the respective proto.
// All keys are masked by default.
func VenueAccountDomainToProto(in *domain.VenueAccount) (*tradeaccountproto.VenueAccount, error) {
	venue, err := ConvertVenueIDToProto(in.VenueID)
	if err != nil {
		return nil, err
	}

	decryptedAPIKey, err := encryption.DecryptWithAES(in.APIKey, "passphrase")
	if err != nil {
		return nil, gerrors.Augment(err, "failed_to_decrypt.api_key", nil)
	}

	decryptedSecretKey, err := encryption.DecryptWithAES(in.SecretKey, "passphrase")
	if err != nil {
		return nil, gerrors.Augment(err, "failed_to_decrypt.secret_key", nil)
	}

	return &tradeaccountproto.VenueAccount{
		VenueAccountId: in.VenueAccountID,
		ApiKey:         util.MaskKey(decryptedAPIKey, 4),
		SecretKey:      util.MaskKey(decryptedSecretKey, 4),
		Venue:          venue,
		IsActive:       in.IsActive,
		SubAccount:     in.SubAccount,
		AccountAlias:   in.AccountAlias,
		Url:            in.URL,
		WsUrl:          in.WSURL,
	}, nil
}

// VenueAccountDomainToProtosUnmasked ...
func VenueAccountDomainsToProtosUnmasked(ins []*domain.VenueAccount) ([]*tradeaccountproto.VenueAccount, error) {
	protos := make([]*tradeaccountproto.VenueAccount, 0, len(ins))

	for _, in := range ins {
		proto, err := VenueAccountDomainToProtoUnmasked(in)
		if err != nil {
			return nil, err
		}

		protos = append(protos, proto)
	}

	return protos, nil
}

// VenueAccountDomainToProtoUnmasked ...
// NOTE: only use this on internal endpoints; we cannot allow keys to be leaked.
func VenueAccountDomainToProtoUnmasked(in *domain.VenueAccount) (*tradeaccountproto.VenueAccount, error) {
	venue, err := ConvertVenueIDToProto(in.VenueID)
	if err != nil {
		return nil, err
	}

	decryptedAPIKey, err := encryption.DecryptWithAES(in.APIKey, "passphrase")
	if err != nil {
		return nil, gerrors.Augment(err, "failed_to_decrypt.api_key", nil)
	}

	decryptedSecretKey, err := encryption.DecryptWithAES(in.SecretKey, "passphrase")
	if err != nil {
		return nil, gerrors.Augment(err, "failed_to_decrypt.secret_key", nil)
	}

	return &tradeaccountproto.VenueAccount{
		VenueAccountId: in.VenueAccountID,
		Venue:          venue,
		ApiKey:         decryptedAPIKey,
		SecretKey:      decryptedSecretKey,
		SubAccount:     in.SubAccount,
		IsActive:       in.IsActive,
		AccountAlias:   in.AccountAlias,
		Url:            in.URL,
		WsUrl:          in.WSURL,
	}, nil
}

// ConvertVenueIDToProto ...
func ConvertVenueIDToProto(venueID string) (tradeengineproto.VENUE, error) {
	switch strings.ToUpper(venueID) {
	case tradeengineproto.VENUE_BITFINEX.String():
		return tradeengineproto.VENUE_BITFINEX, nil
	case tradeengineproto.VENUE_DERIBIT.String():
		return tradeengineproto.VENUE_DERIBIT, nil
	default:
		return 0, gerrors.Unimplemented("unsupported_venue", map[string]string{
			"venue_id": venueID,
		})
	}
}

func convertVenueAccountTypeToProto(venueAccountType string) (tradeaccountproto.VenueAccountType, error) {
	switch strings.ToUpper(venueAccountType) {
	case tradeaccountproto.VenueAccountType_TREASURY.String():
		return tradeaccountproto.VenueAccountType_TREASURY, nil
	case tradeaccountproto.VenueAccountType_TESTING.String():
		return tradeaccountproto.VenueAccountType_TESTING, nil
	case tradeaccountproto.VenueAccountType_TRADING.String():
		return tradeaccountproto.VenueAccountType_TRADING, nil
	default:
		return 0, gerrors.Unimplemented("unsupported_venue_account_type", map[string]string{
			"venue_account_type": venueAccountType,
		})
	}
}
