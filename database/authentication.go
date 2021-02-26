package database

import (
	"crypto/rsa"
	"errors"
)

// AuthenticationData is the database model of the Authentication data
type AuthenticationData struct {
	AccountID       string
	SessionID       string
	AuthToken       string
	AuthTokenKey    rsa.PublicKey
	Refreshtoken    string
	RefreshTokenKey rsa.PublicKey
}

var authenticationTable []AuthenticationData

//StoreSession Stores the session in the database
func StoreSession(accountID string, sessionID string, authToken string, authPK rsa.PublicKey, refreshToken string, refreshPK rsa.PublicKey) error {
	authenticationTable = append(authenticationTable, AuthenticationData{accountID, sessionID, authToken, authPK, refreshToken, refreshPK})
	return nil
}

//UpdateSessionAuthToken Updates the authtoken of the session in the database
func UpdateSessionAuthToken(accountID string, sessionID string, authToken string, authPK rsa.PublicKey) error {
	err := errors.New("could not find session")

	for i := range authenticationTable {
		if authenticationTable[i].AccountID == accountID {
			authenticationTable[i].AuthToken = authToken
			authenticationTable[i].AuthTokenKey = authPK
			return nil
		}
	}
	return err
}

// GetSessionData retrieves a session from the database
func GetSessionData(accountID string, sessionID string) (AuthenticationData, error) {
	err := errors.New("could not find session")

	for i := range authenticationTable {
		if authenticationTable[i].AccountID == accountID && authenticationTable[i].SessionID == sessionID {
			return authenticationTable[i], nil
		}
	}
	return AuthenticationData{}, err
}
