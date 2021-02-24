package authentication

import (
	"time"

	accountmodels "github.com/DutchDestroyer/eutychia-api-gateway/account/models"
	database "github.com/DutchDestroyer/eutychia-api-gateway/database"
	"github.com/google/uuid"
)

// CreateAccountAuthentication create authentication of the account
func CreateAccountAuthentication(account *accountmodels.Account) error {
	sessionID := uuid.New().String()

	authToken, authErr := createAuthToken(account.AccountID, sessionID)
	if authErr != nil {
		return authErr
	}
	refreshToken, refreshErr := createRefreshToken(account.AccountID, sessionID)
	if refreshErr != nil {
		return refreshErr
	}

	dbErr := database.StoreSession(account.AccountID, sessionID, authToken, refreshToken)

	if dbErr != nil {
		return dbErr
	}

	account.SessionID = sessionID
	account.AuthToken = authToken
	account.RefreshToken = refreshToken

	return nil
}

// UpdateAccountAuthentication create authentication of the account when logging in with refreshtoken
func UpdateAccountAuthentication(account *accountmodels.Account) error {
	sessionID := uuid.New().String()

	authToken, authErr := createAuthToken(account.AccountID, sessionID)
	if authErr != nil {
		return authErr
	}

	dbErr := database.StoreSession(account.AccountID, sessionID, authToken, account.RefreshToken)

	if dbErr != nil {
		return dbErr
	}

	account.SessionID = sessionID
	account.AuthToken = authToken

	return nil
}

func createAuthToken(accountID string, sessionID string) (string, error) {
	return CreateToken(accountID, sessionID, time.Duration(time.Duration.Minutes(15)))
}

func createRefreshToken(accountID string, sessionID string) (string, error) {
	return CreateToken(accountID, sessionID, time.Duration(time.Duration.Hours(24)))
}
