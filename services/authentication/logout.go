package authentication

import "github.com/DutchDestroyer/eutychia-api-gateway/database"

func LogOutWithAccount(sessionID string, accountID string, accessToken string) error {
	err := ValidateToken(accessToken, accountID, sessionID, "authToken")

	if err != nil {
		return err
	}

	return database.RemoveSession(accountID, sessionID)
}
