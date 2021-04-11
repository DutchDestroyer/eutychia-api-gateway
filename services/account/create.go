package account

import (
	"github.com/DutchDestroyer/eutychia-api-gateway/database"
	"github.com/DutchDestroyer/eutychia-api-gateway/models"
	"github.com/DutchDestroyer/eutychia-api-gateway/services/authentication"
)

//GetAccount creates an account after making an http request after logging in
func GetAccount(emailAddress string, password string, token string, accountID string, sessionID string) *models.Account {

	return &models.Account{
		Username:     emailAddress,
		Password:     password,
		RefreshToken: token,
		AccountID:    accountID,
		SessionID:    sessionID,
	}
}

// FinaleAccountCreation finalizes the creation of the account by adding the password, but first checks though if this is legitimate
func FinaleAccountCreation(accountID string, emailAddress string, password string, firstName string, lastName string) (bool, error) {
	isNewAccount, err1 := IsNewAccount(accountID, emailAddress, firstName, lastName)

	if err1 != nil || !isNewAccount {
		return isNewAccount, err1
	}

	encPW, err2 := authentication.EncryptPassword(password)

	if err2 != nil {
		return isNewAccount, err2
	}

	err3 := database.FinalizeAccountCreation(accountID, encPW)

	return isNewAccount, err3
}
