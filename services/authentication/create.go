package authentication

import (
	"crypto/rsa"
	"time"

	database "github.com/DutchDestroyer/eutychia-api-gateway/database"
	models "github.com/DutchDestroyer/eutychia-api-gateway/models"
	"github.com/google/uuid"
)

// CreateAccountAuthentication create authentication of the account
func CreateAccountAuthentication(account *models.Account) error {
	sessionID := uuid.New().String()

	authToken, authTokenKey, authErr := createAuthToken(account.AccountID, sessionID)
	if authErr != nil {
		return authErr
	}

	refreshToken, refreshTokenKey, refreshErr := createRefreshToken(account.AccountID, sessionID)
	if refreshErr != nil {
		return refreshErr
	}

	dbErr := database.StoreSession(account.AccountID, sessionID, authToken, authTokenKey, refreshToken, refreshTokenKey)

	if dbErr != nil {
		return dbErr
	}

	account.SessionID = sessionID
	account.AuthToken = authToken
	account.RefreshToken = refreshToken

	return nil
}

// UpdateAccountAuthentication create authentication of the account when logging in with refreshtoken
func UpdateAccountAuthentication(account *models.Account) error {

	authToken, authTokenKey, authErr := createAuthToken(account.AccountID, account.SessionID)
	if authErr != nil {
		return authErr
	}

	dbErr := database.UpdateSessionAuthToken(account.AccountID, account.SessionID, authToken, authTokenKey)
	if dbErr != nil {
		return dbErr
	}

	account.AuthToken = authToken

	return nil
}

func createAuthToken(accountID string, sessionID string) (string, rsa.PublicKey, error) {
	return CreateToken(accountID, sessionID, time.Duration(time.Duration.Minutes(15)))
}

func createRefreshToken(accountID string, sessionID string) (string, rsa.PublicKey, error) {
	return CreateToken(accountID, sessionID, time.Duration(time.Duration.Hours(24)))
}
