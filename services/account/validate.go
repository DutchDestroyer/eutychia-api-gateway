package account

import account "github.com/DutchDestroyer/eutychia-api-gateway/database"

// IsResearcherAccount determines whether the account is a researcher account, which means it has certain admin rights
func IsResearcherAccount(accountID string) (bool, error) {

	acc, err1 := account.GetDatabaseEntry(accountID)

	if err1 != nil {
		return false, err1
	}

	return (acc.AccountType == "researcher"), nil
}
