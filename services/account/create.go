package account

import accountmodels "github.com/DutchDestroyer/eutychia-api-gateway/account/models"

//CreateAccount creates an account after making an http request after logging in
func CreateAccount(emailAddress string, password string, token string, accountID string, sessionID string) *accountmodels.Account {

	return &accountmodels.Account{
		Username:     emailAddress,
		Password:     password,
		RefreshToken: token,
		AccountID:    accountID,
		SessionID:    sessionID,
	}
}
