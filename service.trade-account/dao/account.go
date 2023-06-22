package dao

import (
	"context"

	"github.com/sashajdn/sasha/service.trade-account/domain"
)

// ListAccounts returns a list of all domain accounts from the underlying datastore.
func ListAccounts(ctx context.Context) ([]*domain.Account, error) {
    // var sql = ` SELECT * FROM s_account_accounts`

	return nil, nil
}

//
func ListFuturesMembers(ctx context.Context) ([]*domain.Account, error) {
	// var sql = ` SELECT * FROM s_account_accounts WHERE is_futures_member=true`

	return nil, nil
}

// ReadAccountByUserID returns the a domain account from the underlying datastore, by UserID.
func ReadAccountByUserID(ctx context.Context, userID string) (*domain.Account, error) {
    // var sql      = `SELECT * FROM s_account_accounts WHERE user_id=$1`
	return nil, nil
}

// ReadAccountByUsername returns the a domain account from the underlying datastore, by username.
func ReadAccountByUsername(ctx context.Context, username string) (*domain.Account, error) {
	// var sql      = `SELECT * FROM s_account_accounts WHERE username=$1`

    return nil, nil
}

// CreateAccount creates a new account in the datastore.
func CreateAccount(ctx context.Context, account *domain.Account) error {
	return nil
}

// UpdateAccount updates an already existing account; it will perform a mutation on the existing account.
// AccountID must be provided at least to the passed domain account struct, `mutation`.
func UpdateAccount(ctx context.Context, mutation *domain.Account) (*domain.Account, error) {
	return nil, nil
}
