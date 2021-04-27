package models

// Account contains all data of the account
type Account struct {
	Username IEmailAddress

	Password string

	AuthToken string

	RefreshToken string

	AccountID string

	AccountType string

	SessionID string

	GrantType string
}
