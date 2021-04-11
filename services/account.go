package services

import (
	"errors"

	accountDB "github.com/DutchDestroyer/eutychia-api-gateway/database"
	models "github.com/DutchDestroyer/eutychia-api-gateway/models"
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

	encPW, err2 := EncryptPassword(password)

	if err2 != nil {
		return isNewAccount, err2
	}

	err3 := accountDB.FinalizeAccountCreation(accountID, encPW)

	return isNewAccount, err3
}

// IsResearcherAccount determines whether the account is a researcher account, which means it has certain admin rights
func IsResearcherAccount(accountID string) (bool, error) {

	acc, err1 := accountDB.GetDatabaseEntry(accountID)

	if err1 != nil {
		return false, err1
	}

	return (acc.AccountType == "researcher"), nil
}

func IsNewAccount(accountID string, emailAddress string, firstName string, lastName string) (bool, error) {

	account, err := accountDB.GetDatabaseEntry(accountID)

	if err != nil {
		return true, err
	}

	if account.EmailAddress != emailAddress || account.FirstName != firstName || account.LastName != lastName {
		return false, errors.New("data in database is different from given values")
	}

	if account.Password != "" {
		return false, errors.New("Password has already been created")
	}

	return true, nil
}
