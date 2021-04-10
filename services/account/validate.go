package account

import (
	"errors"

	"github.com/DutchDestroyer/eutychia-api-gateway/database"
	account "github.com/DutchDestroyer/eutychia-api-gateway/database"
)

// IsResearcherAccount determines whether the account is a researcher account, which means it has certain admin rights
func IsResearcherAccount(accountID string) (bool, error) {

	acc, err1 := account.GetDatabaseEntry(accountID)

	if err1 != nil {
		return false, err1
	}

	return (acc.AccountType == "researcher"), nil
}

func IsNewAccount(accountID string, emailAddress string, firstName string, lastName string) (bool, error) {

	account, err := database.GetDatabaseEntry(accountID)

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
