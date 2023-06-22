package dao

import (
	"context"

	"github.com/sashajdn/sasha/service.trade-account/domain"
)

// ReadVenueAccountByAccountAlias ...
func ReadVenueAccountByAccountAlias(ctx context.Context, userID, accountAlias string) (*domain.VenueAccount, error) {
    return nil, nil
}

// ReadVenueAccountByVenueAccountID ...
func ReadVenueAccountByVenueAccountID(ctx context.Context, venueAccountID string) (*domain.VenueAccount, error) {
    return nil, nil
}

// ReadVenueAccountByVenueAccountDetails ...
func ReadVenueAccountByVenueAccountDetails(ctx context.Context, venueID, userID string) (*domain.VenueAccount, error) {
    return nil, nil
}

func ReadInternalVenueAccount(ctx context.Context, venueID, subaccount, internalAccountType string) (*domain.InternalVenueAccount, error) {
    return nil, nil
}

// ListVenueAccountsByUserID ...
func ListVenueAccountsByUserID(ctx context.Context, userID string, isActive bool) ([]*domain.VenueAccount, error) {
    return nil, nil
}

// AddVenueAccount ...
func AddVenueAccount(ctx context.Context, venueAccount *domain.VenueAccount) error {
    return nil
}

// CreateOrUpdateInternalVenueAccount ...
func CreateOrUpdateInternalVenueAccount(ctx context.Context, venueAccount *domain.InternalVenueAccount, allowUpdate bool) error {
    return nil
}

// RemoveVenueAccount ...
func RemoveVenueAccount(ctx context.Context, venueAccountID string) error {
    return nil
}

// UpdateVenueAccount ...
func UpdateVenueAccount(ctx context.Context, mutation *domain.VenueAccount) (*domain.VenueAccount, error) {
    return nil, nil
}
