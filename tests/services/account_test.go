package services_test

import (
	"errors"
	"testing"

	accountDB "github.com/DutchDestroyer/eutychia-api-gateway/database"
	"github.com/DutchDestroyer/eutychia-api-gateway/services"
	"github.com/DutchDestroyer/eutychia-api-gateway/tests"
)

type accDBTest struct{}

var finalizeAccountCreationMock func(accountID string, encryptedpassword []byte) error
var getDatabaseEntryMock func(accountID string) (accountDB.AccountDAO, error)

func (a accDBTest) finalizeAccountCreation(accountID string, encryptedpassword []byte) error {
	return finalizeAccountCreationMock(accountID, encryptedpassword)
}

func (a accDBTest) getDatabaseEntry(accountID string) (accountDB.AccountDAO, error) {
	return getDatabaseEntryMock(accountID)
}

func TestCreateAccountWithValidEmailAddress(t *testing.T) {

	email := "wijnbergenmark@gmail.com"
	pw := "password"
	token := "token"
	accountID := "accountID"
	sessionID := "sessionID"

	account, err := services.GetAccount(email, pw, token, accountID, sessionID)

	tests.CompareErrors(t, err, nil)
	tests.CompareStrings(t, account.Username.EmailAddress, email)
	tests.CompareStrings(t, account.Password, pw)
	tests.CompareStrings(t, account.RefreshToken, token)
	tests.CompareStrings(t, account.AccountID, accountID)
	tests.CompareStrings(t, account.SessionID, sessionID)
	tests.CompareStrings(t, account.AuthToken, "")
	tests.CompareStrings(t, account.AccountType, "")
}

func TestCreateAccountWithInvalidEmailAddress(t *testing.T) {
	email := "invalid"
	pw := "password"
	token := "token"
	accountID := "accountID"
	sessionID := "sessionID"

	account, err := services.GetAccount(email, pw, token, accountID, sessionID)

	tests.CompareErrors(t, err, errors.New("invalid format"))
	tests.CompareStrings(t, account.Username.EmailAddress, "")
	tests.CompareStrings(t, account.Password, "")
	tests.CompareStrings(t, account.RefreshToken, "")
	tests.CompareStrings(t, account.AccountID, "")
	tests.CompareStrings(t, account.SessionID, "")
	tests.CompareStrings(t, account.AuthToken, "")
	tests.CompareStrings(t, account.AccountType, "")
}
