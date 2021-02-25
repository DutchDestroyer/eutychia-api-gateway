package authentication

import models "github.com/DutchDestroyer/eutychia-api-gateway/models"

// IsValidTokenLogin determines whether the token validation is done correctly
func IsValidTokenLogin(acc models.Account) error {
	return ValidateToken(acc.RefreshToken, acc.AccountID, acc.SessionID, "refreshToken")
}
