package services

import (
	"database/sql"
	"errors"

	"github.com/DutchDestroyer/eutychia-api-gateway/database"
)

type IAccountService interface {
	FinaleAccountCreation(string, string, *sql.Tx) (bool, error)
	IsResearcherAccount(string, *sql.Tx) (bool, error)
}

type AccountService struct {
	AccDBService database.IAccountDBService
	AuthService  IAuthenticationService
}

// FinaleAccountCreation finalizes the creation of the account by adding the password, but first checks though if this is legitimate
func (a AccountService) FinaleAccountCreation(accountID string, password string, tx *sql.Tx) (bool, error) {

	account, err := a.AccDBService.GetDatabaseEntry(accountID)

	if err != nil {
		return false, err
	}

	if account.FirstName == "" {
		return false, errors.New("account does not exist")
	}

	hasNoPassword, err := a.isAccountWithoutPassword(accountID, account)

	if err != nil || hasNoPassword {
		return hasNoPassword, err
	}

	encPW, err := a.AuthService.encryptPassword(password)

	if err != nil {
		return hasNoPassword, err
	}

	err = a.AccDBService.FinalizeAccountCreation(accountID, encPW)

	return hasNoPassword, err
}

// IsResearcherAccount determines whether the account is a researcher account, which means it has certain admin rights
func (a AccountService) IsResearcherAccount(accountID string, tx *sql.Tx) (bool, error) {

	acc, err1 := a.AccDBService.GetDatabaseEntry(accountID)

	if err1 != nil {
		return false, err1
	}

	return (acc.AccountType == "researcher"), nil
}

func (a AccountService) isAccountWithoutPassword(accountID string, account database.AccountDAO) (bool, error) {

	if account.Password == "" {
		return true, nil
	}

	return false, errors.New("account already has a password")
}
