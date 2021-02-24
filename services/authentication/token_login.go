package authentication

import accountmodels "github.com/DutchDestroyer/eutychia-api-gateway/account/models"

// IsValidTokenLogin determines whether the token validation is done correctly
func IsValidTokenLogin(acc accountmodels.Account) error {
	return ValidateToken(acc.AuthToken, acc.AccountID, acc.SessionID)
}
