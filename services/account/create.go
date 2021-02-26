package account

import "github.com/DutchDestroyer/eutychia-api-gateway/models"

//CreateAccount creates an account after making an http request after logging in
func CreateAccount(emailAddress string, password string, token string, accountID string, sessionID string) *models.Account {

	return &models.Account{
		Username:     emailAddress,
		Password:     password,
		RefreshToken: token,
		AccountID:    accountID,
		SessionID:    sessionID,
	}
}
