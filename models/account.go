package models

// Account contains all data of the account
type Account struct {
	Username string

	Password string

	AuthToken string

	RefreshToken string

	AccountID string

	AccountType string

	SessionID string
}
